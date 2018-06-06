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

	Name        string        // name of the job, useful in logs
	MaxFailures int           // max # of consecutive failures, retries endlessly if 0
	Do          func() error  // when started the job calls the do function
	Delay       time.Duration // delays first job start
	Interval    time.Duration // attempts to call the job every interval, if the job takes longer than the interval it will be restarted immediately after finishing
}

func formatDuration(d time.Duration) string {
	return fmt.Sprintf("%dm %ds", int(d.Minutes()), int(d.Seconds())%60)
}

func do(job *Job, output chan *Job) {
	if job.sleep > 0 {

		log.Printf("job '%v' going to sleep for: %s", job.Name, formatDuration(job.sleep))
		time.Sleep(job.sleep)
	}

	job.start = time.Now()
	log.Printf("job '%v' running", job.Name)
	job.lastErr = job.Do()
	job.end = time.Now()
	job.lastDuration = job.end.Sub(job.start)
	log.Printf("job '%v' ended after %s", job.Name, formatDuration(job.lastDuration))

	if job.lastErr != nil {
		job.errors = append(job.errors, job.lastErr)
	} else {
		job.errors = nil
	}

	output <- job
}

// StartJobs starts one more jobs in a fixed interval
// If something is written to quit it waits for all jobs to finish and than returns
func StartJobs(quit chan int, jobs ...*Job) error {
	var jobsRunning uint32

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

		go do(job, output)
		jobsRunning++
		jobsStarted++
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

			if !cancel {
				if job.MaxFailures <= 0 || len(job.errors) < job.MaxFailures {
					job.sleep = job.Interval - job.lastDuration

					go do(job, output)
					jobsRunning++
				} else {
					log.Printf("WARNING: job '%v' failed max. amount of times (%v), stopping...", job.Name, len(job.errors))
				}
			}
		case <-quit:
			log.Printf("user cancelation: waiting for remaining jobs...")
			cancel = true
		}
	}

	return nil
}
