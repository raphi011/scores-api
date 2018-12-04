package job

import (
	"fmt"
	"sync"
	"time"
)

type Execution struct {
	Start        time.Time     `json:"start"`
	End          time.Time     `json:"end"`
	Sleep        time.Duration `json:"sleep"`
	LastDuration time.Duration `json:"lastDuration"`
	Errors       []error       `json:"errors"`
	Runs         uint          `json:"runs"`
	State        State         `json:"state"`

	Job Job `json:"job"`

	quit     chan int
	quitOnce sync.Once
}

func (e *Execution) canStart() bool {
	return e.State == StateStopped || e.State == StateWaiting
}

func (e *Execution) stop() {
	e.quitOnce.Do(func() {
		close(e.quit)
	})
}

func (e *Execution) run() {
	e.State = StateRunning

	e.Start = time.Now()
	err := e.Job.Do()
	e.End = time.Now()

	e.LastDuration = e.End.Sub(e.Start)
	e.Runs++

	if err != nil {
		e.Errors = append(e.Errors, err)

		if e.Job.MaxFailures > 0 && len(e.Errors) >= int(e.Job.MaxFailures) {
			e.State = StateErrored
			return
		}
	} else {
		e.Errors = nil
	}

	if e.Job.MaxRuns > 0 && e.Runs >= e.Job.MaxRuns {
		e.State = StateStopped
		return
	}

	e.Sleep = e.Job.Interval - e.LastDuration
	e.State = StateWaiting
}

func formatDuration(d time.Duration) string {
	return fmt.Sprintf("%dm %ds", int(d.Minutes()), int(d.Seconds())%60)
}

func sleep(sleep time.Duration, wake chan int) {
	time.Sleep(sleep)
	wake <- 1
}
