package main

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

type ladderJob struct {
	syncService *sync.Service
	genders     []string
}

func (j *ladderJob) do() error {
	for _, gender := range j.genders {
		_, err := j.syncService.Ladder(gender)

		if err != nil {
			return err
		}
	}

	return nil
}

var leagues = []string{"amateur-tour", "pro-tour", "junior-tour"}
var genders = []string{"M", "W"}

type tournamentsJob struct {
	syncService *sync.Service
	leagues     []string
	genders     []string
	season      int
}

func (j *tournamentsJob) do() error {
	for _, league := range j.leagues {
		for _, gender := range j.genders {
			err := j.syncService.Tournaments(gender, league, j.season)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
