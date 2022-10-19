package handler

import "github.com/gin-gonic/gin"

type router struct {
	router      *gin.Engine
	todoHandler *todoHandler
}

func NewRouter(todoHandler *todoHandler) *router {
	r := gin.Default()
	return &router{router: r, todoHandler: todoHandler}
}

func (r *router) Run(port string) {
	r.router.POST("/todos", r.todoHandler.CreateTodo)

	r.router.Run(port)
}
