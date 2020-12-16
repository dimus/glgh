package ghapi

import (
	"github.com/dimus/glgh/config"
	"github.com/dimus/glgh/entity/github"
)

type githubAPI struct {
	cfg config.Config
}

func NewGitHubAPI(cfg config.Config) github.GitHub {
	return githubAPI{cfg: cfg}
}

func (g githubAPI) WriteIssues() error {
	return nil
}
