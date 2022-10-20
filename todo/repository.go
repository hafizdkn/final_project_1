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
	DeleteTodo(id int) *helper.Response
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

	todoFormatter := TodoFormatter(todo)
	return helper.SuccessCreateResponse(todoFormatter, "Success create todo task")
}

func (r *repository) GetTodos() *helper.Response {
	t := make([]Todo, 0)

	err := r.db.Debug().Find(&t).Error
	if err != nil {
		return helper.InternalServerError(err)
	}

	todoFormatter := TodoFormatter(t)
	return helper.SuccessResponse(todoFormatter, "Success get all todo task")
}

func (r *repository) GetTodoByid(id int) *helper.Response {
	var t Todo

	t.Id = id
	err := r.db.Debug().First(&t).Error
	if err != nil {
		return helper.InternalServerError(err)
	}

	todoFormatter := TodoFormatter(t)
	return helper.SuccessResponse(todoFormatter, "Success get todo")
}

func (r *repository) UpdateTodo(todo Todo) *helper.Response {
	if err := r.checkFirstRecord(Todo{Id: todo.Id}); err != nil {
		return helper.InternalServerError(err)
	}

	err := r.db.Debug().Updates(&todo).Error
	if err != nil {
		return helper.InternalServerError(err)
	}

	todoFormatter := TodoFormatter(todo)
	return helper.SuccessResponse(todoFormatter, "Success update todo")
}

func (r *repository) DeleteTodo(id int) *helper.Response {
	if err := r.checkFirstRecord(Todo{Id: id}); err != nil {
		return helper.InternalServerError(err)
	}

	err := r.db.Debug().Delete(&Todo{}, id).Error
	if err != nil {
		return helper.InternalServerError(err)
	}
	return helper.SuccessResponse(nil, "Success delete todo")
}

func (r *repository) checkFirstRecord(todo Todo) error {
	err := r.db.Debug().First(&todo).Error
	if err != nil {
		return err
	}
	return nil
}
