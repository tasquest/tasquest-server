package tasks

import "github.com/google/uuid"

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
	AdventurerID uuid.UUID `json:"adventurer_id" binding:"required"`
	TaskID       uuid.UUID `json:"task_id" binding:"required"`
}
