package toolserver

import "github.com/gofiber/fiber/v2"

// newApp init fiber app
func newApp(conf *Config) *fiber.App {
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

	return app
}
