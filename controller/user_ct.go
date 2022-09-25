package controller

import (
	"home-work-z-api/model/vo"
	"home-work-z-api/service"
	"home-work-z-api/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary login
// @Description 登入
// @Tags authorization
// @Accept json
// @Produce plain
// @param postVO body vo.UserPostVO true "body for UserPostVO content"
// @Success 200 {string} string "Token"
// @Failure 400 {string} string
// @router /user/login [post]
func UserLoginPost(ctx *gin.Context) {

	var uBody vo.UserPostVO

	if err := ctx.ShouldBindJSON(&uBody); err != nil {
		util.Failure(ctx, "string", http.StatusBadRequest, err.Error())
		return
	}

	token, err := service.PostUserLogin(uBody)

	if err != nil {
		util.Failure(ctx, "string", http.StatusBadRequest, err.Error())
		return
	}

	util.Success(ctx, "json", token)
}
