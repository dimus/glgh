package glapi

import (
	"context"

	"github.com/dimus/glgh/entity/gitlab"
	"github.com/machinebox/graphql"
)

func (g gitlabAPI) runGraphQL() (gitlab.IssuesData, error) {
	client := graphql.NewClient("https://gitlab.com/api/graphql")
	req := graphqlRequest()
	req.Header.Add("Authorization", "Bearer "+g.cfg.GitLabToken)
	req.Var("repo", g.cfg.GitLabRepoPath())
	ctx := context.Background()
	var resp gitlab.IssuesData
	err := client.Run(ctx, req, &resp)
	if err != nil {
		return gitlab.IssuesData{}, err
	}
	return resp, nil
}
