package job

import (
	"errors"
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// Job is the definition of a job which is run the Manager in defined intervals.
type Job struct {
	MaxRuns     uint          `json:"maxRuns"`     // limit the # of runs, no limit if 0
	Name        string        `json:"name"`        // name of the job, useful in logs
	MaxFailures uint          `json:"maxFailures"` // max # of consecutive failures, retries endlessly if 0
	Interval    time.Duration `json:"interval"`    // attempts to call the job every interval, if the job takes longer than the interval it will be restarted immediately after finishing
	Delay       time.Duration `json:"delay"`       // delays first job start

	Do func() error `json:"-"` // when started the job calls the do function
}

// Manager runs jobs in defined intervals
type Manager struct {
	waitGroup sync.WaitGroup
	running   bool

	quit     chan int
	executed chan *Execution

	runningJobs map[string]*Execution
	jobs        []*Job
}

func (s *Manager) executionLoop() {
	run := true

	for run {
		select {
		case exec := <-s.executed:
			if exec.State == StateWaiting {
				s.schedule(exec)
			}
		case <-s.quit:
			s.StopJobs()
			run = false
		}
	}
}

func (s *Manager) schedule(exec *Execution) {
	s.waitGroup.Add(1)
	wake := make(chan int)

	if exec.Sleep > 0 {
		log.Printf("job '%v' going to sleep for: %s", exec.Job.Name, formatDuration(exec.Sleep))
		go sleep(exec.Sleep, wake)

		select {
		case <-wake:
			log.Printf("job '%v' running", exec.Job.Name)
		case <-exec.quit:
			log.Printf("job '%v' canceled", exec.Job.Name)
			exec.State = StateStopped
		}
	}

	if exec.State == StateWaiting {
		exec.run()
	}

	if len(exec.Errors) > 0 {
		err := exec.Errors[len(exec.Errors)-1]
		log.Warnf("job %q failed: %v", exec.Job.Name, err)	
	}

	s.waitGroup.Done()
	s.executed <- exec
}

// Run attempts to run a job referenced by its name (job.Name)
func (s *Manager) Run(jobName string) error {
	exec, ok := s.runningJobs[jobName]

	if !ok {
		return fmt.Errorf("job %q does not exist", jobName)
	}

	return s.RunJob(exec.Job)
}

// Executions returns all running jobs
func (s *Manager) Executions() []Execution {
	executions := []Execution{}

	for _, execution := range s.runningJobs {
		executions = append(executions, *execution)
	}

	return executions
}

// RunJob runs a job
func (s *Manager) RunJob(job Job) error {
	if !s.running {
		return errors.New("Manager must be running before starting a job")
	}

	if job.Do == nil {
		return errors.New("job has no 'Do' function")
	}

	exec, ok := s.runningJobs[job.Name]

	if ok && !exec.canStart() {
		return fmt.Errorf("job %q can't start because it's in state %q",
			job.Name, exec.State)
	}

	var sleep time.Duration

	if !ok && job.Delay > 0 {
		sleep = job.Delay
	}

	exec = &Execution{
		Job:   job,
		Sleep: sleep,
		State: StateWaiting,

		quit: make(chan int),
	}

	s.runningJobs[job.Name] = exec

	go s.schedule(exec)

	return nil
}

// HasJob returns true if a job with the name `jobName` exists.
func (s *Manager) HasJob(jobName string) bool {
	_, ok := s.runningJobs[jobName]

	return ok
}

// StopJob stops a job.
func (s *Manager) StopJob(jobName string) error {
	exec, ok := s.runningJobs[jobName]

	if !ok {
		return fmt.Errorf("job %q does not exist", jobName)
	}

	exec.stop()

	return nil
}

// Start runs the Manager and queues the `jobs`
func (s *Manager) Start(jobs ...Job) (err error) {
	if s.running || len(jobs) == 0 {
		return
	}

	s.runningJobs = make(map[string]*Execution)
	s.quit = make(chan int)
	s.executed = make(chan *Execution)
	s.running = true

	go s.executionLoop()

	for _, job := range jobs {
		err = s.RunJob(job)

		if err != nil {
			s.Stop()
			break
		}
	}

	return err
}

// StopJobs stops all jobs
func (s *Manager) StopJobs() {
	for _, exec := range s.runningJobs {
		exec.stop()
	}
}

// Stop stops the manager and all its jobs
func (s *Manager) Stop() {
	if !s.running {
		return
	}

	close(s.quit)
	s.waitGroup.Wait()
	s.running = false
}
