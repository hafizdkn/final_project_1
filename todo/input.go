package todo

type TodoInput struct {
	Task string `json:"task" binding:"required" example:"push up 1 sets"`
}

type TodoUpdateInput struct {
	Task      string `json:"task" binding:"required" example:"push up 1 sets"`
	Completed bool   `json:"completed" binding:"required" example:"true"`
}
