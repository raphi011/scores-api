package job

import (
	"errors"
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type Job struct {
	MaxRuns     uint          `json:"maxRuns"`     // limit the # of runs, no limit if 0
	Name        string        `json:"name"`        // name of the job, useful in logs
	MaxFailures uint          `json:"maxFailures"` // max # of consecutive failures, retries endlessly if 0
	Interval    time.Duration `json:"interval"`    // attempts to call the job every interval, if the job takes longer than the interval it will be restarted immediately after finishing
	Delay       time.Duration `json:"delay"`       // delays first job start

	Do func() error `json:"-"` // when started the job calls the do function
}

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

	s.waitGroup.Done()
	s.executed <- exec
}

func (s *Manager) Run(jobName string) error {
	exec, ok := s.runningJobs[jobName]

	if !ok {
		return fmt.Errorf("job %q does not exist", jobName)
	}

	return s.RunJob(exec.Job)
}

func (s *Manager) Executions() []Execution {
	executions := []Execution{}

	for _, execution := range s.runningJobs {
		executions = append(executions, *execution)
	}

	return executions
}

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

func (s *Manager) HasJob(jobName string) bool {
	_, ok := s.runningJobs[jobName]

	return ok
}

func (s *Manager) StopJob(jobName string) error {
	exec, ok := s.runningJobs[jobName]

	if !ok {
		return fmt.Errorf("job %q does not exist", jobName)
	}

	exec.stop()

	return nil
}

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

func (s *Manager) StopJobs() {
	for _, exec := range s.runningJobs {
		exec.stop()
	}
}

func (s *Manager) Stop() {
	if !s.running {
		return
	}

	close(s.quit)
	s.waitGroup.Wait()
	s.running = false
}
