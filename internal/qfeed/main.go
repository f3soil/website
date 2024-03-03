package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gorilla/feeds"
)

func main() {
	feed := GenerateFeed()
	rss, err := feed.ToRss()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s", err.Error())
		os.Exit(1)
	}
	fmt.Println(rss)
}

func GenerateFeed() feeds.Feed {
	startDate := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	feed := feeds.Feed{
		Id:          "f3soil.com/rss",
		Title:       "F3 QSource",
		Subtitle:    "The F3 Manual of Virtuous Leadership",
		Description: "An annual feed for F3 QSource",
		Created:     startDate,
		Updated:     time.Now().UTC(),
		Author: &feeds.Author{
			Name:  "Rowengartner",
			Email: "nnutter@duck.com",
		},
	}

	elapsed := time.Since(startDate)
	weekNumber := int(elapsed/(7*24*time.Hour)) + 1
	qPointsIndex := weekNumber + 1
	if weekNumber < 1 {
		weekNumber = 1
	}
	if weekNumber >= len(qPoints) {
		weekNumber = len(qPoints)
	}
	for i, q := range qPoints[0:qPointsIndex] {
		qPointPubDate := startDate.Add(time.Duration(i) * 7 * 24 * time.Hour)
		feed.Items = append(feed.Items, &feeds.Item{
			Id:      strings.TrimSuffix(strings.TrimPrefix(q.Link, "https://"), "/"),
			Created: qPointPubDate,
			Updated: qPointPubDate,
			Title:   q.Title,
			Link: &feeds.Link{
				Href: q.Link,
			},
		})
	}
	return feed
}

type QPoint struct {
	Title string
	Link  string
}

