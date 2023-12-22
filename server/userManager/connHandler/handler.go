package connHandler

import (
	"IoT-backend/server/configManager"
	"IoT-backend/server/dataChannel"
	"IoT-backend/server/userManager/domain"
	"IoT-backend/server/userManager/mongo"
	"IoT-backend/server/userManager/repository"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type sensorNchan struct {
	sensor     string
	channelVal int32
	channel    chan int32
}

func Handler(conn net.Conn, channelMap *dataChannel.ChannelMap, env *configManager.Env, db mongo.Database) {
	defer conn.Close()
	ur := repository.NewUserRepository(db, domain.CollectionUser)

	// Verify if userID(24 bytes hex string) match its token(6 bytes digital string)
	buf := make([]byte, 30)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		fmt.Printf("Error verifying from IP %s\n", conn.RemoteAddr().String())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(env.ContextTimeout))
	defer cancel()

	userID := string(buf[0:24])
	token := string(buf[24:])
	user, err := ur.GetByID(ctx, userID)
	if err != nil {
		log.Fatal(err)
		return
	}

	if token != user.OneTimeToken {
		fmt.Printf("Error verifying one time token from IP %s\n", conn.RemoteAddr().String())
		return
	}

	// Get all channel of user's sensor
	var channels []sensorNchan
	for _, sensor := range user.Sensors {
		channels = append(channels, sensorNchan{sensor, 0, channelMap.GetChannel(sensor)})
	}

	// Send back all sensor data from channels
	for {
		for _, pair := range channels {
			io.WriteString(conn, pair.sensor)
			writeBuf := make([]byte, 4)

			// Key Step of goroutine communication
			select {
			case <-pair.channel:
				pair.channelVal = <-pair.channel
			default:
			}

			binary.BigEndian.PutUint32(writeBuf, uint32(pair.channelVal))
			_, err = conn.Write(writeBuf)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
		time.Sleep(time.Duration(env.SendInterval))
	}
}
