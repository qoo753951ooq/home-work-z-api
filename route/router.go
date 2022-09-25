package route

import (
	"context"
	"home-work-z-api/controller"
	"home-work-z-api/docs"
	"home-work-z-api/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
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
