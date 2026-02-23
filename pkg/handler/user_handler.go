package handler

import (
	"strconv"
	"steixsith/pkg/model"
	"steixsith/pkg/service"
	"steixsith/pkg/utils"

	"github.com/gofiber/fiber/v2"
)


func GetUserHandler(c *fiber.Ctx) error {
	idParam := c.Query("id", "1")
	id, _ := strconv.Atoi(idParam)

	data, err := service.GetUserLogic(id)
	if err != nil {
		return utils.SendError(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SendSuccess[model.User](c, fiber.StatusOK, "Berhasil mengambil data user", data)
}


func GetUsersListHandler(c *fiber.Ctx) error {
	data, total := service.GetUsersWithPagination()

	meta := model.MetaPagination{
		Page:      1,
		Limit:     10,
		TotalData: total,
		TotalPage: (total / 10), 
	}


	return utils.SendPagination[[]model.User](c, fiber.StatusOK, "Berhasil mengambil list user", data, meta)
}