package glgh

import (
	"fmt"
	"log"

	"github.com/dimus/glgh/config"
	"github.com/dimus/glgh/entity/github"
	"github.com/dimus/glgh/entity/gitlab"
	"github.com/dimus/glgh/io/ghapi"
	"github.com/dimus/glgh/io/glapi"
	"github.com/gnames/gnlib/encode"
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
	json, _ := encode.GNjson{}.Encode(ghissues)
	fmt.Printf("%+v\n", string(json))
	return nil
}
