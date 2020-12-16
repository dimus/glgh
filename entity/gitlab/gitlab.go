package gitlab

import "time"

type IssuesData struct {
	Project `json:"project"`
}

type Project struct {
	Issues `json:"issues"`
}

type Issues struct {
	Count int     `json:"count"`
	Nodes []Issue `json:"nodes"`
}

type Issue struct {
	IID         string `json:"id"`
	Author      User   `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Discussions `json:"discussions"`
	Notes       `json:"notes"`
	CreatedAt   time.Time `json:"createdAt"`
	ClosedAt    time.Time `json:"closedAt"`
}

type Discussions struct {
	Nodes []Discussion `json:"nodes"`
}

type Notes struct {
	Nodes []Note `json:"nodes"`
}

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type Discussion struct {
	CreatedAt time.Time `json:"createdAt"`
	Notes     `json:"notes"`
}

type Note struct {
	Author    User      `json:"author"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}
