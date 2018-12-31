package events

import "sync"

type subscriptions struct {
	mutex    sync.Mutex
	handlers []EventHandler
}

func (s *subscriptions) Add(handler EventHandler) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.handlers = append(s.handlers, handler)
}
