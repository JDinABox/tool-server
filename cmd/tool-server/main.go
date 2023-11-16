package main

import (
	"os"
	"strings"

	toolserver "github.com/JDinABox/tool-server"
)

// main is the entry point of the tool-server application.
// It initializes the configuration and starts the server.
func main() {
	conf := &toolserver.Config{
		Listen: strings.TrimSpace(os.Getenv("LISTEN")),
	}
	conf.Default()
	toolserver.Start(conf)
}
