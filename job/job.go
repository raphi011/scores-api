package job

import (
	"log"
	"time"
)

type scrapeFunction func() error

type Job struct {
	start   time.Time
	end     time.Time
	sleep   time.Duration
	lastErr error
	errors  []error

	Name        string
	MaxFailures int
	Do          func() error
}

func do(job *Job, output chan *Job) {
	if job.sleep > 0 {
		log.Printf("job '%v' going to sleep for: %v", job.Name, job.sleep)
		time.Sleep(job.sleep)
	}

	job.start = time.Now()
	log.Printf("job '%v' running at %v", job.Name, job.start)
	job.lastErr = job.Do()
	job.end = time.Now()
	log.Printf("job '%v' ended at %v", job.Name, job.end)

	job.errors = append(job.errors, job.lastErr)

	output <- job
}

func StartJob(interval time.Duration, quit chan int, jobs ...*Job) error {
	output := make(chan *Job)
	jobsStarted := 0

	for _, job := range jobs {
		if job.Do == nil {
			log.Printf("WARNING: job '%v' has no 'Do' function, skipping...", job)
			continue
		}

		go do(job, output)
		jobsStarted++
	}

	if jobsStarted == 0 {
		log.Print("WARNING: no jobs started")
	} else {
		log.Printf("started %v job(s)", jobsStarted)
	}

	for {
		select {
		case currentJob := <-output:
			if currentJob.lastErr != nil {
				log.Print(currentJob.lastErr)
			}

			if currentJob.sleep < 0 {
				log.Printf("WARNING: job '%v' ran longer than the interval duration", currentJob.Name)
			}

			if currentJob.MaxFailures <= 0 || len(currentJob.errors) < currentJob.MaxFailures {
				duration := currentJob.end.Sub(currentJob.start)
				currentJob.sleep = interval - duration

				go do(currentJob, output)
			} else {
				log.Printf("WARNING: job '%v' failed max. amount of times (%v), stopping...", currentJob.Name, len(currentJob.errors))
			}

		case <-quit:
			return nil
		}
	}
}
