package job

import "time"

// Job is the definition of a job which is run the Manager in defined intervals.
type Job struct {
	MaxRuns     uint          `json:"maxRuns"`     // limit the # of runs, no limit if 0
	Name        string        `json:"name"`        // name of the job, useful in logs
	MaxFailures uint          `json:"maxFailures"` // max # of consecutive failures, retries endlessly if 0
	Interval    time.Duration `json:"interval"`    // attempts to call the job every interval, if the job takes longer than the interval it will be restarted immediately after finishing
	Delay       time.Duration `json:"delay"`       // delays first job start

	Execution Execution `json:"execution"`

	Do func() error `json:"-"` // when started the job calls the do function
}

func (j *Job) hasFailed() bool {
	errors := len(j.Execution.Errors)
	maxFailures := int(j.MaxFailures)

	return errors > 0 && maxFailures > 0 && errors >= maxFailures
}

func (j *Job) canStart() bool {
	state := j.Execution.State

	return state == StateStopped || state == StateWaiting || state == StateErrored
}

func (j *Job) start() {
	j.Execution.signal <- SignalStart
}

func (j *Job) canRun() bool {
	return j.Execution.State == StateWaiting
}

func (j *Job) isActive() bool {
	return j.Execution.State == StateWaiting || j.Execution.State == StateRunning

}

func (j *Job) shouldStop() bool {
	return j.MaxRuns > 0 && j.Execution.Runs >= j.MaxRuns
}

func (j *Job) stop() {

}
