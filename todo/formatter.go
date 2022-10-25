package todo

import (
	"time"
)

type Formatter struct {
	Id        int       `json:"id" example:"1"`
	Task      string    `json:"task" example:"push up 1 sets"`
	Completed bool      `json:"completed" example:"false"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2022-10-25T17:50:24.701221716+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2022-10-25T17:50:24.701221716+07:00"`
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
