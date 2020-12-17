package github

// GitHub interface contains methods that are needed for creating issues for
// a repository.
type GitHub interface {
	// RepositoryID finds a respoitory metadata based on owner and name of the
	// repository.
	RepositoryData() (Repository, error)
	// MapUsers takes all login strings used in issues and searches GitHub for
	// IDs for the logins. It returns a map of logins used as keys, and IDs used
	// as values. If ID is not found, the value is empty string.
	MapUsers([]string) (map[string]string, error)
	// Creates issues for a repository. It assumes that the repository has no
	// issues.
	WriteIssues([]Issue, Repository, map[string]string) error
}
