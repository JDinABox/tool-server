package toolserver

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Config struct {
	// Listen addr:port
	//
	// Default: 127.0.0.1:80
	Listen string
}

// Default values if empty
func (c *Config) Default() {
	if c.Listen == "" {
		c.Listen = "127.0.0.1:80"
	}
}

func Start(conf *Config) {
	mux := newApp()

	http.ListenAndServe(conf.Listen, mux)
}
