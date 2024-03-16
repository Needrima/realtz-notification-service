package handler

import (
	"realtz-notification-service/internal/core/domain/dto"
	errorHelper "realtz-notification-service/internal/core/helpers/error-helper"
	logHelper "realtz-notification-service/internal/core/helpers/log-helper"
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

func (h HttpHandler) SendNotification(c *gin.Context) {
	body := dto.SendNotificationDto{}
	if err := c.BindJSON(&body); err != nil {
		logHelper.LogEvent(logHelper.ErrorLog, "binding signup request body: "+err.Error())
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	response, err := h.httpPort.SendNotification(body)
	if err != nil {
		c.AbortWithStatusJSON(err.(errorHelper.ServiceError).Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response)
}
