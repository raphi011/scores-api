package events

type Event struct {
	Name string      `json:"name"`
	Body interface{} `json:"body"`
}

type EventHandler func(event Event)

type Publisher interface {
	Publish(event Event)
}

type Subscriber interface {
	Subscribe(eventName string, handler EventHandler)
}

type PublisherSubscriber interface {
	Publisher
	Subscriber
}
