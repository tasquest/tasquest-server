package tasks

import "github.com/google/uuid"

const AdventurerTaskTopic = "event.adventurer.completed.task"

type AdventurerCompletedTaskEvent struct {
	AdventurerID      uuid.UUID `json:"adventurerId" binding:"required"`
	TaskId            uuid.UUID `json:"taskId" binding:"required"`
	ExperienceAwarded int64     `json:"experienceAwarded"  binding:"required"`
}
