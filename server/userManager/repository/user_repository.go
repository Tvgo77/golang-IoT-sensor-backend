package repository

import (
	"IoT-backend/server/userManager/domain"
	"IoT-backend/server/userManager/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []domain.User

	err = cursor.All(c, &users)
	if users == nil {
		return []domain.User{}, err
	}

	return users, err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}

func (ur *userRepository) AddSensor(c context.Context, id string, serialNum string) error {
	collection := ur.database.Collection(ur.collection)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": idHex}
	update := bson.M{"$addToSet": bson.M{"sensors": serialNum}}
	_, err = collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}
	return err
}

func (ur *userRepository) RemoveSensor(c context.Context, id string, serialNum string) error {
	collection := ur.database.Collection(ur.collection)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": idHex}
	update := bson.M{"$pull": bson.M{"sensors": serialNum}}
	_, err = collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}
	return err
}

func (ur *userRepository) AddOneTimeToken(c context.Context, id string, token string) error {
	collection := ur.database.Collection(ur.collection)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": idHex}
	update := bson.M{"$set": bson.M{"oneTimeToken": token}}
	_, err = collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetTokenByID(c context.Context, id string) (string, error) {
	collection := ur.database.Collection(ur.collection)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	var userToken domain.UserToken
	filter := bson.M{"_id": idHex}
	projection := bson.D{{Key: "email", Value: 1}, {Key: "_id", Value: 0}} // Include email, exclude _id
	err = collection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&userToken)
	if err != nil {
		return "", err
	}

	return userToken.OneTimeToken, err
}
