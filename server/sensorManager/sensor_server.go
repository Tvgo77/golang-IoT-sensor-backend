package sensorManager

import (
	"IoT-backend/server/configManager"
	"IoT-backend/server/dataChannel"
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
