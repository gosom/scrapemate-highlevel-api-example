package hockey

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Team struct {
	Name         string
	Year         int
	Wins         int
	Losses       int
	OTLosses     int
	WinPct       float64
	GoalsFor     int
	GoalsAgainst int
	GoalDiff     int
}

func (t Team) CsvHeaders() []string {
	return []string{
		"Name",
		"Year",
		"Wins",
		"Losses",
		"OTLosses",
		"WinPct",
		"GoalsFor",
		"GoalsAgainst",
		"GoalDiff",
	}
}

func (t Team) CsvRow() []string {
	return []string{
		t.Name,
		strconv.Itoa(t.Year),
		strconv.Itoa(t.Wins),
		strconv.Itoa(t.Losses),
		strconv.Itoa(t.OTLosses),
		strconv.FormatFloat(t.WinPct, 'f', 2, 64),
		strconv.Itoa(t.GoalsFor),
		strconv.Itoa(t.GoalsAgainst),
		strconv.Itoa(t.GoalDiff),
	}
}

func parseTeams(doc *goquery.Document) ([]Team, error) {
	sel := "table.table tr.team"
	var teams []Team
	doc.Find(sel).Each(func(i int, s *goquery.Selection) {
		teams = append(teams, parseTeam(s))
	})
	return teams, nil
}

func parseTeam(s *goquery.Selection) Team {
	var team Team
	team.Name = cleanText(s.Find("td.name").Text())
	team.Year = parseInt(s.Find("td.year").Text())
	team.Wins = parseInt(s.Find("td.wins").Text())
	team.Losses = parseInt(s.Find("td.losses").Text())
	team.OTLosses = parseInt(s.Find("td.ot-losses").Text())
	team.WinPct = parseFloat(s.Find("td.pct").Text())
	team.GoalsFor = parseInt(s.Find("td.gf").Text())
	team.GoalsAgainst = parseInt(s.Find("td.ga").Text())
	team.GoalDiff = parseInt(s.Find("td.diff").Text())
	return team
}

func parseNextLink(doc *goquery.Document) (string, map[string]string) {
	sel := "ul.pagination>li:last-child>a[aria-label=Next]"
	s := doc.Find(sel).AttrOr("href", "")
	if s == "" {
		return "", nil
	}
	s = "https://www.scrapethissite.com" + s
	parts := strings.Split(s, "?")
	nextLink := parts[0]
	params := make(map[string]string)
	for _, p := range strings.Split(parts[1], "&") {
		kv := strings.Split(p, "=")
		params[kv[0]] = kv[1]
	}
	return nextLink, params
}

func cleanText(s string) string {
	s = strings.TrimFunc(s, func(r rune) bool {
		return r == '\n'
	})
	return strings.TrimSpace(s)
}

func parseInt(s string) int {
	s = cleanText(s)
	if s == "" {
		return 0
	}
	ans, _ := strconv.Atoi(s)
	return ans
}

func parseFloat(s string) float64 {
	s = cleanText(s)
	if s == "" {
		return 0
	}
	ans, _ := strconv.ParseFloat(s, 64)
	return ans
}
