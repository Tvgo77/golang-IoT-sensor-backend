package domain

import "context"

type UpdateSensorRequest struct {
	SerialNum string `form:"serial" binding:"required"`
	Operation string `form:"operation" binding:"required"`
}

type UpdateSensorResponse struct {
	Success string `json:"status"`
}

type UpdateSensorUsecase interface {
	AddSensor(c context.Context, id string, serialNum string) error
	RemoveSensor(c context.Context, id string, serialNum string) error
}
