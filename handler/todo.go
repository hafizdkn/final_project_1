package handler

import (
	"strconv"

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

// Create todo godoc
// @Summary Create new todo task
// @Description Create todo task
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body todo.TodoInput true "param to create todo"
// @Success 200 {array} todo.Formatter
// @Failure 400 {array} helper.Response
// @Router /user/person/login [post]
// @Router /todos [post]
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

// Get Todo godoc
// @Summary Get all of todo task
// @Description Get detail of all todo task
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} todo.Formatter
// @Router /todos [get]
func (h *todoHandler) GetTodos(c *gin.Context) {
	resp := h.todoService.GetTodos()
	helper.WriteJsonRespnse(c, resp)
}

// Get Todo By Id godoc
// @Summary Get todo by id
// @Description Get detail todo task by id
// @Tags todos
// @Accept json
// @Produce json
// @Param id body int true "id todo task"
// @Success 200 {array} todo.Formatter
// @Failure 500 {array} helper.Response "Internal server error, When id not in record"
// @Router /todos/:id [get]
func (h *todoHandler) GetTodoByid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.WriteJsonRespnse(c, helper.BadRequestResponse(err))
		return
	}
	resp := h.todoService.GetTodoByid(id)
	helper.WriteJsonRespnse(c, resp)
}

// Update Todo By Id godoc
// @Summary Update todo by id
// @Description update doto by id
// @Tags todos
// @Accept json
// @Produce json
// @Param param body todo.TodoUpdateInput true "param to update todo task"
// @Success 200 {array} todo.Formatter
// @Failure 500 {array} helper.Response "Internal server error, When id not in record"
// @Router /todos/:id [put]
func (h *todoHandler) UpdateTodo(c *gin.Context) {
	var t todo.TodoUpdateInput
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.WriteJsonRespnse(c, helper.BadRequestResponse(err))
		return
	}

	err = c.ShouldBindJSON(&t)
	if err != nil {
		helper.WriteJsonRespnse(c, helper.BadRequestResponse(err))
		return
	}

	resp := h.todoService.UpdateTodo(id, t)
	helper.WriteJsonRespnse(c, resp)
}

// Delete Todo By id
// @Summary Delete todo by id
// @Description delete todo by id
// @Tags todos
// @Accept json
// @Produce json
// @Param id body int true "id todo task"
// @Success 200 {array} helper.Response "Success, When id in record"
// @Failure 500 {array} helper.Response "Internal server error, When id not in record"
// @Router /todos/:id [delete]
func (h *todoHandler) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.WriteJsonRespnse(c, helper.BadRequestResponse(err))
	}

	resp := h.todoService.DeleteTodo(id)
	helper.WriteJsonRespnse(c, resp)
}
