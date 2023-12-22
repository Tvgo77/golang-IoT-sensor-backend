package usecase

import (
	"context"
	"time"

	"IoT-backend/server/userManager/domain"
)

type requestSensorUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewRequestSensorUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) domain.RequestSensorUsecase {
	return &requestSensorUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}

func (rsu *requestSensorUsecase) AddOneTimeToken(c context.Context, id string, token string) error {
	ctx, cancel := context.WithTimeout(c, rsu.contextTimeout)
	defer cancel()

	err := rsu.userRepository.AddOneTimeToken(ctx, id, token)
	if err != nil {
		return err
	}
	return nil
}
