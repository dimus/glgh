package gitlab

import (
	"strconv"
	"strings"
	"time"

	"github.com/dimus/glgh/entity/github"
)

type IssuesData struct {
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

func (idata IssuesData) ToGithubIssueData() github.IssuesData {
	res := github.IssuesData{
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
		issue := github.Issue{
			Author:    github.Actor{Login: issues[i].Author.Username},
			Number:    id,
			Title:     issues[i].Title,
			Body:      issues[i].Description,
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
