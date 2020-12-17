package github

import "time"

type Data struct {
	Repository `json:"repository"`
}

type UserData struct {
	User `json:"user"`
}

type Repository struct {
	ID     string `json:"id,omitempty"`
	Labels `json:"labels"`
	Issues `json:"issues,omitempty"`
}

type Issues struct {
	TotalCount int     `json:"totalCount"`
	Nodes      []Issue `json:"nodes"`
}

type Issue struct {
	Author    Actor      `json:"author"`
	ID        string     `json:"id,omitempty"`
	Number    int        `json:"number"`
	Title     string     `json:"title"`
	Body      string     `json:"body,omitempty"`
	URL       string     `json:"url"`
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
	ID    string `json:"id,omitempty"`
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
	Nodes []Label `json:"nodes"`
}

type Label struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (idata Data) UniqueLogins() []string {
	users := make(map[string]struct{})
	issues := idata.Issues.Nodes
	for i := range issues {
		users[issues[i].Author.Login] = struct{}{}
		notes := issues[i].Comments.Nodes
		for ii := range notes {
			users[notes[ii].Author.Login] = struct{}{}
		}
	}

	res := make([]string, len(users))
	var count int
	for k := range users {
		res[count] = k
		count++
	}
	return res
}
