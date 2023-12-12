package userManager

import (
	"IoT-backend/server/configManager"
	"IoT-backend/server/dataChannel"
	"IoT-backend/server/userManager/mongo"
	"context"
	"fmt"
	"log"
	"time"
)

type userServer struct {
	env        *configManager.Env
	channelMap *dataChannel.ChannelMap
}

func Setup(ev *configManager.Env, chMap *dataChannel.ChannelMap) *userServer {
	return &userServer{
		env:        ev,
		channelMap: chMap,
	}
}

func (s *userServer) Run() {
	/* Database setup */
	client := NewMongoDatabase(s.env)
	db := client.Database(s.env.DBName)

	dummy(db)

	/* Router setup */
}

func NewMongoDatabase(env *configManager.Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}

func dummy(any interface{}) {}
