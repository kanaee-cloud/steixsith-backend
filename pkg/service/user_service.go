package service

import (
	"errors"
	"steixsith/pkg/model"
)


func GetUserLogic(id int) (model.User, error) {
	if id != 1 {
		return model.User{}, errors.New("user tidak ditemukan di database")
	}

	user := model.User{
		ID:    1,
		Name:  "Budi Santoso",
		Email: "budi@example.com",
	}

	return user, nil
}


func GetUsersWithPagination() ([]model.User, int) {
	users := []model.User{
		{ID: 1, Name: "Budi", Email: "budi@example.com"},
		{ID: 2, Name: "Andi", Email: "andi@example.com"},
	}
	totalData := 20 

	return users, totalData
}