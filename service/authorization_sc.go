package service

import (
	"home-work-z-api/model/vo"
	"home-work-z-api/service/authorization"
)

func PostUserLogin(body vo.UserPostVO) (string, error) {
	token, err := authorization.UserLogin(body)
	return token, err
}
