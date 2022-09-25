package route

import (
	"home-work-z-api/controller"
	"home-work-z-api/docs"
	"home-work-z-api/middleware"

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

	authRouter := apiGroup.Group("/user")
	{
		authRouter.POST("/login", controller.UserLoginPost)
	}

	orderRouter := apiGroup.Group("/order").Use(middleware.JWTAuth())
	{
		orderRouter.GET("", controller.OrderGetList)
		orderRouter.GET("/:id", controller.OrderGet)
		orderRouter.POST("", controller.OrderPost)
		orderRouter.PUT("/:id", controller.OrderPut)
		orderRouter.DELETE("/:id", controller.OrderDelete)
	}
}
