package events

// Event contains the actual event (Body) and meta information
// like the name of the Event.
type Event struct {
	Name string      `json:"name"`
	Body interface{} `json:"body"`
}

// Unsubscribe allows unsubscribing of an Event.
type Unsubscribe func()

// Publisher can publish Events.
type Publisher interface {
	Publish(event Event)
}

// Subscriber can subscribe to events.
type Subscriber interface {
	Subscribe(eventName string) (<-chan Event, Unsubscribe)
}

// PublisherSubscriber implements both Publisher and Subscriber.
type PublisherSubscriber interface {
	Publisher
	Subscriber
}
