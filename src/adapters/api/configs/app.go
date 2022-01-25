package configs

import (
	"go-viacep-wrapper/adapters/api/handlers"
	"go-viacep-wrapper/internal/application/gateway/viacep"
	"go-viacep-wrapper/internal/application/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func App() *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(logger.New())

	viacepgateway := viacep.NewGateway(http.DefaultClient, "https://viacep.com.br")
	service := services.NewProfileService(viacepgateway)
	handler := handlers.NewHandler(service)

	api := app.Group("/api")
	handler.SetUpRoutes(api)

	return app
}
