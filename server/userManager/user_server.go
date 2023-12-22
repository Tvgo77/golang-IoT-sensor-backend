package userManager

import (
	"IoT-backend/server/configManager"
	"IoT-backend/server/dataChannel"
	"IoT-backend/server/userManager/api/route"
	"IoT-backend/server/userManager/connHandler"
	"IoT-backend/server/userManager/mongo"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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

	timeout := time.Duration(s.env.ContextTimeout) * time.Second
	gin := gin.Default()
	/* Router setup */
	route.Setup(s.env, timeout, db, gin)

	/* Run data sending server */
	go s.dataSendServer(db)

	/* Run gin engine*/
	gin.Run(fmt.Sprintf(":%d", s.env.ServerPort1))
}

func (s *userServer) dataSendServer(db mongo.Database) {
	/* Start Listening on port */
	port := s.env.ServerPort2
	fmt.Printf("User Server Start Listening on port: %d\n", port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error Listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	/* Accept connection in loop */
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting:", err.Error())
			os.Exit(1)
		}
		/* Handle connection in go routine */
		go connHandler.Handler(conn, s.channelMap, s.env, db)
	}
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
