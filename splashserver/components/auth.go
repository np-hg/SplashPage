package components

import (
	"context"
	"os"
	"strings"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

// AuthCredentials return auth object
type AuthCredentials struct {
	ClientID     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	RedirectURL  string   `yaml:"callback_url"`
	ProviderURL  string   `yaml:"provider_url"`
	Scopes       []string `yaml:"scopes,omitempty"`
	State        string   `yaml:"state"`
}

// ToOAuth returns the oauth struct from the yaml struct
func (a *AuthCredentials) ToOAuth(ctx context.Context, provider oidc.Provider) oauth2.Config {
	var ClientSecret string
	if strings.TrimSpace(a.ClientSecret) == "" {
		ClientSecret = os.Getenv("OAUTH_CLIENT_SECRET")
	} else {
		ClientSecret = a.ClientSecret
	}
	return oauth2.Config{
		ClientID:     a.ClientID,
		ClientSecret: ClientSecret,
		RedirectURL:  a.RedirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       append(a.Scopes, oidc.ScopeOpenID),
	}
}
