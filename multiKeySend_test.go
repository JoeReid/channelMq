package channelMq_test

import (
	"github.com/JoeReid/channelMq"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestMultiKeyUnsubscribe(t *testing.T) {
	keys := []string{"key1", "key2", "key3"}
	msg := []byte("SomeData")

	mq := getFreshMQ()
	mq.EmptyKeyGets(channelMq.ALL_MESSAGES)

	ch := mq.Subscribe(keys)
	mq.Unsubscribe(ch)

	mq.Send(msg, keys)
	for range keys {
		b, _ := chanReceives(ch, msg)
		assert.False(t, b)
	}
}

func TestMultiKeyMessage(t *testing.T) {
	keys := []string{"key1", "key2", "key3"}
	msg := []byte("SomeData")

	mq := getFreshMQ()
	mq.EmptyKeyGets(channelMq.ALL_MESSAGES)

	ch := mq.Subscribe(keys)

	// test only subscribed keys
	mq.Send(msg, keys)
	for range keys {
		if b, data := chanReceives(ch, msg); b {
			assert.Equal(t, data, msg)
		} else {
			assert.Fail(t, "chan didnt receive")
		}
	}

	// add some keys we are not subscribed to
	keys = append(
		keys,
		[]string{
			"key4", "key5", "key6", "key7",
		}...,
	)

	// shuffle the key list
	for i, _ := range keys {
		j := rand.Intn(i + 1)
		keys[i], keys[j] = keys[j], keys[i]
	}

	// test sub and non sub keys
	mq.Send(msg, keys)

	// Expect to receive 3 times
	for i := 0; i < 3; i++ {
		if b, data := chanReceives(ch, msg); b {
			assert.Equal(t, data, msg)
		} else {
			assert.Fail(t, "chan didnt receive")
		}
	}

	// then no more
	for i := 3; i < len(keys); i++ {
		b, _ := chanReceives(ch, msg)
		assert.False(t, b)
	}
}
