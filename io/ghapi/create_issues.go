package ghapi

import (
	"context"
	"fmt"
	"time"

	"github.com/dimus/glgh/entity/github"
	"github.com/machinebox/graphql"
)

const (
	defaultDeveloper = "dimus"
	pause            = 3000
)

var client = graphql.NewClient(apiURL)

func (g githubAPI) WriteIssues(issues []github.Issue,
	repo github.Repository, logins map[string]string) error {
	for i := range issues {
		time.Sleep(pause * time.Millisecond)
		id, err := g.createIssue(issues[i], repo, logins)
		if err != nil {
			return err
		}
		fmt.Printf("Issue %d: %s", issues[i].Number, id)

		comments := issues[i].Comments.Nodes
		for i := range comments {
			time.Sleep(pause * time.Millisecond)
			err = g.createComment(id, comments[i], issues[i].URL)
			if err != nil {
				return err
			}
		}
		if issues[i].Closed {
			time.Sleep(pause * time.Millisecond)
			err = g.closeIssue(id)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (g githubAPI) createIssue(issue github.Issue, repo github.Repository,
	logins map[string]string) (string, error) {
	req := createIssueReq()
	req.Header.Add("Authorization", "Bearer "+g.cfg.GitHubToken)
	req.Var("repo", repo.ID)
	req.Var("title", issue.Title)
	req.Var("body", body(issue))
	req.Var("assignees", []string{logins[defaultDeveloper]})
	req.Var("labels", labels(issue.Labels.Nodes, repo.Labels.Nodes))
	ctx := context.Background()
	var resp struct {
		CreateIssue struct {
			Issue struct {
				ID string `json:"id"`
			} `json:"issue"`
		} `json:"createIssue"`
	}
	err := client.Run(ctx, req, &resp)
	fmt.Printf("issue: %+v", resp)
	if err != nil {
		return "", err
	}
	return resp.CreateIssue.Issue.ID, nil
}

func (g githubAPI) closeIssue(id string) error {
	req := closeIssueReq()
	req.Header.Add("Authorization", "Bearer "+g.cfg.GitHubToken)
	req.Var("issue", id)
	fmt.Printf("reqClose: %+v", req)
	ctx := context.Background()
	resp := struct{}{}
	return client.Run(ctx, req, &resp)
}

func (g githubAPI) createComment(id string,
	comment github.Comment, url string) error {
	req := addCommentReq()
	req.Header.Add("Authorization", "Bearer "+g.cfg.GitHubToken)
	req.Var("subj", id)
	req.Var("body", commentBody(comment, url))
	ctx := context.Background()
	var resp struct{}
	err := client.Run(ctx, req, &resp)
	if err != nil {
		fmt.Println("comment trouble")
		return err
	}
	return nil
}

func body(issue github.Issue) string {
	return fmt.Sprintf("created by @%s at %s\n\n%s", issue.Author.Login,
		issue.URL, issue.Body)
}

func commentBody(comment github.Comment, url string) string {
	return fmt.Sprintf("created by @%s at %s\n\n%s", comment.Author.Login, url,
		comment.Body)
}

func labels(labels []github.Label, lookup []github.Label) []string {
	// not efficient, but good enough
	var res []string
	for i := range labels {
		for ii := range lookup {
			if lookup[ii].Name == labels[i].Name {
				res = append(res, lookup[ii].ID)
			}
		}
	}
	return res
}
