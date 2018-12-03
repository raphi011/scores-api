package main

import (
	"time"

	"github.com/raphi011/scores/job"
	"github.com/raphi011/scores/volleynet/sync"
)

type JobReport struct {
	State        job.State     `json:"state"`
	RunningSince time.Duration `json:"runningSince"`
}

type ladderJob struct {
	syncService *sync.SyncService
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

var leagues = []string{"AMATEUR TOUR", "PRO TOUR", "JUNIOR TOUR"}
var genders = []string{"M", "W"}

type tournamentsJob struct {
	syncService *sync.SyncService
	leagues     []string
	genders     []string
}

func (j *tournamentsJob) do() error {
	season := time.Now().Year()

	for _, league := range j.leagues {
		for _, gender := range j.genders {
			_, err := j.syncService.Tournaments(gender, league, season)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
