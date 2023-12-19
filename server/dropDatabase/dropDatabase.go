package main

import (
	"IoT-backend/server/configManager"
	"IoT-backend/server/userManager/mongo"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	env := configManager.GetEnv("../.env")
	client := newMongoDatabase(env)
	db := client.Database(env.DBName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := db.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func newMongoDatabase(env *configManager.Env) mongo.Client {
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
