package services

import (
	"fmt"
	model "meet-api/internal/models"
	"meet-api/internal/repositories"
)

type UserService struct {
}

func (u *UserService) Login(code string) bool {
	fmt.Printf(code)
	return true
}

func (u *UserService) Info(userId int) model.Users {
	userRep := repositories.UserRepository{}
	userInfo := userRep.GetUserInfoByUserId(userId)

	return userInfo
}
