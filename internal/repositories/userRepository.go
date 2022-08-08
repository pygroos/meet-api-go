package repositories

import (
	model "meet-api/internal/models"
	. "meet-api/internal/pkg/database"
)

type UserRepository struct {
}

func (u *UserRepository) CreateUser(userInfo model.Users) bool {
	if err := DbHandle.Create(&userInfo).Error; err != nil {
		return false
	}
	return true
}

func (u *UserRepository) GetUserInfoByUserId(userId int) model.Users {
	var m = model.Users{}
	DbHandle.Where("user_id=(?)", []int{userId}).Find(&m)
	return m
}
