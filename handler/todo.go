package handler

import (
	"github.com/gin-gonic/gin"

	"final_project_1/helper"
	"final_project_1/todo"
)

type todoHandler struct {
	todoService todo.Service
}

func NewHandler(todoService todo.Service) *todoHandler {
	return &todoHandler{todoService}
}

func (h *todoHandler) CreateTodo(c *gin.Context) {
	var input todo.TodoInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		helper.WriteJsonRespnse(c, helper.BadRequestResponse(err))
		return
	}
	resp := h.todoService.CreateTodo(input)
	helper.WriteJsonRespnse(c, resp)
}

func (h *todoHandler) GetTodos(c *gin.Context) {
	resp := h.todoService.GetTodos()
	helper.WriteJsonRespnse(c, resp)
}
