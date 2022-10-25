package handler

import (
	"final_project_1/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type router struct {
	router      *gin.Engine
	todoHandler *todoHandler
}

func NewRouter(todoHandler *todoHandler) *router {
	r := gin.Default()
	return &router{router: r, todoHandler: todoHandler}
}

func (r *router) Run(port string) {
	docs.SwaggerInfo.Title = "Todo API"
	docs.SwaggerInfo.Description = "This is a sample Todo server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "todo.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.router.POST("/todos", r.todoHandler.CreateTodo)
	r.router.GET("/todos", r.todoHandler.GetTodos)
	r.router.GET("/todos/:id", r.todoHandler.GetTodoByid)
	r.router.PUT("/todos/:id", r.todoHandler.UpdateTodo)
	r.router.DELETE("/todos/:id", r.todoHandler.DeleteTodo)

	r.router.Run(port)
}
