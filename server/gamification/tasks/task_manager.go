package tasks

type TaskManager interface {
	CreateTask(command CreateTaskCommand) (Task, error)
	UpdateTask(taskID string, command UpdateTaskCommand) (Task, error)
	DeleteTask(taskID string) (Task, error)
	AdventurerCompletesTask(command AdventurerCompletedTaskCommand) error
}
