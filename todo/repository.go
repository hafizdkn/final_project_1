package todo

import (
	"gorm.io/gorm"

	"final_project_1/helper"
)

type Repository interface {
	CreateTodo(todo Todo) *helper.Response
	GetTodos() *helper.Response
	GetTodoByid(id int) *helper.Response
	UpdateTodo(Todo Todo) *helper.Response
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTodo(todo Todo) *helper.Response {
	err := r.db.Debug().Create(&todo).Error
	if err != nil {
		return helper.InternalServerError(err)
	}
	return helper.SuccessCreateResponse(todo, "Succes create todo task")
}

func (r *repository) GetTodos() *helper.Response {
	t := make([]Todo, 0)

	err := r.db.Debug().Find(&t).Error
	if err != nil {
		return helper.InternalServerError(err)
	}
	return helper.SuccessResponse(t, "Success get all todo task")
}

func (r *repository) GetTodoByid(id int) *helper.Response {
	var t Todo

	t.Id = id
	err := r.db.Debug().First(&t).Error
	if err != nil {
		return helper.InternalServerError(err)
	}
	return helper.SuccessResponse(t, "Success get todo")
}

func (r *repository) UpdateTodo(todo Todo) *helper.Response {
	_ = r.GetTodoByid(todo.Id)

	err := r.db.Debug().Updates(&todo).Error
	if err != nil {
		return helper.InternalServerError(err)
	}
	return helper.SuccessResponse(todo, "Success update todo")
}
