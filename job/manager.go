package job

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Manager runs jobs in defined intervals
type Manager struct {
	waitGroup sync.WaitGroup

	jobs map[string]*Job

	log *zap.SugaredLogger
}

// NewManager constructs a new Manager.
func NewManager() *Manager {
	log, _ := zap.NewProduction()

	return &Manager{
		log: log.Sugar(),
	}
}

// run starts an execution and sets the appropriate state.
func (s *Manager) run(job *Job) {
	job.Execution.lock.Lock()
	job.Execution.State = StateRunning
	job.Execution.lock.Unlock()
	s.log.Debugf("job %q running", job.Name)

	start := time.Now()
	err := job.Do()
	end := time.Now()

	job.Execution.lock.Lock()
	defer job.Execution.lock.Unlock()

	job.Execution.LastDuration = end.Sub(start)
	job.Execution.LastRun = time.Now()
	job.Execution.Runs++

	if err != nil {
		s.log.Warnf("job %q failed: %v", job.Name, err)
		job.Execution.Errors = append(job.Execution.Errors, err)
	} else {
		s.log.Debugf("job %q finished", job.Name)
		job.Execution.Errors = nil
	}

	if job.hasFailed() {
		s.log.Warnf("job %q failed %d times, stopping", job.Name, job.MaxFailures)
		job.Execution.State = StateErrored
	} else if job.shouldStop() {
		s.log.Infof("job %q ran %d times, stopping", job.Name, job.Execution.Runs)
		job.Execution.State = StateStopped
	} else {
		job.Execution.State = StateWaiting
	}
}

// schedule adds an execution to the run queue during the next interval
// if the execution's state allows it (has not errored, isn't stopped).
// Must be run in a go routine.
func (s *Manager) schedule(job *Job) {
	s.waitGroup.Add(1)

	for {
		sleep := time.Duration(0)

		if job.Execution.Runs == 0 {
			sleep = job.Delay
		} else {
			sleep = job.Interval
		}

		if sleep > 0 {
			s.log.Debugf("job %q going to sleep for: %s", job.Name, formatDuration(sleep))

			sleepTimer := time.AfterFunc(sleep, func() {
				job.Execution.signal <- SignalStart
			})

			if signal := <-job.Execution.signal; signal == SignalStop {
				sleepTimer.Stop()

				job.Execution.lock.Lock()
				job.Execution.State = StateStopped
				job.Execution.lock.Unlock()
				s.log.Debugf("job %q stopped", job.Name)
				break
			}

			s.log.Debugf("job %q woken up", job.Name)
		}

		s.run(job)

		if job.Execution.State != StateWaiting {
			break
		}
	}

	s.waitGroup.Done()
}

func formatDuration(d time.Duration) string {
	return fmt.Sprintf("%dm %ds", int(d.Minutes()), int(d.Seconds())%60)
}

// Jobs returns all running jobs.
func (s *Manager) Jobs() []Job {
	jobs := []Job{}

	for _, job := range s.jobs {
		jobs = append(jobs, *job)
	}

	return jobs
}

// HasJob returns true if a job with the name `jobName` exists.
func (s *Manager) HasJob(jobName string) bool {
	_, ok := s.jobs[jobName]

	return ok
}

// Job retrieves a job.
func (s *Manager) Job(jobName string) (Job, bool) {
	j, ok := s.jobs[jobName]

	if ok {
		j.Execution.lock.Lock()
		job := *j
		j.Execution.lock.Unlock()

		return job, true
	}

	return Job{}, false
}

// Start runs the Manager and queues the `jobs`
func (s *Manager) Start(jobs ...Job) error {
	if len(jobs) == 0 {
		return nil
	}

	s.jobs = make(map[string]*Job)

	for _, job := range jobs {
		if job.Do == nil {
			return errors.New("job has no 'Do' function")
		}
	}

	for i := range jobs {
		job := &jobs[i]

		job.Execution.signal = make(chan int)
		s.jobs[job.Name] = job

		go s.schedule(job)
	}

	return nil
}
