package glgh

import (
	"log"

	"github.com/dimus/glgh/config"
	"github.com/dimus/glgh/entity/github"
	"github.com/dimus/glgh/entity/gitlab"
	"github.com/dimus/glgh/io/ghapi"
	"github.com/dimus/glgh/io/glapi"
)

type glgh struct {
	cfg config.Config
	gl  gitlab.GitLab
	gh  github.GitHub
}

func NewGlGh(cfg config.Config) GlGh {
	return glgh{
		cfg: cfg,
		gl:  glapi.NewGitLabAPI(cfg),
		gh:  ghapi.NewGitHubAPI(cfg),
	}
}

func (g glgh) Issues() error {
	issues, err := g.gl.ReadIssues()
	if err != nil {
		log.Fatal(err)
	}
	ghissues := issues.ToGithubIssueData()
	logins := ghissues.UniqueLogins()
	users, err := g.gh.MapUsers(logins)
	if err != nil {
		log.Printf("Error %T: %s", err, err)
	}
	repo, err := g.gh.RepositoryData()
	if err != nil {
		log.Printf("Error %T: %s", err, err)
	}
	err = g.gh.WriteIssues(ghissues.Issues.Nodes, repo, users)
	if err != nil {
		log.Printf("Error %T: %s", err, err)
	}
	return nil
}
