package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/raphi011/scores/job"
)

func scrape() {
	scrapeJob := &job.Job{Name: "Tournaments", Do: scrapeTournament}

	channel := make(chan int)

	job.StartJob(30*time.Second, channel, scrapeJob)
}

func scrapeTournament() error {
	resp, err := http.Get("http://localhost:8080/volleynet/scrape/tournaments")

	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return errors.New("scraping failed")
	}

	return nil
}
