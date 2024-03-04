package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gorilla/feeds"
	"github.com/samber/lo"
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
	publishInterval := 7 * 24 * time.Hour

	feed := feeds.Feed{
		Id:          "f3soil.com/rss",
		Title:       "F3 QSource",
		Subtitle:    "The F3 Manual of Virtuous Leadership",
		Description: "An annual feed for F3 QSource",
		Created:     startDate,
		Updated:     time.Now().UTC(),
		Author: &feeds.Author{
			Name:  "Rowengartner (F3 SOIL)",
			Email: "nnutter@duck.com",
		},
	}

	elapsed := time.Since(startDate)
	qPointsIndex := int(elapsed/publishInterval) % len(qPoints)
	for i, q := range qPoints[0:qPointsIndex] {
		qPointPubDate := startDate.Add(time.Duration(i) * publishInterval)
		item := feeds.Item{
			Id:      strings.TrimSuffix(strings.TrimPrefix(q.Link, "https://"), "/"),
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
			item.Content += "<ul>\n" + strings.Join(lines, "\n") + "\n</ul>"
		}
		feed.Items = append(feed.Items, &item)
	}
	return feed
}

type QPoint struct {
	Title     string
	Link      string
	Socratics []string
}

var qPoints = []QPoint{
	{
		Title: "QSource The F3 Manual of Virtuous Leadership",
		Link:  "https://f3nation.com/q/",
		Socratics: []string{
			"Also, download the <a href=\"https://f3nation.com/qsource-best-practices/\">\"Best Practice Manual\"</a>.",
		},
	},
	{
		Title: "Disruption (F1)",
		Link:  "https://f3nation.com/q/disruption/",
		Socratics: []string{
			"What do leaders do?",
			"What effect does leadership have on the status quo?",
			"Is there a common characteristic shared by leaders who leave a legacy?",
		},
	},
	{
		Title: "Language (F2)",
		Link:  "https://f3nation.com/q/language/",
		Socratics: []string{
			"How has the theory of leadership changed over time?",
			"Does the culture have any effect upon the language of leadership?",
			"Is there one common language of leadership?",
		},
	},
	{
		Title: "Group (F3)",
		Link:  "https://f3nation.com/q/group/",
		Socratics: []string{
			"Is man’s natural desire solitude or combination?",
			"Are there different types of groups that people form?",
			"Are there any types of groups that are more important than others?",
		},
	},
	{
		Title: "LDP (F4)",
		Link:  "https://f3nation.com/q/ldp/",
		Socratics: []string{
			"Is a Leader born or made?",
			"Are there common elements to leadership development?",
			"Should the leadership development process be controlled from the top of an Organization?",
		},
	},
	{
		Title: "Preparedness (F5)",
		Link:  "https://f3nation.com/q/preparedness/",
		Socratics: []string{
			"Does a man’s need to be prepared change over time?",
			"Is preparedness a mind set or a process?",
			"What drives a man’s desire to be prepared?",
		},
	},
	{
		Title: "Get Right (Q1)",
		Link:  "https://f3nation.com/q/get-right/",
		Socratics: []string{
			"In a man’s pursuit of change in his life, is there anything that must occur first?",
			"Who does the unfit man serve?",
			"Does service to others require a man to abandon service to himself?",
		},
	},
	{
		Title: "DRP (Q1.1)",
		Link:  "https://f3nation.com/q/drp/",
		Socratics: []string{
			"Can a man Get Right by himself?",
			"Do New Year’s Resolutions work?",
			"Are choices good?",
		},
	},
	{
		Title: "King (Q1.2)",
		Link:  "https://f3nation.com/q/king/",
		Socratics: []string{
			"What is the most important component of Fitness?",
			"What is the purpose of exercise?",
			"What is the best way to Accelerate Fitness?",
		},
	},
	{
		Title: "Queen (Q1.3)",
		Link:  "https://f3nation.com/q/queen/",
		Socratics: []string{
			"Can a good exercise routine overcome a bad diet?",
			"Are some foods bad?",
			"What leads a man to descend into gluttony?",
		},
	},
	{
		Title: "Jester (Q1.4)",
		Link:  "https://f3nation.com/q/jester/",
		Socratics: []string{
			"Are some temptations worse than others?",
			"Don’t bad men often do good things, and good men often do bad things?",
			"Can a man ultimately conquer his own demons?",
		},
	},
	{
		Title: "M (Q1.5)",
		Link:  "https://f3nation.com/q/m/",
		Socratics: []string{
			"Is work-life balance a useful concept?",
			"Are some relationships more important than others?",
			"Is meeting your wife halfway good enough?",
		},
	},
	{
		Title: "Shorties (Q1.6)",
		Link:  "https://f3nation.com/q/shorties/",
		Socratics: []string{
			"What can a father hold back from his children?",
			"How important is a father?",
			"Why must a father be strong for his children?",
		},
	},
	{
		Title: "Shield Lock (Q1.7)",
		Link:  "https://f3nation.com/q/shield-lock/",
		Socratics: []string{
			"Are some men born to be lone wolves?",
			"Do Communities benefit when men combine?",
			"Can a man begin the fight against adversity before he encounters it?",
		},
	},
	{
		Title: "Whetstone (Q1.8)",
		Link:  "https://f3nation.com/q/whetstone/",
		Socratics: []string{
			"What is the purpose of mentorship?",
			"Are there any required components to Effective mentorship?",
			"How does a mentor benefit?",
		},
	},
	{
		Title: "Mammon (Q1.9)",
		Link:  "https://f3nation.com/q/mammon/",
		Socratics: []string{
			"Is there anything inherently good about work?",
			"Is work anything more than something we do?",
			"What priority should a man give to his work?",
		},
	},
	{
		Title: "Prayer (Q1.10)",
		Link:  "https://f3nation.com/q/prayer/",
		Socratics: []string{
			"Does Faith play a role in Leadership?",
			"How did we get here?",
			"Are fear and worry bad things?",
		},
	},
	{
		Title: "Study (Q1.11)",
		Link:  "https://f3nation.com/q/study/",
		Socratics: []string{
			"Is there a connection between Leadership and Faith?",
			"How does a man know his belief system is based upon truth?",
			"Is Faith static or dynamic?",
		},
	},
	{
		Title: "Meeting (Q1.12)",
		Link:  "https://f3nation.com/q/meeting/",
		Socratics: []string{
			"Is there something to be gained through corporate worship?",
			"Is there any relationship between Leadership and attendance at church?",
			"Are pastors affected by who is sitting in their pews?",
		},
	},
	{
		Title: "Live Right (Q2)",
		Link:  "https://f3nation.com/q/live-right/",
		Socratics: []string{
			"Are we alive for a reason outside of our own existence?",
			"What happens to a man who lives primarily for himself?",
			"Is there an intrinsic benefit in living for others?",
		},
	},
	{
		Title: "Impact (Q2.1)",
		Link:  "https://f3nation.com/q/impact/",
		Socratics: []string{
			"Can a man self-initiate positive change in his life?",
			"Can a broken man be helped?",
			"Can an unbroken man be convinced to risk abandoning his Status Quo?",
		},
	},
	{
		Title: "Influence (Q2.2)",
		Link:  "https://f3nation.com/q/influence/",
		Socratics: []string{
			"Can you force a man to help himself?",
			"How can one man help another man change his life?",
			"What does Influence look like?",
		},
	},
	{
		Title: "Missionality (Q2.3)",
		Link:  "https://f3nation.com/q/missionality/",
		Socratics: []string{
			"Should a man keep trying new things in order to be Happy?",
			"How does a man know where his IMPACT will be greatest?",
			"What happens if a man does not stay Missional?",
		},
	},
	{
		Title: "Positive Habit Transfer (Q2.4)",
		Link:  "https://f3nation.com/q/positive-habit-transfer/",
		Socratics: []string{
			"Can you learn to do things right from a man who does things wrong?",
			"What is the purpose of building positive Habits?",
			"What is the purpose in passing positive Habits on to others?",
		},
	},
	{
		Title: "Accountability (Q2.5)",
		Link:  "https://f3nation.com/q/accountability/",
		Socratics: []string{
			"Can a man hold himself Accountable?",
			"Are there necessary components to Effective Accountability?",
			"Is there an ideal relationship for mutual Accountability between men?",
		},
	},
	{
		Title: "Correction (Q2.6)",
		Link:  "https://f3nation.com/q/correction/",
		Socratics: []string{
			"Can Virtue be instilled by example alone?",
			"Why would a man decline to tell another man the truth about his shortcomings?",
			"Can a man be a bully even if he is right?",
		},
	},
	{
		Title: "Targeting (Q2.7)",
		Link:  "https://f3nation.com/q/targeting/",
		Socratics: []string{
			"Can a man Live Right without being pointed in any particular direction?",
			"Is there a peace that surpasses all human understanding?",
			"Is there any particular tempo required to Live Right?",
		},
	},
	{
		Title: "The Practice Of Virtuous Leadership",
		Link:  "https://f3nation.com/q/lead-right/",
		Socratics: []string{
			"What enables a man to do what a Leader does?",
			"What makes a man what a Leader is?",
			"What is the difference between a good and bad Leader?",
		},
	},
	{
		Title: "Shared Leadership (Q3.1)",
		Link:  "https://f3nation.com/q/shared-leadership/",
		Socratics: []string{
			"Why do men pool their efforts?",
			"Can anyone Lead together?",
			"Does a Leader gain anything by sharing power?",
		},
	},
	{
		Title: "Mutual Competence (Q3.2)",
		Link:  "https://f3nation.com/q/mutual-competence/",
		Socratics: []string{
			"Is Competence important to Group Success?",
			"Does every Member of a Team have to be equally Competent?",
			"How Competent must the Leader be?",
		},
	},
	{
		Title: "Trust (Q3.3)",
		Link:  "https://f3nation.com/q/trust/",
		Socratics: []string{
			"Is the reward of Shared Leadership worth the risk?",
			"Can a man’s resolve be determined before he is called upon to put it in action?",
			"What is the purpose of a CSAUP?",
		},
	},
	{
		Title: "Vision (Q3.4)",
		Link:  "https://f3nation.com/q/vision/",
		Socratics: []string{
			"How does a Group determine what to do next?",
			"Does an idea have to be big to be good?",
			"Who has Vision?",
		},
	},
	{
		Title: "Articulation (Q3.5)",
		Link:  "https://f3nation.com/q/articulation/",
		Socratics: []string{
			"Is Articulation a Skill that all people possess?",
			"Are great Articulators born or made?",
			"What is the primary goal of Articulation?",
		},
	},
	{
		Title: "Persuasion (Q3.6)",
		Link:  "https://f3nation.com/q/persuasion/",
		Socratics: []string{
			"Can you lie and and bully a man into doing what’s good for him?",
			"What is the biggest Obstacle to Movement?",
			"Must a man agree before he Moves?",
		},
	},
	{
		Title: "Exhortation (Q3.7)",
		Link:  "https://f3nation.com/q/exhortation/",
		Socratics: []string{
			"Is there a difference between Exhortation and encouragement?",
			"What exactly are we afraid of?",
			"Can a man Exhort the breach of something he himself has never experienced?",
		},
	},
	{
		Title: "Candor (Q3.8)",
		Link:  "https://f3nation.com/q/candor/",
		Socratics: []string{
			"Are all Truths self-evident?",
			"Are some things better left unsaid?",
			"What happens if you kill the messenger?",
		},
	},
	{
		Title: "Commitment (Q3.9)",
		Link:  "https://f3nation.com/q/commitment/",
		Socratics: []string{
			"Is there anything a Leader must surrender in order to be Committed?",
			"Is there anything a Leader must do to demonstrate Commitment?",
			"What is more important, loyalty or determination?",
		},
	},
	{
		Title: "Consistency (Q3.10)",
		Link:  "https://f3nation.com/q/consistency/",
		Socratics: []string{
			"Why do complex men often lead simple lives?",
			"How can a man keep from going off the rails?",
			"How can a man learn to be Consistent?",
		},
	},
	{
		Title: "Contentment (Q3.11)",
		Link:  "https://f3nation.com/q/contentment/",
		Socratics: []string{
			"What is the difference between Happiness and Joy?",
			"How does a man learn to govern his emotions?",
			"Does control over one’s environment foster Contentment?",
		},
	},
	{
		Title: "Courage (Q3.12)",
		Link:  "https://f3nation.com/q/courage/",
		Socratics: []string{
			"Does Courage require the absence of Fear?",
			"Can an un-Fit man be brave?",
			"What is Grace?",
		},
	},
	{
		Title: "Leave Right (Q4)",
		Link:  "https://f3nation.com/q/leave-right/",
		Socratics: []string{
			"How does Legacy come about?",
			"Is a man’s Legacy preordained?",
			"What role does Leadership development play in Legacy?",
		},
	},
	{
		Title: "Sua Sponte Leader (Q4.1)",
		Link:  "https://f3nation.com/q/sua-sponte-leader/",
		Socratics: []string{
			"How does a Leader know what to do?",
			"How does a Leader know how to do?",
			"How does a Group ensure that its Leaders know what and how to do?",
		},
	},
	{
		Title: "Schooling (Q4.2)",
		Link:  "https://f3nation.com/q/schooling/",
		Socratics: []string{
			"Should a the theory of Leadership change over time and place?",
			"Can a natural born Leader be Effective without formal training?",
			"When does a Leader know that he has learned enough?",
		},
	},
	{
		Title: "Apprenticeship (Q4.3)",
		Link:  "https://f3nation.com/q/apprenticeship/",
		Socratics: []string{
			"Is there a pattern of learning that best serves the apprentice?",
			"How does a Leader turn head-knowledge into heart-knowledge?",
			"What Incentive does a Leader have to train his subordinates?",
		},
	},
	{
		Title: "Opportunity (Q4.4)",
		Link:  "https://f3nation.com/q/opportunity/",
		Socratics: []string{
			"How do most Groups choose their Leaders?",
			"What criteria should a Group use to evaluate its Leaders?",
			"What is the ideal time to start Leading?",
		},
	},
	{
		Title: "Failure (Q4.5)",
		Link:  "https://f3nation.com/q/failure/",
		Socratics: []string{
			"Is Failure necessary to Leadership development?",
			"Does Failure also require Pain?",
			"What should a Group do with a Leader after he has Failed?",
		},
	},
	{
		Title: "Team Development (Q4.6)",
		Link:  "https://f3nation.com/q/team-development/",
		Socratics: []string{
			"What role does the Team play in Advantage-seeking?",
			"What role does the Q play in Team Development?",
			"How does the Q extend his authority?",
		},
	},
	{
		Title: "Equipping (Q4.7)",
		Link:  "https://f3nation.com/q/equipping/",
		Socratics: []string{
			"Is it human nature to combine together to accomplish things?",
			"How does the Q match Members to Mission within the Team?",
			"Do men know their role by nature?",
		},
	},
	{
		Title: "Accountability.Team (Q4.8)",
		Link:  "https://f3nation.com/q/accountability-team/",
		Socratics: []string{
			"Do Teams need Standards in the way that a man does?",
			"How important is it for the Q to enforce Standards within the Team?",
			"What does the Q do about a man who doesn’t meet the Team’s Standards?",
		},
	},
	{
		Title: "Missionality.Team (Q4.9)",
		Link:  "https://f3nation.com/q/missionality-team/",
		Socratics: []string{
			"Can a Team Prosper without a Mission?",
			"How does the Q keep his Team focused on the Mission?",
			"What danger is there for a Team that loses focus?",
		},
	},
	{
		Title: "Lizard Building (Q4.10)",
		Link:  "https://f3nation.com/q/lizard-building/",
		Socratics: []string{
			"What is the glue that holds an Effective Organization together?",
			"What does the Organizational chart of an Effective Organization look like?",
			"How does a Leader maintain control over an Effective Organization?",
		},
	},
}
