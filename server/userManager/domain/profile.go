package domain

import "context"

type Profile struct {
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Sensors []string `json:"sensors"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*Profile, error)
}
