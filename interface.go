package glgh

// GlGh interface migrates metadata from GitLab to GitHub.
type GlGh interface {
	// Issues transfers issues from GitLab repo to GitHub repo.
	Issues() error
}
