package todo

import "time"

type Todo struct {
	Id        int
	Task      string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
