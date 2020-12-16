package gitlab

type GitLab interface {
	ReadIssues() (IssuesData, error)
}
