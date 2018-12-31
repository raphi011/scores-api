package events

import (
	"fmt"
	"strings"
	"sync"
)

type Broker struct {
	subscribers sync.Map
}

func (b *Broker) Publish(event Event) {
	if event.Name == "" {
		return
	}

	handlers := expandPossibleHandlers(event.Name)

	for _, handler := range handlers {
		s, found := b.subscribers.Load(handler)

		if !found {
			continue
		}

		subscriptions, ok := s.(*subscriptions)

		if !ok {
			panic(fmt.Sprintf("incompatible eventhandler for event %s", event.Name))
		}

		for _, subscriber := range subscriptions.handlers {
			subscriber(event)
		}
	}
}

// expandPossibleHandlers returns a list containing the eventName itself
// and possible wildcard eventHandler names.
// e.g.: expandPossibleHandlers("volleynet/sync/new-tournament") =>
// { "volleynet/*", "volleynet/sync/*", "volleynet/sync/new-tournament" }
func expandPossibleHandlers(eventName string) []string {
	parts := strings.Split(eventName, "/")

	if strings.Index(eventName, "/") == -1 {
		return []string{eventName}
	}

	handlers := []string{}

	current := ""

	for i, part := range parts {
		if i == len(parts)-1 {
			break
		}

		current += part + "/"

		handlers = append(handlers, current+"*")
	}

	handlers = append(handlers, eventName)

	return handlers
}

func (b *Broker) Subscribe(eventName string, handler EventHandler) {
	var subs *subscriptions

	if s, found := b.subscribers.Load(eventName); found {
		subs = s.(*subscriptions)
	} else {
		subs = &subscriptions{}
		b.subscribers.Store(eventName, subs)
	}

	subs.Add(handler)
}
