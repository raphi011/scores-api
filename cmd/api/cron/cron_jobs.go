package cron

import (
	"time"

	"github.com/raphi011/scores-backend/job"
	"github.com/raphi011/scores-backend/volleynet/sync"
)

// JobReport contains information about a running job
type JobReport struct {
	State        job.State     `json:"state"`
	RunningSince time.Duration `json:"runningSince"`
}

// LadderJob is a job that scrapes the volleynet ladder.
type LadderJob struct {
	SyncService *sync.Service
	Genders     []string
}

// Do runs the scrape job.
func (j *LadderJob) Do() error {
	for _, gender := range j.Genders {
		_, err := j.SyncService.Ladder(gender)

		if err != nil {
			return err
		}
	}

	return nil
}

var leagues = []string{"amateur-tour", "pro-tour", "junior-tour"}
var genders = []string{"M", "W"}

// TournamentsJob is a job that scrapes tournaments with the given filters.
type TournamentsJob struct {
	SyncService *sync.Service
	Leagues     []string
	Genders     []string
	Season      int
}

// Do runs the scrape job.
func (j *TournamentsJob) Do() error {
	for _, league := range j.Leagues {
		for _, gender := range j.Genders {
			err := j.SyncService.Tournaments(gender, league, j.Season)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
