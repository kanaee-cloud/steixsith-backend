package router

import (
	"steixsith/pkg/handler"

	"github.com/gofiber/fiber/v2"
)


func UserRouter(router fiber.Router) {
	router.Get("/user", handler.GetUserHandler)
}