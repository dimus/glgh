package glapi

import "github.com/machinebox/graphql"

func graphqlRequest() *graphql.Request {
	req := graphql.NewRequest(`
query($repo: ID!) {
  project(fullPath: $repo) {
    issues {
      count
      nodes {
        iid
        author {
          username
          name
        }
        title
        description
        notes {
          nodes {
            author {
              username
              name
            }
            body
            createdAt
          }
        }
        createdAt
        closedAt
        labels {
          nodes {
            title
          }
        }
      }
    }
  }
}`)
	return req
}
