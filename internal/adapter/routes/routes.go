package routes

import (
	"realtz-notification-service/docs"
	handler "realtz-notification-service/internal/adapter/http-handler"
	"realtz-notification-service/internal/core/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(handler handler.HttpHandler) *gin.Engine {
	//Swagger meta data
	docs.SwaggerInfo.Title = "Realtz User Service"
	docs.SwaggerInfo.Description = "User microservice for realtz app"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/notifications"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.Use(middlewares.CORSMiddleware)

	notificationsApiAuthGroup := router.Group("/api/notifications/auth")
	notificationsApiAuthGroup.Use(middlewares.JWTMiddleware)
	{
		notificationsApiAuthGroup.GET("/get-notifications/:amount/:page_no", handler.GetNotifications)
	}

	// for swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
