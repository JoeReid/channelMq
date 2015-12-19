package channelMq_test

import (
	"github.com/JoeReid/channelMq"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func getFreshMQ() channelMq.MQ {
	return channelMq.MQ{}
}

func assertChanReceives(t *testing.T, ch chan []byte, msg []byte) {
	select {
	case <-time.After(20 * time.Millisecond):
		assert.Fail(t, "chan receive timeout")

	case b := <-ch:
		assert.Equal(t, b, msg)
	}
}

func assertNotChanReceives(t *testing.T, ch chan []byte) {
	select {
	case <-time.After(20 * time.Millisecond):
		return
	case <-ch:
		assert.Fail(t, "chan receive timeout")
	}
}

func TestSubscribeNoKey(t *testing.T) {
	keys := []string{}
	msg := []byte("SomeData")

	// ------------------------------------------
	// Test for all messages
	// ------------------------------------------
	mq := getFreshMQ()
	mq.EmptyKeyGets(channelMq.ALL_MESSAGES)

	ch, err := mq.Subscribe(keys)
	assert.Nil(t, err)

	// send to blank key
	mq.Send(msg, keys)
	assertChanReceives(t, ch, msg)

	mq.Send(msg, []string{"ActualKey"})
	assertChanReceives(t, ch, msg)

	// ------------------------------------------
	// Test for no messages
	// ------------------------------------------
	mq = getFreshMQ()
	mq.EmptyKeyGets(channelMq.NO_MESSAGES)

	ch, err = mq.Subscribe(keys)
	assert.Nil(t, err)

	// send to blank key
	mq.Send(msg, keys)
	assertNotChanReceives(t, ch)

	mq.Send(msg, []string{"ActualKey"})
	assertNotChanReceives(t, ch)
}
