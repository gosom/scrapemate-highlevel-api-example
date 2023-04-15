package hockey

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
	"github.com/gosom/scrapemate"
)

type TeamCollectJob struct {
	scrapemate.Job
}

func NewTeamCollectJob(u string, params map[string]string) *TeamCollectJob {
	return &TeamCollectJob{
		Job: scrapemate.Job{
			// just give it a random id
			ID:        uuid.New().String(),
			Method:    http.MethodGet,
			URL:       u,
			UrlParams: params,
			Headers: map[string]string{
				"User-Agent": scrapemate.DefaultUserAgent,
			},
			Timeout:    10 * time.Second,
			MaxRetries: 3,
		},
	}
}

func (o *TeamCollectJob) Process(ctx context.Context, resp scrapemate.Response) (any, []scrapemate.IJob, error) {
	doc, ok := resp.Document.(*goquery.Document)
	if !ok {
		return nil, nil, fmt.Errorf("invalid document type %T expected *goquery.Document", resp.Document)
	}
	teams, err := parseTeams(doc)
	if err != nil {
		return nil, nil, err
	}

	var nextJobs []scrapemate.IJob

	nextLink, params := parseNextLink(doc)
	if nextLink != "" {
		nextJobs = append(nextJobs, NewTeamCollectJob(nextLink, params))
	}

	return teams, nextJobs, nil
}
