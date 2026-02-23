package router

import (
	"github.com/gofiber/fiber/v2"
)

func Router(route fiber.Router) {
	UserRouter(route)
}