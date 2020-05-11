package events

import (
	"fmt"
	"strings"
	"sync"
)

// Broker handles subscribing and publishing of events
type Broker struct {
	subscribers sync.Map
}

// Publish publishes an event to all subscribers that are listening to an event
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
			panic(fmt.Sprintf("incompatible subscription for event %s", event.Name))
		}

		subscriptions.mutex.Lock()

		for _, listener := range subscriptions.listeners {
			listener <- event
		}

		subscriptions.mutex.Unlock()
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

// Subscribe subscribes to an event in the form of "some/interesting/event" or
// (with wildcards) "some/interesting/*". Wildcards are only valid after directly
// preceding slashes.
func (b *Broker) Subscribe(eventName string) (<-chan Event, Unsubscribe) {
	var subs *subscriptions

	messageChan := make(chan Event)

	if s, found := b.subscribers.Load(eventName); found {
		subs = s.(*subscriptions)
	} else {
		subs = &subscriptions{}
		b.subscribers.Store(eventName, subs)
	}

	subs.Add(messageChan)

	return messageChan, func() {
		subs.Remove(messageChan)
		close(messageChan)
	}
}
