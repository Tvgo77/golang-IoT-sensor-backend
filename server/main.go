package main

import (
	"IoT-backend/server/configManager"
	"IoT-backend/server/dataChannel"
	"IoT-backend/server/sensorManager"
	"IoT-backend/server/userManager"
	"fmt"
)

func dummyUse(dummy interface{}) {}

func main() {
	fmt.Printf("\n")
	/* Read configuration from file */
	env := configManager.GetEnv(".env")
	dummyUse(env)

	/* Init Data channel */
	var channelMap = dataChannel.NewChannelMap()
	dummyUse(channelMap)

	/* Setup SensorManager */
	sensorServer := sensorManager.Setup(env, channelMap)
	dummyUse(sensorServer)

	/* Setup UserManager */
	userServer := userManager.Setup(env, channelMap)

	/* Run */
	// go sensorServer.Run()
	userServer.Run()
}
