package main

import (
	"net/http"
	"splash/components"
)

func main() {

	splash := components.NewSplash("./values.yaml")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, splash.Host+"/splash", http.StatusFound)
	})
	http.HandleFunc("/splash", splash.MainSplashHandler)
	http.HandleFunc(splash.CallbackPath, splash.CallbackHandler)

	http.ListenAndServe(":"+splash.Port, nil)
}
