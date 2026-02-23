package utils

import (
	"steixsith/pkg/model"
	"github.com/gofiber/fiber/v2"
)

// SendSuccess membungkus response sukses standar
func SendSuccess[T any](c *fiber.Ctx, code int, message string, data T) error {
	response := model.ResponseEntity[T]{
		Code:    code,
		Status:  true,
		Message: message,
		Data:    data,
	}
	return c.Status(code).JSON(response)
}


func SendPagination[T any](c *fiber.Ctx, code int, message string, data T, meta model.MetaPagination) error {
	response := model.ResponseEntity[T]{
		Code:    code,
		Status:  true,
		Message: message,
		Data:    data,
		Meta:    &meta,
	}
	return c.Status(code).JSON(response)
}

func SendError(c *fiber.Ctx, code int, message string) error {
	response := model.ResponseError[any]{
		ResponseEntity: model.ResponseEntity[any]{
			Code:    code,
			Status:  false,
			Message: message,
			Data:    nil,
		},
		Path: c.Path(), 
	}
	return c.Status(code).JSON(response)
}