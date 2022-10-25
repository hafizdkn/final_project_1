package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"final_project_1/database"
	"final_project_1/handler"
	"final_project_1/todo"
)

func main() {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.Generate(db)

	todoRepo := todo.NewRepository(db)
	todoService := todo.NewService(todoRepo)
	todoHandler := handler.NewHandler(todoService)

	app := handler.NewRouter(todoHandler)
	app.Run(":8080")
}

