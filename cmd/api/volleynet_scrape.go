package main

import (
	msync "sync"
	"time"

	"github.com/raphi011/scores/job"
	"github.com/sirupsen/logrus"
)

type Status int

const (
	Stopped  Status = 0
	Running  Status = 1
	Stopping Status = 2
	Errored  Status = 3
)

type Scraper struct {
	mux    msync.Mutex
	active bool
	jobs   []*job.Job
	quit   chan int

	log logrus.FieldLogger
}

func New() *Scraper {
	scraper := &Scraper{
		quit: make(chan int),
		jobs: []*job.Job{
			&job.Job{
				Name:        "Players",
				Do:          players,
				MaxFailures: 3,
				Interval:    1 * time.Hour,
			},
			&job.Job{
				Name:        "Tournaments",
				Do:          tournaments,
				MaxFailures: 3,
				Interval:    5 * time.Minute,
				Delay:       1 * time.Minute,
			},
		},
	}

	return scraper
}

type JobReport struct {
	Status Status `json:"status"`
}

type Report struct {
	Status Status      `json:"status"`
	Jobs   []JobReport `json:"jobs"`
}

func (s *Scraper) Report() Report {
	return Report{}
}

func (s *Scraper) Active() bool {
	return s.active
}

func (s *Scraper) Start() {
	s.active = true

	job.StartJobs(s.quit, s.jobs...)
}

func (s *Scraper) Stop() {
	if !s.active {
		return
	}

	s.quit <- 0
	s.active = false
}

func players() error {
	// for _, gender := range genders {
	// resp, err := http.Get(*urlArg + "/volleynet/scrape/ladder?gender=" + url.QueryEscape(gender))

	// 	if err != nil {
	// 		return err
	// 	} else if resp.StatusCode != http.StatusOK {
	// 		return fmt.Errorf("scraping failed with code: %d", resp.StatusCode)
	// 	}
	// }

	return nil
}

var leagues = []string{"AMATEUR TOUR", "PRO TOUR", "JUNIOR TOUR"}
var genders = []string{"M", "W"}

func tournaments() error {
	var season = time.Now().Year()

	for _, league := range leagues {
		for _, gender := range genders {
			report, err := h.syncService.Tournaments(gender, league, season)
		}
	}

	return nil
}
