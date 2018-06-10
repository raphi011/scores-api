package job

import (
	"fmt"
	"log"
	"time"
)

type scrapeFunction func() error

type Job struct {
	start        time.Time
	end          time.Time
	sleep        time.Duration
	lastDuration time.Duration
	lastErr      error
	errors       []error
	runs         uint
	canceled     bool

	MaxRuns     uint          // limit the # of runs, no limit if 0
	Name        string        // name of the job, useful in logs
	MaxFailures uint          // max # of consecutive failures, retries endlessly if 0
	Do          func() error  // when started the job calls the do function
	Delay       time.Duration // delays first job start
	Interval    time.Duration // attempts to call the job every interval, if the job takes longer than the interval it will be restarted immediately after finishing
}

func formatDuration(d time.Duration) string {
	return fmt.Sprintf("%dm %ds", int(d.Minutes()), int(d.Seconds())%60)
}

func sleep(sleep time.Duration, wake chan int) {
	time.Sleep(sleep)
	wake <- 1
}

func do(job *Job, output chan *Job, quit chan int) {
	wake := make(chan int)

	if job.sleep > 0 {
		log.Printf("job '%v' going to sleep for: %s", job.Name, formatDuration(job.sleep))
		go sleep(job.sleep, wake)

		select {
		case <-wake:
			log.Printf("job '%v' running", job.Name)
		case <-quit:
			job.canceled = true
			log.Printf("job '%v' canceled", job.Name)
		}
	}

	if !job.canceled {
		job.start = time.Now()
		job.lastErr = job.Do()
		job.end = time.Now()
		job.lastDuration = job.end.Sub(job.start)
		log.Printf("job '%v' ended after %s", job.Name, formatDuration(job.lastDuration))

		job.runs++

		if job.lastErr != nil {
			job.errors = append(job.errors, job.lastErr)
		} else {
			job.errors = nil
		}
	}

	output <- job
}

// StartJobs starts one more jobs in a fixed interval
// If something is written to quit it waits for all jobs to finish and than returns
func StartJobs(quit chan int, jobs ...*Job) error {
	quitJob := make(chan int, len(jobs))
	var jobsRunning int

	output := make(chan *Job)
	jobsStarted := 0
	cancel := false

	for _, job := range jobs {
		if job.Do == nil {
			log.Printf("WARNING: job '%v' has no 'Do' function, skipping...", job)
			continue
		}

		if job.Delay > 0 {
			job.sleep = job.Delay
		}

		jobsRunning++
		jobsStarted++
		go do(job, output, quitJob)
	}

	if jobsStarted == 0 {
		log.Print("WARNING: no jobs started")
	} else {
		log.Printf("started %v job(s)", jobsStarted)
	}

	for !cancel || jobsRunning > 0 {
		select {
		case job := <-output:
			jobsRunning--

			if job.lastErr != nil {
				log.Printf("job '%v' error: %v", job.Name, job.lastErr)
			}

			if job.sleep < 0 {
				log.Printf("WARNING: job '%v' ran longer than the interval duration", job.Name)
			}

			if !cancel && (job.MaxRuns == 0 || job.runs < job.MaxRuns) {
				if job.MaxFailures <= 0 || len(job.errors) < int(job.MaxFailures) {
					job.sleep = job.Interval - job.lastDuration

					jobsRunning++
					go do(job, output, quitJob)
				} else {
					log.Printf("WARNING: job '%v' failed max. amount of times (%v), stopping...", job.Name, len(job.errors))
				}
			}
		case <-quit:
			cancel = true

			log.Printf("user cancelation: waiting for %d remaining jobs...", jobsRunning)

			for i := 0; i < jobsRunning; i++ {
				quitJob <- 1
			}
		}
	}

	return nil
}
