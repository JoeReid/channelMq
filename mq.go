package channelMq

import (
	"sync"
)

func NewMQ() *MQ {
	return &MQ{
		ALL_MESSAGES,
		make(map[string]*group),
		&sync.Mutex{},
		NewGroup(),
	}
}

type MQ struct {
	messageSet       MessageSet
	subscribers      map[string]*group
	subscribersMu    *sync.Mutex
	subscribersNoKey *group
}

func (m *MQ) EmptyKeyGets(ms MessageSet) {
	m.messageSet = ms
}

func (m *MQ) Send(message []byte, keys []string) {
	if m.messageSet == ALL_MESSAGES {
		m.subscribersNoKey.send(message)
	}

	if len(keys) == 0 {
		for _, v := range m.subscribers {
			v.send(message)
		}
		return
	}

	keySet := asSet(keys)

	for key, _ := range keySet {
		g, ok := m.subscribers[key]
		if ok {
			g.send(message)
		}
	}
}

func (m *MQ) Subscribe(keys []string) chan []byte {
	rtn := make(chan []byte, CHANNEL_SIZE)

	if len(keys) == 0 {
		m.subscribersNoKey.add(rtn)
		return rtn
	}

	keySet := asSet(keys)
	for key, _ := range keySet {
		m.subscribersMu.Lock()

		_, ok := m.subscribers[key]
		if !ok {
			m.subscribers[key] = NewGroup()
		}
		m.subscribers[key].add(rtn)

		m.subscribersMu.Unlock()
	}

	return rtn
}

func (m *MQ) Unsubscribe(ch chan []byte) {
	m.subscribersNoKey.remove(ch)

	m.subscribersMu.Lock()

	for _, g := range m.subscribers {
		g.remove(ch)
	}

	m.subscribersMu.Unlock()
}

func asSet(strs []string) map[string]struct{} {
	rtn := make(map[string]struct{})

	for _, str := range strs {
		rtn[str] = struct{}{}
	}

	return rtn
}
