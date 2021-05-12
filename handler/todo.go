package handler

import (
	"go-basic-crud/todo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type todoHandler struct {
	todoService todo.Service
}

func NewTodoHandler(todoService todo.Service) *todoHandler {
	return &todoHandler{todoService}
}

func (handler *todoHandler) Store(c *gin.Context) {
	var input todo.InputTodo
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newTodo, err := handler.todoService.Store(input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := Response{
		Success: true,
		Message: "New todo has benn stored successfully",
		Data:    newTodo,
	}

	c.JSON(http.StatusOK, response)
}
