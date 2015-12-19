package channelMq_test

import (
	"fmt"
	"github.com/JoeReid/channelMq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiUnsubscribe(t *testing.T) {
	numberOfSubscribers := 500
	subscribers := []chan []byte{}
	keys := [][]string{}
	msg := []byte("SomeData")

	for i := 0; i < numberOfSubscribers; i++ {
		keys = append(keys, []string{
			fmt.Sprintf("key%v", i),
		})
	}

	mq := getFreshMQ()
	mq.EmptyKeyGets(channelMq.ALL_MESSAGES)

	for i, _ := range keys {
		subscribers = append(subscribers, mq.Subscribe(keys[i]))
	}

	for _, elm := range subscribers {
		mq.Unsubscribe(elm)
	}

	for i, _ := range keys {
		mq.Send(msg, keys[i])
	}

	// None recieve
	for _, elm := range subscribers {
		b, _ := chanReceives(elm, msg)
		assert.False(t, b)
	}
}

func TestMultiSend(t *testing.T) {
	numberOfSubscribers := 500
	subscribers := []chan []byte{}
	keys := [][]string{}
	msg := []byte("SomeData")

	for i := 0; i < numberOfSubscribers; i++ {
		keys = append(keys, []string{
			fmt.Sprintf("key%v", i),
		})
	}

	mq := getFreshMQ()
	mq.EmptyKeyGets(channelMq.ALL_MESSAGES)

	for i, _ := range keys {
		subscribers = append(subscribers, mq.Subscribe(keys[i]))
	}

	for i, _ := range keys {
		mq.Send(msg, keys[i])
	}

	// Each chan receives once
	for _, elm := range subscribers {
		if b, data := chanReceives(elm, msg); b {
			assert.Equal(t, data, msg)
		} else {
			assert.Fail(t, "chan didnt receive")
		}
	}

	// and only once
	for _, elm := range subscribers {
		b, _ := chanReceives(elm, msg)
		assert.False(t, b)
	}
}
