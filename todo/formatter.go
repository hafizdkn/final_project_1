package todo

import (
	"time"
)

type Formatter struct {
	Id        int       `json:"id"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func TodoFormatter(todo interface{}) interface{} {
	var sliceOfFormatter []Formatter
	var todoFormatter interface{}

	if t, ok := todo.([]Todo); ok {
		for _, todo := range t {
			t := parseTodoToFormatter(todo)
			sliceOfFormatter = append(sliceOfFormatter, t)
		}
		todoFormatter = sliceOfFormatter
	}

	if t, ok := todo.(Todo); ok {
		t := parseTodoToFormatter(t)
		todoFormatter = t
	}
	return todoFormatter
}

func parseTodoToFormatter(todo Todo) Formatter {
	return Formatter{
		Id:        todo.Id,
		Task:      todo.Task,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}
