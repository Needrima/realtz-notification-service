package routes

import (
	handler "realtz-notification-service/internal/adapter/http-handler"
	// "realtz-notification-service/internal/core/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(handler handler.HttpHandler) *gin.Engine {
	router := gin.Default()
	// router.Use(middlewares.CORSMiddleware)
	// notificationApiGroup := router.Group("/api/notification")
	// {
	// 	notificationApiGroup.POST("/send-notification", handler.SendNotification)
	// }

	// for swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
