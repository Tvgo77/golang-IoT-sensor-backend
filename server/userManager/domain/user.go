package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "Users"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Email        string             `bson:"email"`
	Password     string             `bson:"password"`
	Sensors      []string           `bson:"sensors"`
	OneTimeToken string             `bson:"oneTimeToken"`
}

type UserToken struct {
	OneTimeToken string `bson:"oneTimeToken"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
	AddSensor(c context.Context, id string, serialNum string) error
	RemoveSensor(c context.Context, id string, serialNum string) error
	AddOneTimeToken(c context.Context, id string, token string) error
	GetTokenByID(c context.Context, id string) (string, error)
}
