package todo

import (
	"errors"
	"time"

	"final_project_1/helper"
)

type Service interface {
	CreateTodo(input TodoInput) *helper.Response
	GetTodos() *helper.Response
	GetTodoByid(id int) *helper.Response
	UpdateTodo(id int, input TodoUpdateInput) *helper.Response
	DeleteTodo(id int) *helper.Response
}
type serivce struct {
	repository Repository
}

func NewService(repository Repository) *serivce {
	return &serivce{repository}
}

func (s *serivce) CreateTodo(input TodoInput) *helper.Response {
	var t Todo
	t.Task = input.Task
	t.CreatedAt = time.Now()

	todo, err := s.repository.CreateTodo(t)
	if err != nil {
		return helper.InternalServerError(err)
	}

	todoFormatter := TodoFormatter(todo)
	return helper.SuccessCreateResponse(todoFormatter, "Success create todo task")
}

func (s *serivce) GetTodos() *helper.Response {
	todos, err := s.repository.GetTodos()
	if err != nil {
		return helper.InternalServerError(err)
	}

	todoFormatter := TodoFormatter(todos)
	return helper.SuccessResponse(todoFormatter, "Success get all todo task")
}

func (s *serivce) GetTodoByid(id int) *helper.Response {
	if id == 0 {
		return helper.BadRequestResponse(errors.New("cant find data, try again"))
	}

	todo, err := s.repository.GetTodoByid(id)
	if err != nil {
		return helper.InternalServerError(err)
	}

	todoFormatter := TodoFormatter(todo)
	return helper.SuccessResponse(todoFormatter, "Success get todo tas")
}

func (s *serivce) UpdateTodo(id int, input TodoUpdateInput) *helper.Response {
	if id == 0 {
		return helper.BadRequestResponse(errors.New("cant find data, try again"))
	}

	var t Todo
	t.Id = id
	t.Task = input.Task
	t.Completed = input.Completed
	t.UpdatedAt = time.Now()

	todo, err := s.repository.UpdateTodo(t)
	if err != nil {
		return helper.InternalServerError(err)
	}

	todoFormatter := TodoFormatter(todo)
	return helper.SuccessResponse(todoFormatter, "Success update todo task")
}

func (s *serivce) DeleteTodo(id int) *helper.Response {
	if id == 0 {
		return helper.BadRequestResponse(errors.New("cant find data, try again"))
	}

	err := s.repository.DeleteTodo(id)
	if err != nil {
		return helper.InternalServerError(err)
	}

	return helper.SuccessResponse(nil, "Success delete todo task")
}
