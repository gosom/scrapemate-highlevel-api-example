package hockey

import (
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/require"
)

func Test_parseTeams(t *testing.T) {
	fd, err := os.Open("../testdata/teams.html")
	require.NoError(t, err)
	defer fd.Close()
	doc, err := goquery.NewDocumentFromReader(fd)
	require.NoError(t, err)

	teams, err := parseTeams(doc)
	require.NoError(t, err)
	require.Equal(t, 100, len(teams))

	team := teams[0]
	require.Equal(t, "Boston Bruins", team.Name)
	require.Equal(t, 1990, team.Year)
	require.Equal(t, 44, team.Wins)
	require.Equal(t, 24, team.Losses)
	require.Equal(t, 0, team.OTLosses)
	require.Equal(t, 0.55, team.WinPct)
	require.Equal(t, 299, team.GoalsFor)
	require.Equal(t, 264, team.GoalsAgainst)
	require.Equal(t, 35, team.GoalDiff)
}

func Test_parseNextLink(t *testing.T) {
	fd, err := os.Open("../testdata/teams.html")
	require.NoError(t, err)
	defer fd.Close()
	doc, err := goquery.NewDocumentFromReader(fd)
	require.NoError(t, err)

	nextLink, params := parseNextLink(doc)
	require.Equal(t, "https://www.scrapethissite.com/pages/forms/", nextLink)
	require.Equal(t, "2", params["page_num"])
	require.Equal(t, "100", params["per_page"])
}
