package server

import (
	"net/http"
	"splashserver/components"
)

// Serve the main webserver
func Serve(splash components.SplashService) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, splash.Host+"/splash", http.StatusFound)
	})
	http.HandleFunc("/splash", splash.MainSplashHandler)
	http.HandleFunc(splash.CallbackPath, splash.CallbackHandler)

	http.ListenAndServe(":"+splash.Port, nil)
}
