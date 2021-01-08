package main

import (
	"splashserver/cmd"
)

func main() {

	indexHTML, err := Asset("data/index.html")
	if err != nil {
		panic(err)
	}

	cmd.LoadAssets(string(indexHTML))
	cmd.Execute()
}
