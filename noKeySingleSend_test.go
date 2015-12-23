package channelMq_test

import (
	"github.com/SuperLimitBreak/channelMq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoKeyUnsubscribe(t *testing.T) {
	keys := []string{}
	msg := []byte("SomeData")

	mq := getFreshMQ()
	mq.EmptyKeyGets(channelMq.ALL_MESSAGES)

	ch := mq.Subscribe(keys)
	mq.Unsubscribe(ch)

	mq.Send(msg, keys)
	b, _ := chanReceives(ch, msg)
	assert.False(t, b)
}

func TestNoKeyAllMessages(t *testing.T) {
	keys := []string{}
	msg := []byte("SomeData")

	mq := getFreshMQ()
	mq.EmptyKeyGets(channelMq.ALL_MESSAGES)

	ch := mq.Subscribe(keys)

	// send to blank key
	mq.Send(msg, keys)
	if b, data := chanReceives(ch, msg); b {
		assert.Equal(t, data, msg)
	} else {
		assert.Fail(t, "chan didnt receive")
	}

	mq.Send(msg, []string{"ActualKey"})
	if b, data := chanReceives(ch, msg); b {
		assert.Equal(t, data, msg)
	} else {
		assert.Fail(t, "chan didnt receive")
	}
}

func TestNoKeyNoMessages(t *testing.T) {
	keys := []string{}
	msg := []byte("SomeData")

	mq := getFreshMQ()
	mq.EmptyKeyGets(channelMq.NO_MESSAGES)

	ch := mq.Subscribe(keys)

	// send to blank key
	mq.Send(msg, keys)
	b, _ := chanReceives(ch, msg)
	assert.False(t, b)

	mq.Send(msg, []string{"ActualKey"})
	b, _ = chanReceives(ch, msg)
	assert.False(t, b)
}
