package main

import (
	"fmt"
	"os"
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
		feed.Items = append(feed.Items, &feeds.Item{
			Created: startDate.Add(time.Duration(i+1) * 7 * 24 * time.Hour),
			Title:   q.Title,
			Link: &feeds.Link{
				Href:   q.Link,
				Rel:    "",
				Type:   "",
				Length: "",
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
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
	{
		Title: "",
		Link:  "",
	},
}
