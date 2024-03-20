package handler

import (
	// "realtz-notification-service/internal/core/domain/dto"
	errorHelper "realtz-notification-service/internal/core/helpers/error-helper"
	// logHelper "realtz-notification-service/internal/core/helpers/log-helper"
	tokenHelper "realtz-notification-service/internal/core/helpers/token-helper"
	"realtz-notification-service/internal/ports"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	httpPort ports.HTTPPort
}

func NewHTTPHandler(httpPort ports.HTTPPort) HttpHandler {
	return HttpHandler{
		httpPort: httpPort,
	}
}

// func (h HttpHandler) SendNotification(c *gin.Context) {
// 	body := dto.SendNotificationDto{}
// 	if err := c.BindJSON(&body); err != nil {
// 		logHelper.LogEvent(logHelper.ErrorLog, "binding signup request body: "+err.Error())
// 		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request body: " + err.Error()})
// 		return
// 	}

// 	response, err := h.httpPort.SendNotification(body)
// 	if err != nil {
// 		c.AbortWithStatusJSON(err.(errorHelper.ServiceError).Code, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, response)
// }

// @Summary Get Notifications
// @Description Get first n notificatiobs n is the amount parameter
// @Tags Notification
// @Accept json
// @Produce json
// @Param Token header string true "Authentication token"
// @Param amount path string true "amount of notifications to be queried"
// @Param page_no path string true "page to be gotten e.g if amount param is 8, 1 means first 8 products, 2 means from 9th to 16th notification"
// @Success 200 {object} interface{} "Successfully retrieved notifications"
// @Failure 500 {object} errorHelper.ServiceError "something went wrong"
// @Router /auth/get-notifications/{amount}/{page_no} [get]
func (h HttpHandler) GetNotifications(c *gin.Context) {
	amount := c.Param("amount")
	pageNo := c.Param("page_no")

	// get user from jwt-token
	currentUser, _ := tokenHelper.ValidateToken(c.GetHeader("Token"))

	response, err := h.httpPort.GetNotifications(c.Request.Context(), *currentUser, amount, pageNo)
	if err != nil {
		c.AbortWithStatusJSON(err.(errorHelper.ServiceError).Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response)
}
