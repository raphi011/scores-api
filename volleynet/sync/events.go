package sync

import (
	"time"

	"github.com/google/uuid"
	"github.com/raphi011/scores/events"
)

const (
	ScrapeEventsType = "volleynet/scrape/*"
	StartScrapeEventType = "volleynet/scrape/start"
	EndScrapeEventType   = "volleynet/scrape/end"
)

type StartScrapeEvent struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"time"`
	Type      string    `json:"type"`
}

func (s *Service) publishStartScrapeEvent(scrapeType string, start time.Time) {
	s.Subscriptions.Publish(events.Event{
		Name: StartScrapeEventType,
		Body: StartScrapeEvent{
			ID:        uuid.New().String(),
			Timestamp: start,
			Type:      scrapeType,
		},
	})
}

type EndScrapeEvent struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"time"`
	Report    *Changes  `json:"report"`
}

func (s *Service) publishEndScrapeEvent(report *Changes, end time.Time) {
	s.Subscriptions.Publish(events.Event{
		Name: EndScrapeEventType,
		Body: EndScrapeEvent{
			ID:        uuid.New().String(),
			Timestamp: end,
			Report:    report,
		},
	})
}
