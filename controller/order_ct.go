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
// @Description 訂單列表
// @Tags trade
// @Produce  json
// @param starttime query string false "開始時間 (yyyy-MM-dd)" default(2001-01-01)
// @param endtime query string false "結束時間 (yyyy-MM-dd)" default(2001-01-02)
// @param top query int false "取前幾筆" default(10)
// @Security BearerAuth
// @Success 200 {object} []vo.OrderVO
// @Router /order [get]
func OrderGetList(ctx *gin.Context) {

	starttime := ctx.Query("starttime")
	endtime := ctx.Query("endtime")
	top := ctx.Query("top")

	util.Success(ctx, "json", service.GetOrderList(starttime, endtime, top))
}

// @Summary getOne
// @Description 單筆訂單
// @Tags trade
// @Produce  json
// @param id path int true "編號"
// @Security BearerAuth
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
// @Security BearerAuth
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

// @Summary editOne
// @Description 編輯訂單
// @Tags trade
// @Accept json
// @Produce plain
// @param id path int true "編號"
// @param putVO body vo.OrderPutVO true "body for OrderPutVO content"
// @Security BearerAuth
// @Success 200 {string} string
// @Failure 400 {string} string
// @router /order/{id} [put]
func OrderPut(ctx *gin.Context) {

	var oBody vo.OrderPutVO

	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err = ctx.ShouldBindJSON(&oBody); err != nil {
		util.Failure(ctx, "string", http.StatusBadRequest, err.Error())
		return
	}

	result, err := service.PutOrder(id, oBody)

	if err != nil {
		util.Failure(ctx, "string", http.StatusBadRequest, err.Error())
		return
	}

	util.Success(ctx, "string", result)
}

// @Summary deleteOne
// @Description 刪除訂單
// @Tags trade
// @Produce plain
// @param id path int true "編號"
// @Security BearerAuth
// @Success 200 {string} string
// @Failure 400 {string} string
// @router /order/{id} [delete]
func OrderDelete(ctx *gin.Context) {

	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	result, err := service.DeleteOrder(id)

	if err != nil {
		util.Failure(ctx, "string", http.StatusBadRequest, err.Error())
		return
	}

	util.Success(ctx, "string", result)
}
