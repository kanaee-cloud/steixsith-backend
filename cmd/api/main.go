package main

import (
	"github.com/gofiber/fiber/v2"
	// import package internal...
)

func main() {
	app := fiber.New()
	

	
	app.Listen(":3000")
}