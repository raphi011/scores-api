package cron

import (
	"time"

	"github.com/raphi011/scores/job"
	"github.com/raphi011/scores/volleynet/sync"
)

// JobReport contains information about a running job
type JobReport struct {
	State        job.State     `json:"state"`
	RunningSince time.Duration `json:"runningSince"`
}

type LadderJob struct {
	SyncService *sync.Service
	Genders     []string
}

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

type TournamentsJob struct {
	SyncService *sync.Service
	Leagues     []string
	Genders     []string
	Season      int
}

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
