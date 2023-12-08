package sensorManager

import (
	"IoT-backend/server/configManager"
	"IoT-backend/server/dataChannel"
	"IoT-backend/server/sensorManager/connHandler"
	"fmt"
	"net"
	"os"
)

type sensorServer struct {
	env        *configManager.Env
	channelMap *dataChannel.ChannelMap
}

func Setup(ev *configManager.Env, chMap *dataChannel.ChannelMap) *sensorServer {
	return &sensorServer{
		env:        ev,
		channelMap: chMap,
	}
}

/* Start Running */
func (s *sensorServer) Run() {
	/* Start Listening on port */
	port := s.env.SensorPort
	fmt.Printf("Sensor Server Start Listening on port: %d\n", port)
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
		go connHandler.Handler(conn, s.channelMap)
	}
}
