package route

import (
	"home-work-z-api/controller"
	"home-work-z-api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	apiGroup := r.Group(docs.SwaggerInfo.BasePath)
	apiGroup.Static("/static", "./static")
	setInformationRouter(apiGroup)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}

func setInformationRouter(apiGroup *gin.RouterGroup) {

	orderRouter := apiGroup.Group("/order")
	{
		orderRouter.GET("", controller.OrderGetList)
		orderRouter.GET("/:id", controller.OrderGet)
		orderRouter.POST("", controller.OrderPost)
	}
}