var qPoints = []QPoint{
	{
		Title: "QSource The F3 Manual of Virtuous Leadership",
		Link:  "https://f3nation.com/q/",
	},
	{
		Title: "Disruption (F1)",
		Link:  "https://f3nation.com/q/disruption/",
	},
	{
		Title: "Language (F2)",
		Link:  "https://f3nation.com/q/language/",
	},
	{
		Title: "Group (F3)",
		Link:  "https://f3nation.com/q/group/",
	},
	{
		Title: "LDP (F4)",
		Link:  "https://f3nation.com/q/ldp/",
	},
	{
		Title: "Preparedness (F5)",
		Link:  "https://f3nation.com/q/preparedness/",
	},
	{
		Title: "Get Right (Q1)",
		Link:  "https://f3nation.com/q/get-right/",
	},
	{
		Title: "DRP (Q1.1)",
		Link:  "https://f3nation.com/q/drp/",
	},
	{
		Title: "King (Q1.2)",
		Link:  "https://f3nation.com/q/king/",
	},
	{
		Title: "Queen (Q1.3)",
		Link:  "https://f3nation.com/q/queen/",
	},
	{
		Title: "Jester (Q1.4)",
		Link:  "https://f3nation.com/q/jester/",
	},
	{
		Title: "M (Q1.5)",
		Link:  "https://f3nation.com/q/m/",
	},
	{
		Title: "SHORTIES (Q1.6)",
		Link:  "https://f3nation.com/q/shorties/",
	},
	{
		Title: "Shield Lock (Q1.7)",
		Link:  "https://f3nation.com/q/shield-lock/",
	},
	{
		Title: "Whetstone (Q1.8)",
		Link:  "https://f3nation.com/q/whetstone/",
	},
	{
		Title: "Mammon (Q1.9)",
		Link:  "https://f3nation.com/q/mammon/",
	},
	{
		Title: "Prayer (Q1.10)",
		Link:  "https://f3nation.com/q/prayer/",
	},
	{
		Title: "Study (Q1.11)",
		Link:  "https://f3nation.com/q/study/",
	},
	{
		Title: "Meeting (Q1.12)",
		Link:  "https://f3nation.com/q/meeting/",
	},
	{
		Title: "Live Right (Q2)",
		Link:  "https://f3nation.com/q/live-right/",
	},
	{
		Title: "Impact (Q2.1)",
		Link:  "https://f3nation.com/q/impact/",
	},
	{
		Title: "Influence (Q2.2)",
		Link:  "https://f3nation.com/q/influence/",
	},
	{
		Title: "Missionality (Q2.3)",
		Link:  "https://f3nation.com/q/missionality/",
	},
	{
		Title: "Positive Habit Transfer (Q2.4)",
		Link:  "https://f3nation.com/q/positive-habit-transfer/",
	},
	{
		Title: "Accountability (Q2.5)",
		Link:  "https://f3nation.com/q/accountability/",
	},
	{
		Title: "Correction (Q2.6)",
		Link:  "https://f3nation.com/q/correction/",
	},
	{
		Title: "Targeting (Q2.7)",
		Link:  "https://f3nation.com/q/targeting/",
	},
	{
		Title: "The Practice Of Virtuous Leadership",
		Link:  "https://f3nation.com/q/lead-right/",
	},
	{
		Title: "Shared Leadership (Q3.1)",
		Link:  "https://f3nation.com/q/shared-leadership/",
	},
	{
		Title: "Mutual Competence (Q3.2)",
		Link:  "https://f3nation.com/q/mutual-competence/",
	},
	{
		Title: "Trust (Q3.3)",
		Link:  "https://f3nation.com/q/trust/",
	},
	{
		Title: "Vision (Q3.4)",
		Link:  "https://f3nation.com/q/vision/",
	},
	{
		Title: "Articulation (Q3.5)",
		Link:  "https://f3nation.com/q/articulation/",
	},
	{
		Title: "Persuasion (Q3.6)",
		Link:  "https://f3nation.com/q/persuasion/",
	},
	{
		Title: "Exhortation (Q3.7)",
		Link:  "https://f3nation.com/q/exhortation/",
	},
	{
		Title: "Candor (Q3.8)",
		Link:  "https://f3nation.com/q/candor/",
	},
	{
		Title: "Commitment (Q3.9)",
		Link:  "https://f3nation.com/q/commitment/",
	},
	{
		Title: "Consistency (Q3.10)",
		Link:  "https://f3nation.com/q/consistency/",
	},
	{
		Title: "Contentment (Q3.11)",
		Link:  "https://f3nation.com/q/contentment/",
	},
	{
		Title: "Courage (Q3.12)",
		Link:  "https://f3nation.com/q/courage/",
	},
	{
		Title: "Leave Right (Q4)",
		Link:  "https://f3nation.com/q/leave-right/",
	},
	{
		Title: "Sua Sponte Leader (Q4.1)",
		Link:  "https://f3nation.com/q/sua-sponte-leader/",
	},
	{
		Title: "Schooling (Q4.2)",
		Link:  "https://f3nation.com/q/schooling/",
	},
	{
		Title: "Apprenticeship (Q4.3)",
		Link:  "https://f3nation.com/q/apprenticeship/",
	},
	{
		Title: "Opportunity (Q4.4)",
		Link:  "https://f3nation.com/q/opportunity/",
	},
	{
		Title: "Failure (Q4.5)",
		Link:  "https://f3nation.com/q/failure/",
	},
	{
		Title: "Team Development (Q4.6)",
		Link:  "https://f3nation.com/q/team-development/",
	},
	{
		Title: "Equipping (Q4.7)",
		Link:  "https://f3nation.com/q/equipping/",
	},
	{
		Title: "Accountability.Team (Q4.8)",
		Link:  "https://f3nation.com/q/accountability-team/",
	},
	{
		Title: "Missionality.Team (Q4.9)",
		Link:  "https://f3nation.com/q/missionality-team/",
	},
	{
		Title: "Lizard Building (Q4.10)",
		Link:  "https://f3nation.com/q/lizard-building/",
	},
}
