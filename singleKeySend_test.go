package channelMq_test

import (
	"github.com/SuperLimitBreak/channelMq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleKeyUnsubscribe(t *testing.T) {
	keys := []string{"key1"}
	msg := []byte("SomeData")

	mq := getFreshMQ()
	mq.EmptyKeyGets(channelMq.ALL_MESSAGES)

	ch := mq.Subscribe(keys)
	mq.Unsubscribe(ch)

	mq.Send(msg, keys)
	b, _ := chanReceives(ch, msg)
	assert.False(t, b)
}

func TestSingleKeyMessage(t *testing.T) {
	keys := []string{"key1"}
	msg := []byte("SomeData")

	mq := getFreshMQ()
	mq.EmptyKeyGets(channelMq.ALL_MESSAGES)

	ch := mq.Subscribe(keys)

	// send to correct key
	mq.Send(msg, keys)
	if b, data := chanReceives(ch, msg); b {
		assert.Equal(t, data, msg)
	} else {
		assert.Fail(t, "chan didnt receive")
	}

	// send to incorrect key
	mq.Send(msg, []string{"key2"})
	b, _ := chanReceives(ch, msg)
	assert.False(t, b)

	// send to both keys
	mq.Send(msg, []string{"key1", "key2"})

	// Expect to receive on key1
	if b, data := chanReceives(ch, msg); b {
		assert.Equal(t, data, msg)
	} else {
		assert.Fail(t, "chan didnt receive")
	}

	// But not key2
	b, _ = chanReceives(ch, msg)
	assert.False(t, b)

}
