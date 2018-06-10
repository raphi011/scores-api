package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/raphi011/scores/job"
)

var url = flag.String("url", "http://localhost:8080", "url of scores backend e.g.: http(s)://hostname:port")
var scrapeOnce = flag.Bool("once", false, "run each job only once")

func scrape() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	quit := make(chan int)

	go func() {
		<-sigs
		quit <- 1
	}()

	log.Printf("scraping via url %s", *url)

	maxRuns := uint(0)

	if *scrapeOnce {
		maxRuns = 1
	}

	jobs := []*job.Job{
		&job.Job{
			Name:        "Players",
			Do:          players,
			MaxFailures: 3,
			Interval:    5 * time.Minute,
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
	resp, err := http.Get(*url + "/volleynet/scrape/ladder")

	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("scraping failed with code: %d", resp.StatusCode)
	}

	return nil
}

func tournaments() error {
	resp, err := http.Get(*url + "/volleynet/scrape/tournaments")

	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("scraping failed with code: %d", resp.StatusCode)
	}

	return nil
}
