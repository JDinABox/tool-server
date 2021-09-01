package toolserver

import (
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
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

func Start(conf Config) {
	conf.Default()
	// Init fiber app
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		BodyLimit:   1 * 1024 * 2024, // 1MB
	})

	// Api base path
	apiV1 := app.Group("/api/v1")

	// Return pong on ping
	apiV1.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON("pong")
	})

	// Return client ip
	apiV1.Get("/ip", func(c *fiber.Ctx) error {
		ips := c.IPs()
		var ip string
		if len(ips) > 0 {
			ip = ips[0]
		} else {
			ip = c.IP()
		}
		return c.JSON(fiber.Map{
			"ip": ip,
		})
	})

	// 404 error
	app.Use(func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"error":     "404",
			"not-found": c.OriginalURL(),
		})
	})

	logger.Fatal(app.Listen(conf.Listen))
}
