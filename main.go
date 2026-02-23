package main

import (
	"log"
	"steixsith/pkg/router"
	"steixsith/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(recover.New())

	route := app.Group("/api")
	
	route.Get("/", func(c *fiber.Ctx) error {
		return utils.SendSuccess[*struct{}](c, fiber.StatusOK, "Hi From Stei X Sith Backend 💡", nil)
	})

	
	router.Router(route)

	port := "8080"
	log.Printf("API server is running on http://localhost:%s/api 💡\n", port)
	
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("http server error: %v", err)
	}
}