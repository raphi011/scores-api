package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/raphi011/scores/job"
)

var url = flag.String("url", "http://localhost:8080", "url of scores backend e.g.: http(s)://hostname:port")

func scrape() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT)

	quit := make(chan int)

	go func() {
		<-sigs
		quit <- 1
	}()

	log.Printf("scraping via url %s", *url)

	playersJob := &job.Job{
		Name:        "Players",
		Do:          players,
		MaxFailures: 3,
		Interval:    5 * time.Minute,
	}
	tournamentsJob := &job.Job{
		Name:        "Tournaments",
		Do:          tournaments,
		MaxFailures: 3,
		Interval:    5 * time.Minute,
		Delay:       1 * time.Minute,
	}

	job.StartJobs(quit,
		playersJob,
		tournamentsJob,
	)
}

func players() error {
	resp, err := http.Get(*url + "/volleynet/scrape/ladder")

	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return errors.New("scraping failed")
	}

	return nil
}

func tournaments() error {
	resp, err := http.Get(*url + "/volleynet/scrape/tournaments")

	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return errors.New("scraping failed")
	}

	return nil
}
