package todo

type TodoInput struct {
	Task string `json:"task" binding:"required"`
}
