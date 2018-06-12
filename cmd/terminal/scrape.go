package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/raphi011/scores/job"
)

var urlArg = flag.String("url", "http://localhost:8080", "url of scores backend e.g.: http(s)://hostname:port")
var scrapeOnceArg = flag.Bool("once", false, "run each job only once")

func scrape() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	quit := make(chan int)

	go func() {
		<-sigs
		quit <- 1
	}()

	log.Printf("scraping via url %s", *urlArg)

	maxRuns := uint(0)

	if *scrapeOnceArg {
		maxRuns = 1
	}

	jobs := []*job.Job{
		&job.Job{
			Name:        "Players",
			Do:          players,
			MaxFailures: 3,
			Interval:    1 * time.Hour,
			MaxRuns:     maxRuns,
		},
		&job.Job{
			Name:        "Tournaments",
			Do:          tournaments,
			MaxFailures: 3,
			Interval:    5 * time.Minute,
			Delay:       1 * time.Minute,
			MaxRuns:     maxRuns,
		},
	}

	job.StartJobs(quit,
		jobs...,
	)
}

func players() error {
	for _, gender := range genders {
		resp, err := http.Get(*urlArg + "/volleynet/scrape/ladder?gender=" + url.QueryEscape(gender))

		if err != nil {
			return err
		} else if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("scraping failed with code: %d", resp.StatusCode)
		}
	}

	return nil
}

var leagues = []string{"AMATEUR TOUR", "PRO TOUR", "JUNIOR TOUR"}
var genders = []string{"M", "W"}

func tournaments() error {
	for _, league := range leagues {
		for _, gender := range genders {
			resp, err := http.Get(
				*urlArg +
					"/volleynet/scrape/tournaments?league=" +
					url.QueryEscape(league) +
					"&gender=" +
					url.QueryEscape(gender))

			if err != nil {
				return err
			} else if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("scraping failed with code: %d", resp.StatusCode)
			}
		}
	}

	return nil
}
