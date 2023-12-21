package usecase

import (
	"context"
	"time"

	"IoT-backend/server/userManager/domain"
)

type updateSensorUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUpdateSensorUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) domain.UpdateSensorUsecase {
	return &updateSensorUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}

func (usu *updateSensorUsecase) AddSensor(c context.Context, id string, serialNum string) error {
	ctx, cancel := context.WithTimeout(c, usu.contextTimeout)
	defer cancel()

	err := usu.userRepository.AddSensor(ctx, id, serialNum)
	if err != nil {
		return err
	}
	return nil
}

func (usu *updateSensorUsecase) RemoveSensor(c context.Context, id string, serialNum string) error {
	ctx, cancel := context.WithTimeout(c, usu.contextTimeout)
	defer cancel()

	err := usu.userRepository.RemoveSensor(ctx, id, serialNum)
	if err != nil {
		return err
	}
	return nil
}
