package main

import (
	"os"
	"strings"

	toolserver "github.com/JDinABox/tool-server"
	"github.com/jdinabox/go-await"
)

// main is the entry point of the tool-server application.
// It initializes the configuration and starts the server.
func main() {
	conf := &toolserver.Config{
		Listen: strings.TrimSpace(os.Getenv("LISTEN")),
	}
	conf.Default()
	toolserver.StartAwaitInterupt(conf, await.NewInterrupt())
}
