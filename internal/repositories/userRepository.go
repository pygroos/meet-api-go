package repositories

import (
	model "meet-api/internal/models"
	. "meet-api/internal/pkg/database"
)

type UserRepository struct {
}

func (u *UserRepository) CreateUser(userInfo map[string]string) bool {
	tx := DbHandle
	tx.Create(userInfo)
	return true
}

func (u *UserRepository) GetUserInfoByUserId(userId int) model.Users {
	var m = model.Users{}
	tx := DbHandle
	tx.Where("user_id=(?)", []int{userId})
	tx.Find(&m)

	return m
}
