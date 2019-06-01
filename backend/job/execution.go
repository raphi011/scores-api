package job

import (
	"time"
)

const (
	// SignalStop signals a job to stop
	SignalStop = 0
	// SignalStart signals a job to wake up and run
	SignalStart = 1
)

// Execution represents a running job
type Execution struct {
	LastRun      time.Time     `json:"lastRun"`
	LastDuration time.Duration `json:"lastDuration"`

	Errors []error `json:"errors"`
	Runs   uint    `json:"runs"`
	State  State   `json:"state"`

	start time.Time
	end   time.Time
	sleep time.Duration

	signal chan int
}
