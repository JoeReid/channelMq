package channelMq

import (
	"errors"
)

type MessageSet int

const (
	ALL_MESSAGES MessageSet = iota
	NO_MESSAGES
)

const (
	CHANNEL_SIZE = 256
)

var (
	NotSubscribedError     = errors.New("Channel not subscribed")
	AlreadySubscribedError = errors.New("Channel already subscribed")
)
