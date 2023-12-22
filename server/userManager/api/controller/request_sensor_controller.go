package controller

import (
	"fmt"
	"math/rand"
	"net/http"

	"IoT-backend/server/userManager/domain"

	"github.com/gin-gonic/gin"
)

type RequestSensorController struct {
	RequestSensorUsecase domain.RequestSensorUsecase
}

func (rsc *RequestSensorController) RequestSensor(c *gin.Context) {
	userID := c.GetString("x-user-id")

	// Generate random one time token
	randomNumber := rand.Intn(900000) + 100000
	token := fmt.Sprintf("%d", randomNumber)

	// Record one time token to database
	err := rsc.RequestSensorUsecase.AddOneTimeToken(c, userID, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Do response
	requestSensorResponse := domain.RequestSensorResponse{
		OneTimeToken: token,
		UserID:       userID,
	}

	c.JSON(http.StatusOK, requestSensorResponse)
}
