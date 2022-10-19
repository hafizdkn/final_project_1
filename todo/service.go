package todo

import (
	"time"

	"final_project_1/helper"
)

type Service interface {
	CreateTodo(input TodoInput) *helper.Response
	GetTodos() *helper.Response
	GetTodoByid(id int) *helper.Response
	UpdateTodo(id int, input TodoUpdateInput) *helper.Response
}
type serivce struct {
	repository Repository
}

func NewService(repository Repository) *serivce {
	return &serivce{repository}
}

func (s *serivce) CreateTodo(input TodoInput) *helper.Response {
	var todo Todo
	todo.Task = input.Task
	todo.CreatedAt = time.Now()
	resp := s.repository.CreateTodo(todo)
	return resp
}

func (s *serivce) GetTodos() *helper.Response {
	resp := s.repository.GetTodos()
	return resp
}

func (s *serivce) GetTodoByid(id int) *helper.Response {
	resp := s.repository.GetTodoByid(id)
	return resp
}

func (s *serivce) UpdateTodo(id int, input TodoUpdateInput) *helper.Response {
	var t Todo

	t.Id = id
	t.Task = input.Task
	t.Completed = input.Completed
	t.UpdatedAt = time.Now()

	resp := s.repository.UpdateTodo(t)
	return resp
}
