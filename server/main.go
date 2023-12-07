package main

import (
	"IoT-backend/server/configManager"
	"fmt"
)

func dummyUse(dummy interface{}) {}

func main() {
	fmt.Printf("\n")
	/* Read configuration from file */
	env := configManager.GetEnv()
	dummyUse(env)

	/* Init Data channel */
	// var channelMap = make(map[int]chan int)

	/* Setup SensorManager */
	// sensorServer := sensorManager.Setup(env, channelMap)

	/* Setup UserManager */
	// userServer := userManager.Setup(env, channelMap)

	/* Run */
	// go sensorServer.Run()
	// go userServer.Run()
}
