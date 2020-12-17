package ghapi

import (
	"context"

	"github.com/dimus/glgh/config"
	"github.com/dimus/glgh/entity/github"
	"github.com/machinebox/graphql"
)

const (
	apiURL = "https://api.github.com/graphql"
)

type githubAPI struct {
	cfg config.Config
}

func NewGitHubAPI(cfg config.Config) github.GitHub {
	return githubAPI{cfg: cfg}
}

func (g githubAPI) RepositoryData() (github.Repository, error) {
	client := graphql.NewClient(apiURL)
	req := repoReq()
	req.Header.Add("Authorization", "Bearer "+g.cfg.GitHubToken)
	req.Var("owner", g.cfg.GitHubOwner)
	req.Var("repo", g.cfg.GitHubRepo)
	ctx := context.Background()
	var resp github.Data
	err := client.Run(ctx, req, &resp)
	if err != nil {
		return resp.Repository, err
	}
	return resp.Repository, nil
}

func (g githubAPI) MapUsers(logins []string) (map[string]string, error) {
	res := make(map[string]string)
	client := graphql.NewClient(apiURL)
	req := userIDReq()
	req.Header.Add("Authorization", "Bearer "+g.cfg.GitHubToken)
	for i := range logins {
		req.Var("login", logins[i])
		ctx := context.Background()
		var resp github.UserData
		err := client.Run(ctx, req, &resp)
		if err != nil {
			return res, err
		}
		res[logins[i]] = resp.User.ID
	}
	return res, nil
}
