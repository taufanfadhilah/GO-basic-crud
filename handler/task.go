package handler

import (
	"go-basic-crud/task"
	"net/http"

	"github.com/gin-gonic/gin"
)

// set response structure
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

func (h *taskHandler) Store(c *gin.Context) {
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

	newTask, err := h.taskService.Store(input)
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

func (h *taskHandler) Show(c *gin.Context) {
	var input task.InputTaskDetail
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	task, err := h.taskService.Show(input)
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
		Message: "Get task by id",
		Data:    task,
	}
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) Update(c *gin.Context) {
	var inputDetail task.InputTaskDetail
	var input task.InputTask

	err := c.ShouldBindUri(&inputDetail)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newTask, err := h.taskService.Update(inputDetail, input)
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
		Message: "Task updated successfully",
		Data:    newTask,
	}
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) Destory(c *gin.Context) {
	var inputDetail task.InputTaskDetail
	err := c.ShouldBindUri(&inputDetail)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.taskService.Destroy(inputDetail)
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
		Message: "Task has been deleted successfully",
		Data:    nil,
	}

	c.JSON(http.StatusOK, response)
}
