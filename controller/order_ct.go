package controller

import (
	"home-work-z-api/model/vo"
	"home-work-z-api/service"
	"home-work-z-api/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary getList
// @Tags trade
// @Produce  json
// @param starttime query string false "開始時間 (yyyy-MM-dd)" default(2001-01-01)
// @param endtime query string false "結束時間 (yyyy-MM-dd)" default(2001-01-02)
// @param top query int false "取前幾筆" default(10)
// @Success 200 {object} []vo.OrderVO
// @Router /order [get]
func OrderGetList(ctx *gin.Context) {

	starttime := ctx.Query("starttime")
	endtime := ctx.Query("endtime")
	top := ctx.Query("top")

	util.Success(ctx, "json", service.GetOrderList(starttime, endtime, top))
}

// @Summary getOne
// @Tags trade
// @Produce  json
// @param id path int true "編號"
// @Success 200 {object} vo.OrderVO
// @Router /order/{id} [get]
func OrderGet(ctx *gin.Context) {

	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	util.Success(ctx, "json", service.GetOrder(id))
}

// @Summary addOne
// @Description 新增訂單
// @Tags trade
// @Accept mpfd
// @Produce json
// @param postVO formData vo.OrderPostVO true "formData for OrderPostVO content"
// @Success 200 {object} vo.OrderVO
// @Failure 400 {string} string
// @router /order [post]
func OrderPost(ctx *gin.Context) {

	var oBody vo.OrderPostVO

	if err := ctx.ShouldBind(&oBody); err != nil {
		util.Failure(ctx, "string", http.StatusBadRequest, err.Error())
		return
	}

	order, err := service.PostOrder(oBody)

	if err != nil {
		util.Failure(ctx, "string", http.StatusBadRequest, err.Error())
		return
	}

	util.Success(ctx, "json", order)
}
