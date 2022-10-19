package todo

type TodoInput struct {
	Task string `json:"task" binding:"required"`
}

type TodoUpdateInput struct {
	Task      string `json:"task" binding:"required"`
	Completed bool   `json:"completed" binding:"required"`
}
