package connHandler

import (
	"IoT-backend/server/dataChannel"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func Handler(conn net.Conn, channelMap *dataChannel.ChannelMap) {
	defer conn.Close()

	/* Read first line to get the serialNum and verify identity */
	/* Create new entry in channel map */
	buf := make([]byte, 10)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		fmt.Printf("Error verifying from IP %s\n", conn.RemoteAddr().String())
		return
	}
	serialNum := string(buf)
	ch := make(chan int32, 1)
	channelMap.Insert(serialNum, ch)

	/* Read following lines and send data to channel */
	buf = make([]byte, 4)
	for {
		_, err := io.ReadFull(conn, buf)
		if err != nil {
			fmt.Printf("Bad data format from IP %s\n", conn.RemoteAddr().String())
			channelMap.Delete(serialNum)
		}
		sensorVal := int32(binary.BigEndian.Uint32(buf))
		sendLatest(sensorVal, ch)
	}
}

func sendLatest(val int32, ch chan int32) {
	select {
	case ch <- val: // Able to send to channel. Just send.
	default:
		// Channel not ready to receive
		select {
		case <-ch: // Discard previous val
		default: // Do nothing if channel already empty
		}
		ch <- val
	}
}
