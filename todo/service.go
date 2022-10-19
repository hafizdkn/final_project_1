package todo

import (
	"time"

	"final_project_1/helper"
)

type Service interface {
	CreateTodo(input TodoInput) *helper.Response
	GetTodos() *helper.Response
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
