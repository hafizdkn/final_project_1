package todo

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateTodo(todo Todo) (Todo, error)
	GetTodos() ([]Todo, error)
	GetTodoByid(id int) (Todo, error)
	UpdateTodo(Todo Todo) (Todo, error)
	DeleteTodo(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTodo(todo Todo) (Todo, error) {
	err := r.db.Debug().Create(&todo).Error
	if err != nil {
		return todo, nil
	}

	return todo, nil
}

func (r *repository) GetTodos() ([]Todo, error) {
	t := make([]Todo, 0)

	err := r.db.Debug().Find(&t).Error
	if err != nil {
		return t, err
	}

	return t, nil
}

func (r *repository) GetTodoByid(id int) (Todo, error) {
	var t Todo

	err := r.db.Debug().Where("id=?", id).First(&t).Error
	if err != nil {
		return t, err
	}

	return t, nil
}

func (r *repository) UpdateTodo(todo Todo) (Todo, error) {
	todoId := todo.Id
	_, err := r.GetTodoByid(todoId)
	if err != nil {
		return todo, err
	}

	err = r.db.Debug().Updates(&todo).Error
	if err != nil {
		return todo, err
	}

	return r.GetTodoByid(todoId)
}

func (r *repository) DeleteTodo(id int) error {
	_, err := r.GetTodoByid(id)
	if err != nil {
		return err
	}

	err = r.db.Debug().Delete(&Todo{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
