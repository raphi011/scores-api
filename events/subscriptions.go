package events

import "sync"

type subscriptions struct {
	mutex     sync.Mutex
	listeners []chan<- Event
}

func (s *subscriptions) Add(listener chan<- Event) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.listeners = append(s.listeners, listener)
}

func (s *subscriptions) Remove(listener chan<- Event) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, l := range s.listeners {
		if l == listener {
			s.listeners = append(s.listeners[:i], s.listeners[i+1:]...)
		}
	}
}
