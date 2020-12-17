package gitlab

type GitLab interface {
	ReadIssues() (Data, error)
}
