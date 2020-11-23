package components

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

// SplashService the web service
type SplashService struct {
	authHTML     string
	auth         *oauth2.Config
	context      context.Context
	provider     *oidc.Provider
	oidcConfig   *oidc.Config
	verifier     *oidc.IDTokenVerifier
	state        string
	Host         string
	CallbackPath string
	Port         string
}

// NewSplash creates a new splash service
func NewSplash(valuesYAML string) SplashService {
	values := NewValues(valuesYAML)

	var html string
	if values.RawHTML != nil {
		html = *values.RawHTML
		html = strings.ReplaceAll(html, "%", "%%")

	} else {
		component := values.AsHTMLComponent()
		html = ToHTMLPage(component.ConsumeTemplate(), component.Styles())
	}

	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, values.Auth.ProviderURL)

	if err != nil {
		panic(err)
	}

	auth := values.Auth.ToOAuth(ctx, *provider)
	oidcConfig := &oidc.Config{ClientID: auth.ClientID}

	u, _ := url.Parse(values.Auth.RedirectURL)

	return SplashService{
		authHTML:     html,
		auth:         &auth,
		context:      ctx,
		provider:     provider,
		oidcConfig:   oidcConfig,
		verifier:     provider.Verifier(oidcConfig),
		state:        values.Auth.State,
		Host:         values.Host,
		CallbackPath: u.Path,
		Port:         values.Port,
	}
}

// MainSplashHandler if authorized, runs main page
func (s *SplashService) MainSplashHandler(w http.ResponseWriter, r *http.Request) {
	rawAccessToken := r.Header.Get("Authorization")

	if token := r.URL.Query().Get("token"); rawAccessToken == "" && token != "" {
		rawAccessToken = "Bearer " + token
	}

	if rawAccessToken == "" {
		http.Redirect(w, r, s.auth.AuthCodeURL(s.state), http.StatusFound)
		return
	}

	parts := strings.Split(rawAccessToken, " ")
	if len(parts) != 2 {
		w.WriteHeader(400)
		return
	}
	_, err := s.verifier.Verify(s.context, parts[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	if err != nil {
		http.Redirect(w, r, s.auth.AuthCodeURL(s.state), http.StatusFound)
		return
	}

	fmt.Fprintf(w, s.authHTML)
}

// CallbackHandler if authorized, runs main page
func (s *SplashService) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("state") != s.state {
		http.Error(w, "state did not match", http.StatusBadRequest)
		return
	}

	oauth2Token, err := s.auth.Exchange(s.context, r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}
	idToken, err := s.verifier.Verify(s.context, rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{oauth2Token, new(json.RawMessage)}

	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u, _ := url.Parse(s.Host + "/splash")
	q := u.Query()
	q.Add("token", resp.OAuth2Token.AccessToken)
	u.RawQuery = q.Encode()

	http.Redirect(w, r, u.String(), http.StatusPermanentRedirect)
}
