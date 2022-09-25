package authorization

import (
	"fmt"
	"home-work-z-api/middleware"
	"home-work-z-api/model"
	"home-work-z-api/model/vo"
)

func UserLogin(data vo.UserPostVO) (string, error) {

	if data.Account != "alan" || data.Password != "home_work_z" {
		return model.Empty_string, fmt.Errorf("%s\n", model.Id_not_found_string)
	}

	token, err := middleware.GenerateToken(data.Account)

	if err != nil {
		return model.Empty_string, err
	}

	return token, nil
}
