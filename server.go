package toolserver

import (
	"github.com/jdinabox/go-await"
	jsoniter "github.com/json-iterator/go"
	"k8s.io/klog/v2"
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
	app := newApp(conf)

	logFatal(app.Listen(conf.Listen))
}

func StartAwaitInterupt(conf *Config, ai *await.Interrupt) {
	app := newApp(conf)
	// Wait for interupt and shutdown
	go func() {
		ai.Await()
		klog.Info("Stopping fiber")
		app.Shutdown()
	}()

	logFatal(app.Listen(conf.Listen))
	ai.Done()
}

// logFatal
func logFatal(err error) {
	if err != nil {
		klog.Fatal(err)
	}
}
