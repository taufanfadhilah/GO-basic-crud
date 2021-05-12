package handler

import (
	"go-basic-crud/task"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type taskHandler struct {
	taskService task.Service
}

func NewTaskHandler(taskService task.Service) *taskHandler {
	return &taskHandler{taskService}
}

func (h *taskHandler) Index(c *gin.Context) {
	tasks, err := h.taskService.Index()
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
		Message: "Get all tasks",
		Data:    tasks,
	}
	c.JSON(http.StatusOK, response)
}

func (handler *taskHandler) Store(c *gin.Context) {
	var input task.InputTask
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

	newTask, err := handler.taskService.Store(input)
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
		Message: "New task has benn stored successfully",
		Data:    newTask,
	}

	c.JSON(http.StatusOK, response)
}
