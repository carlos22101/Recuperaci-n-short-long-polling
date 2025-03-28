package polling

import (
	"sync"
)

type Poller struct {
	subscribers map[chan interface{}]bool
	mu         sync.RWMutex
}

func NewPoller() *Poller {
	return &Poller{
		subscribers: make(map[chan interface{}]bool),
	}
}

func (p *Poller) Subscribe() chan interface{} {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	ch := make(chan interface{}, 1)
	p.subscribers[ch] = true
	return ch
}

func (p *Poller) Unsubscribe(ch chan interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	delete(p.subscribers, ch)
}

func (p *Poller) Publish(data interface{}) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	for ch := range p.subscribers {
		select {
		case ch <- data:
		default:
			// Canal lleno, saltamos
		}
	}
}