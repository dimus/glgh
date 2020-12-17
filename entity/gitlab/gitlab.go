package gitlab

import (
	"strconv"
	"strings"
	"time"

	"github.com/dimus/glgh/entity/github"
)

type Data struct {
	Project `json:"project"`
}

type Project struct {
	Issues `json:"issues"`
}

type Issues struct {
	Count int     `json:"count"`
	Nodes []Issue `json:"nodes"`
}

type Issue struct {
	IID         string `json:"iid"`
	WebURL      string `json:"webUrl"`
	Author      User   `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Notes       `json:"notes"`
	Labels      `json:"labels"`
	CreatedAt   time.Time  `json:"createdAt"`
	ClosedAt    *time.Time `json:"closedAt,omitempty"`
}

type Notes struct {
	Nodes []Note `json:"nodes"`
}

type Labels struct {
	Nodes []Label `json:"nodes"`
}

type Label struct {
	Title string
}

type User struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type Discussion struct {
	CreatedAt time.Time `json:"createdAt"`
	Notes     `json:"notes"`
}

type Note struct {
	Author    User      `json:"author"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}

func (idata Data) ToGithubIssueData() github.Data {
	res := github.Data{
		Repository: github.Repository{
			Issues: github.Issues{
				TotalCount: idata.Issues.Count,
			},
		},
	}

	issues := idata.Issues.Nodes
	ghIssues := make([]github.Issue, len(issues))

	for i := range issues {
		id, _ := strconv.Atoi(issues[i].IID)
		login := issues[i].Author.Username
		if login == "mjy1" {
			login = "mjy"
		}
		issue := github.Issue{
			Author:    github.Actor{Login: login},
			Number:    id,
			Title:     issues[i].Title,
			Body:      issues[i].Description,
			URL:       issues[i].WebURL,
			CreatedAt: issues[i].CreatedAt,
			ClosedAt:  issues[i].ClosedAt,
			Closed:    issues[i].ClosedAt != nil,
			Comments: github.Comments{
				Nodes: populateComments(issues[i]),
			},
			Labels: github.Labels{
				Nodes: populateLabels(issues[i]),
			},
		}
		ghIssues[i] = issue
	}
	res.Issues.Nodes = ghIssues
	return res
}

func populateComments(issue Issue) []github.Comment {
	notes := issue.Notes.Nodes
	comments := make([]github.Comment, 0)
	for i := range notes {
		login := notes[i].Author.Username
		if login == "mjy1" {
			login = "mjy"
		}
		comment := github.Comment{
			Author:    github.Actor{Login: notes[i].Author.Username},
			Body:      notes[i].Body,
			CreatedAt: notes[i].CreatedAt,
		}
		if !strings.HasPrefix(comment.Body, "closed via commit") {
			comments = append(comments, comment)
		}
	}
	return comments
}

func populateLabels(issue Issue) []github.Label {
	labels := issue.Labels.Nodes
	ghlabels := make([]github.Label, len(labels))
	for i := range labels {
		ghlabels[i] = github.Label{
			Name: labels[i].Title,
		}
	}
	return ghlabels
}
