package rest

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

// SetupHandler bundles all http rest handlers and returns it as an http.Handler
func SetupHandler() http.Handler {
	app := fiber.New()

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/healthcheck", healthcheck)

	return adaptor.FiberApp(app)
}

func healthcheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Service is healthy",
	})
}
