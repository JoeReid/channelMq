package channelMq

import (
	"sync"
	"time"
)

func NewGroup() *group {
	rtn := &group{
		make(map[chan []byte]struct{}),
		&sync.Mutex{},
		make(chan []byte, CHANNEL_SIZE),
	}

	go rtn.work()
	return rtn
}

type group struct {
	members     map[chan []byte]struct{}
	mu          *sync.Mutex
	messageChan chan []byte
}

func (g *group) work() {
	for {
		msg := <-g.messageChan

		g.mu.Lock()

		for ch, _ := range g.members {
			go func() {
				select {
				case ch <- msg:
					return
				case <-time.After(200 * time.Millisecond):
					return
				}
			}()
		}

		g.mu.Unlock()
	}
}

func (g *group) send(message []byte) {
	g.messageChan <- message
}

func (g *group) add(ch chan []byte) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	_, exists := g.members[ch]
	if exists {
		return AlreadySubscribedError
	}

	g.members[ch] = struct{}{}
	return nil
}

func (g *group) remove(ch chan []byte) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	_, exists := g.members[ch]
	if !exists {
		return NotSubscribedError
	}

	delete(g.members, ch)
	return nil
}
