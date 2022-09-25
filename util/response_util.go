package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, returnType string, data interface{}) {
	response(ctx, http.StatusOK, returnType, data)
}

func Failure(ctx *gin.Context, returnType string, statusCode int, data interface{}) {
	response(ctx, statusCode, returnType, data)
}

func response(ctx *gin.Context, code int, returnType string, data interface{}) {
	switch returnType {
	case "string":
		ctx.String(code, data.(string))
	case "json":
		ctx.JSON(code, data)
	}
}
