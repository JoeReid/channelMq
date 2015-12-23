package channelMq_test

import (
	"github.com/SuperLimitBreak/channelMq"
	"time"
)

func getFreshMQ() *channelMq.MQ {
	return channelMq.NewMQ()
}

func chanReceives(ch chan []byte, msg []byte) (bool, []byte) {
	select {
	case <-time.After(time.Millisecond):
		return false, nil

	case b := <-ch:
		return true, b
	}
}
