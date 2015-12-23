package channelMq_test

import (
	"github.com/SuperLimitBreak/channelMq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoKeyMultiSendAll(t *testing.T) {
	keys := []string{"key1", "key2", "key3"}
	msg := []byte("SomeData")

	mq := getFreshMQ()
	mq.EmptyKeyGets(channelMq.ALL_MESSAGES)

	ch := mq.Subscribe([]string{})

	// send to blank key
	mq.Send(msg, keys)

	// Expect to recive once
	if b, data := chanReceives(ch, msg); b {
		assert.Equal(t, data, msg)
	} else {
		assert.Fail(t, "chan didnt receive")
	}

	// Dont expect any more
	for i := 1; i < len(keys); i++ {
		b, _ := chanReceives(ch, msg)
		assert.False(t, b)
	}
}

func TestNoKeyMultiSendAllNone(t *testing.T) {
	keys := []string{"key1", "key2", "key3"}
	msg := []byte("SomeData")

	mq := getFreshMQ()
	mq.EmptyKeyGets(channelMq.NO_MESSAGES)

	ch := mq.Subscribe([]string{})

	// send to blank key
	mq.Send(msg, keys)

	// Expect no messages
	for range keys {
		b, _ := chanReceives(ch, msg)
		assert.False(t, b)
	}
}
