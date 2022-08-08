package services

import (
	model "meet-api/internal/models"
	"meet-api/internal/repositories"
)

type UserService struct {
}

func (u *UserService) Login(code string) bool {
	userRep := repositories.UserRepository{}
	var userInfo model.Users
	userInfo.Nickname = "摩遇"
	userInfo.AvatarUrl = "https://thirdwx.qlogo.cn/mmopen/vi_32/Q3auHgzwzM5UickrDSgpXBMmjxTic97skZyeq5NxmV0tTLKdhFRiaPXicGlswdBcxkN5nERiaVuQTSbia17VNEgAtQqw/132"
	userInfo.Openid = "o4aFK5GxfcxxyPxz7nmiPjT21Bk1"
	userInfo.Token = "bmzoRthoRmkese1tR3hmY3h4eVB4ejdubWlQalQyMUJrMHwxNjYwNTg3Mzcw"
	return userRep.CreateUser(userInfo)
}

func (u *UserService) Info(userId int) model.Users {
	userRep := repositories.UserRepository{}
	userInfo := userRep.GetUserInfoByUserId(userId)

	return userInfo
}
