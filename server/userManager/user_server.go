package userManager

import (
	"IoT-backend/server/configManager"
	"IoT-backend/server/dataChannel"
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
