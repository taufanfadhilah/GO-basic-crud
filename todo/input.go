package todo

type InputTodo struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
