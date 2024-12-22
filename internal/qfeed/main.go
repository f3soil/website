package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/andybalholm/cascadia"
	"github.com/cenkalti/backoff/v4"
	"github.com/gorilla/feeds"
	"github.com/samber/lo"
	"golang.org/x/net/html"
)

func main() {
	feed, err := GenerateFeed()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: generate feed: %s\n", err.Error())
		os.Exit(1)
	}
	rss, err := feed.ToAtom()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: atom: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(rss)
}

func GenerateFeed() (feeds.Feed, error) {
	now := time.Now()
	chicago, err := time.LoadLocation("America/Chicago")
	if err != nil {
		panic(err)
	}
	firstOfTheYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, chicago)
	startDate := firstOfTheYear
	for {
		if startDate.Weekday() == time.Sunday {
			break
		}
		startDate = startDate.Add(24 * time.Hour)
	}
	publishInterval := 7 * 24 * time.Hour

	feed := feeds.Feed{
		Title:       "F3 QSource",
		Subtitle:    "The F3 Manual of Virtuous Leadership",
		Description: "An annual feed for F3 QSource",
		Created:     startDate,
		Updated:     now.UTC(),
		Author: &feeds.Author{
			Name:  "Rowengartner (F3 SOIL)",
			Email: "nnutter@duck.com",
		},
		Link: &feeds.Link{
			Href: "https://f3soil.com/rss",
		},
		Items: []*feeds.Item{
			{
				Id:      "https://f3nation.com/q/",
				Created: firstOfTheYear,
				Updated: firstOfTheYear,
				Title:   "QSource The F3 Manual of Virtuous Leadership",
				Link: &feeds.Link{
					Href: "https://f3nation.com/q/",
				},
				Content: "Also, download the <a href=\"https://f3nation.com/qsource-best-practices/\">\"Best Practice Manual\"</a>.",
			},
		},
	}

	elapsed := time.Since(startDate)
	qPointsIndex := int(elapsed/publishInterval) + 1
	if qPointsIndex > len(qPoints) {
		return feed, nil
	}
	for i, q := range qPoints[qPointsIndex-1 : qPointsIndex] {
		if err := q.Get(); err != nil {
			return feeds.Feed{}, err
		}
		qPointPubDate := startDate.Add(time.Duration(i) * publishInterval)
		item := feeds.Item{
			Id:      q.Link,
			Created: qPointPubDate,
			Updated: qPointPubDate,
			Title:   q.Title,
			Link: &feeds.Link{
				Href: q.Link,
			},
		}
		if len(q.Socratics) > 0 {
			lines := lo.Map(q.Socratics, func(item string, index int) string {
				return "<li>" + item + "</li>"
			})
			item.Content += "\n<h1>Socratics</h1><ul>\n" + strings.Join(lines, "\n") + "\n</ul>\n"
		}
		if len(q.Spurs) > 0 {
			lines := lo.Map(q.Spurs, func(item string, index int) string {
				return "<li>" + item + "</li>"
			})
			item.Content += "\n<h1>Spurs</h1><ul>\n" + strings.Join(lines, "\n") + "\n</ul>\n"
		}
		if len(q.Body) > 0 {
			item.Content += q.Body
		}
		feed.Items = append(feed.Items, &item)
	}
	return feed, nil
}

func query(n *html.Node, query string) *html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		panic(err)
	}
	return cascadia.Query(n, sel)
}

func queryAll(n *html.Node, query string) []*html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		panic(err)
	}
	return cascadia.QueryAll(n, sel)
}

func hasClass(listNode *html.Node, class string) bool {
	for _, attr := range listNode.Attr {
		if attr.Key != "class" {
			continue
		}
		classes := strings.Fields(attr.Val)
		if lo.Contains(classes, class) {
			return true
		}
	}
	return false
}

type QPoint struct {
	Title string
	Link  string

	Socratics []string
	Spurs     []string
	Body      string
}

func (q *QPoint) Get() error {
	op := func() (*http.Response, error) {
		resp, err := http.Get(q.Link)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("returned %d status code: %s", resp.StatusCode, resp.Request.URL.String())
		}
		return resp, nil
	}
	resp, err := backoff.RetryWithData(op, backoff.NewExponentialBackOff())
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	if err := q.GetBody(doc); err != nil {
		return fmt.Errorf("%w: %s", err, q.Link)
	}
	q.Socratics, err = q.GetPoints(doc, "Socratic")
	if err != nil {
		return fmt.Errorf("%w: %s", err, q.Link)
	}
	q.Spurs, err = q.GetPoints(doc, "Spur")
	if err != nil {
		return fmt.Errorf("%w: %s", err, q.Link)
	}

	return nil
}

func (q *QPoint) GetBody(doc *html.Node) error {
	node := query(doc, "div.c-sub-heading h1")
	if node == nil {
		return fmt.Errorf("could not find body")
	}
	node = node.Parent.Parent.Parent.Parent
	if node == nil {
		return fmt.Errorf("could not find body")
	}

	var b bytes.Buffer
	if err := html.Render(&b, node); err != nil {
		return err
	}

	q.Body = b.String()
	return nil
}

