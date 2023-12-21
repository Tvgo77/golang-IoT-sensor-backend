package controller

import (
	"net/http"

	"IoT-backend/server/userManager/domain"

	"github.com/gin-gonic/gin"
)

type UpdateSensorController struct {
	UpdateSensorUsecase domain.UpdateSensorUsecase
}

func (usc *UpdateSensorController) UpdateSensor(c *gin.Context) {
	userId := c.GetString("x-user-id")

	// Check request format
	var request domain.UpdateSensorRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Check it is Add or Remove operation and do update operation
	if request.Operation == "ADD" {
		err = usc.UpdateSensorUsecase.AddSensor(c, userId, request.SerialNum)
	} else if request.Operation == "Remove" {
		err = usc.UpdateSensorUsecase.RemoveSensor(c, userId, request.SerialNum)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	updateSensorResponse := domain.UpdateSensorResponse{
		Success: "success",
	}

	c.JSON(http.StatusOK, updateSensorResponse)
}
