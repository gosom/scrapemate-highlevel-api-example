package main

import (
	"context"
	"encoding/csv"
	"os"

	"github.com/gosom/scrapemate"
	"github.com/gosom/scrapemate-highlevel-api-example/hockey"
	"github.com/gosom/scrapemate/adapters/writers/csvwriter"
	"github.com/gosom/scrapemate/scrapemateapp"
)

func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
		return
	}
	os.Exit(0)
}

func run() error {
	csvWriter := csvwriter.NewCsvWriter(csv.NewWriter(os.Stdout))

	writers := []scrapemate.ResultWriter{
		csvWriter,
	}

	cfg, err := scrapemateapp.NewConfig(writers)
	if err != nil {
		return err
	}
	app, err := scrapemateapp.NewScrapeMateApp(cfg)
	if err != nil {
		return err
	}
	params := map[string]string{
		"page_num": "1",
		"per_page": "100",
	}
	seedJobs := []scrapemate.IJob{
		hockey.NewTeamCollectJob("https://www.scrapethissite.com/pages/forms/", params),
	}
	return app.Start(context.Background(), seedJobs...)
}
