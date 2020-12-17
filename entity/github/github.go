package github

import "time"

type IssuesData struct {
	Repository `json:"repository"`
}

type Repository struct {
	Issues `json:"issues"`
}

type Issues struct {
	TotalCount int     `json:"count"`
	Nodes      []Issue `json:"nodes"`
}

type Issue struct {
	Author    Actor      `json:"author"`
	ID        string     `json:"id,omitempty"`
	Number    int        `json:"number"`
	Title     string     `json:"title"`
	Body      string     `json:"body,omitempty"`
	Closed    bool       `json:"closed"`
	ClosedAt  *time.Time `json:"closedAt"`
	CreatedAt time.Time  `json:"createdAt"`
	Assignees `json:"assignees,omitempty"`
	Comments  `json:"comments,omitempty"`
	Labels    `json:"labels,omitempty"`
}

type Assignees struct {
	Nodes []User `json:"nodes"`
}

type Actor struct {
	Login string `json:"login"`
}

type User struct {
	Login string `json:"login"`
}

type Comments struct {
	Nodes []Comment `json:"comment"`
}

type Comment struct {
	Author    Actor     `json:"author"`
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}

type Labels struct {
	Nodes []Label `json:"label"`
}

type Label struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
