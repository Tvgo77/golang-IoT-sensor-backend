package domain

import "context"

type RequestSensorResponse struct {
	OneTimeToken string `json:"oneTimeToken"`
}

type RequestSensorUsecase interface {
	AddOneTimeToken(c context.Context, id string, token string) error
}
