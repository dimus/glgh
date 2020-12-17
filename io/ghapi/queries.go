package ghapi

import "github.com/machinebox/graphql"

func userIDReq() *graphql.Request {
	req := graphql.NewRequest(`
  query($login: String!) {
    user(login: $login) {
      id
      login
    }
  }`)
	return req
}

func repoReq() *graphql.Request {
	req := graphql.NewRequest(`
  query($owner: String!, $repo: String!) {
    repository(owner: $owner, name: $repo) {
      id
      labels(first:100) {
        nodes {
          id
          name
        }
      }
    }
  }`)
	return req
}

func createIssueReq() *graphql.Request {
	req := graphql.NewRequest(`
  mutation($repo: ID!, $title: String!, $body: String!, $assignees: [ID!]!,
  $labels: [ID!]){
    createIssue(input:{
      repositoryId: $repo
      title: $title
      body: $body
      assigneeIds: $assignees
      labelIds: $labels
    }) { issue { id } }
  }`)
	return req
}

func closeIssueReq() *graphql.Request {
	req := graphql.NewRequest(`
  mutation($issue: ID!){
    closeIssue(input:{
      issueId: $issue
    }){ issue { id } }
  }`)
	return req
}

func addCommentReq() *graphql.Request {
	req := graphql.NewRequest(`
  mutation($subj: ID!, $body: String!){
    addComment(input:{
      subjectId: $subj
      body: $body
    }) { subject { id } }
  }`)
	return req
}
