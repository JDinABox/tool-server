package main

import (
	"os"
	"strings"

	toolserver "github.com/JDinABox/tool-server"
)

// main is the entry point of the tool-server application.
// It initializes the configuration and starts the server.
func main() {
	listenAddr := strings.TrimSpace(os.Getenv("LISTEN"))
	toolserver.Start(toolserver.WithListenAddr(listenAddr))
}
