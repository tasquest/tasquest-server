package tasks

type CreateTaskCommand struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Experience  int64  `json:"experience" binding:"required"`
}

type UpdateTaskCommand struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Experience  int64  `json:"experience" binding:"required"`
}

type AdventurerCompletedTaskCommand struct {
	AdventurerID string `json:"adventurer_id" binding:"required"`
	TaskID       string `json:"task_id" binding:"required"`
}