func (q *QPoint) GetPoints(doc *html.Node, name string) ([]string, error) {
	nodes := queryAll(doc, "div.c-sub-heading > div > div > h4 > strong")
	listNode := new(html.Node)
	for _, node := range nodes {
		if node.FirstChild == nil {
			continue
		}
		if !strings.Contains(node.FirstChild.Data, name) {
			// Hack because of broken page for Q4.5 Failure
			if !(name == "Socratic" && strings.Contains(node.FirstChild.Data, "SocBullet listratic")) {
				continue
			}
		}
		listNode = node.Parent.Parent.Parent.Parent.NextSibling
		if listNode == nil {
			continue
		}
		for {
			if listNode == nil {
				return nil, fmt.Errorf("could not find a %s node", name)
			}
			if hasClass(listNode, "c-bullet-list") {
				break
			}
			listNode = listNode.NextSibling
		}
		break
	}
	if listNode == nil {
		return nil, fmt.Errorf("could not find a %s node", name)
	}
	if len(listNode.Data) == 0 {
		return nil, fmt.Errorf("could not find a %s node", name)
	}
	itemNodes := queryAll(listNode, "li > p")
	points := make([]string, 0)
	for _, itemNode := range itemNodes {
		points = append(points, itemNode.FirstChild.Data)
	}
	return points, nil
}

var qPoints = []QPoint{
	{Title: "Disruption (F1)", Link: "https://f3nation.com/disruption"},
	{Title: "Language (F2)", Link: "https://f3nation.com/language"},
	{Title: "Group (F3)", Link: "https://f3nation.com/group"},
	{Title: "LDP (F4)", Link: "https://f3nation.com/ldp"},
	{Title: "Preparedness (F5)", Link: "https://f3nation.com/preparedness"},
	{Title: "Get Right (Q1)", Link: "https://f3nation.com/get-right"},
	{Title: "DRP (Q1.1)", Link: "https://f3nation.com/drp"},
	{Title: "King (Q1.2)", Link: "https://f3nation.com/king"},
	{Title: "Queen (Q1.3)", Link: "https://f3nation.com/queen"},
	{Title: "Jester (Q1.4)", Link: "https://f3nation.com/jester"},
	{Title: "M (Q1.5)", Link: "https://f3nation.com/m"},
	{Title: "Shorties (Q1.6)", Link: "https://f3nation.com/shorties"},
	{Title: "Shield Lock (Q1.7)", Link: "https://f3nation.com/shield-lock"},
	{Title: "Whetstone (Q1.8)", Link: "https://f3nation.com/whetstone"},
	{Title: "Mammon (Q1.9)", Link: "https://f3nation.com/mammon"},
	{Title: "Prayer (Q1.10)", Link: "https://f3nation.com/prayer"},
	{Title: "Study (Q1.11)", Link: "https://f3nation.com/study"},
	{Title: "Meeting (Q1.12)", Link: "https://f3nation.com/meeting"},
	{Title: "Live Right (Q2)", Link: "https://f3nation.com/live-right"},
	{Title: "Impact (Q2.1)", Link: "https://f3nation.com/impact"},
	{Title: "Influence (Q2.2)", Link: "https://f3nation.com/influence"},
	{Title: "Missionality (Q2.3)", Link: "https://f3nation.com/missionality"},
	{Title: "Positive Habit Transfer (Q2.4)", Link: "https://f3nation.com/pht"},
	{Title: "Accountability (Q2.5)", Link: "https://f3nation.com/accountability"},
	{Title: "Correction (Q2.6)", Link: "https://f3nation.com/correction"},
	{Title: "Targeting (Q2.7)", Link: "https://f3nation.com/targeting"},
	{Title: "The Practice Of Virtuous Leadership", Link: "https://f3nation.com/lead-right-q3"},
	{Title: "Shared Leadership (Q3.1)", Link: "https://f3nation.com/shared-leadership"},
	{Title: "Mutual Competence (Q3.2)", Link: "https://f3nation.com/mutual-competence"},
	{Title: "Trust (Q3.3)", Link: "https://f3nation.com/trust"},
	{Title: "Vision (Q3.4)", Link: "https://f3nation.com/vision"},
	{Title: "Articulation (Q3.5)", Link: "https://f3nation.com/articulation"},
	{Title: "Persuasion (Q3.6)", Link: "https://f3nation.com/persuasion"},
	{Title: "Exhortation (Q3.7)", Link: "https://f3nation.com/exhortation"},
	{Title: "Candor (Q3.8)", Link: "https://f3nation.com/candor"},
	{Title: "Commitment (Q3.9)", Link: "https://f3nation.com/commitment"},
	{Title: "Consistency (Q3.10)", Link: "https://f3nation.com/consistency"},
	{Title: "Contentment (Q3.11)", Link: "https://f3nation.com/contentment"},
	{Title: "Courage (Q3.12)", Link: "https://f3nation.com/courage"},
	{Title: "Leave Right (Q4)", Link: "https://f3nation.com/leave-right"},
	{Title: "Sua Sponte Leader (Q4.1)", Link: "https://f3nation.com/sua"},
	{Title: "Schooling (Q4.2)", Link: "https://f3nation.com/schooling"},
	{Title: "Apprenticeship (Q4.3)", Link: "https://f3nation.com/apprenticeship"},
	{Title: "Opportunity (Q4.4)", Link: "https://f3nation.com/opp"},
	{Title: "Failure (Q4.5)", Link: "https://f3nation.com/failure"},
	{Title: "Team Development (Q4.6)", Link: "https://f3nation.com/team-development"},
	{Title: "Equipping (Q4.7)", Link: "https://f3nation.com/equip"},
	{Title: "Accountability.Team (Q4.8)", Link: "https://f3nation.com/accountability-team"},
	{Title: "Missionality.Team (Q4.9)", Link: "https://f3nation.com/missionality-team"},
	{Title: "Lizard Building (Q4.10)", Link: "https://f3nation.com/lizard-building"},
}
