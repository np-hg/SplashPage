package main

import (
	"splashserver/components"
	"splashserver/server"
)

func main() {

	splash := components.NewSplash("./values.yaml")

	server.Serve(splash)
}
