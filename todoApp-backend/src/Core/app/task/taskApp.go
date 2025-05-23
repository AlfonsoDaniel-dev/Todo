package task

import "todoApp-backend/src/Core/domain"

type TaskServices struct {
	taskRepository domain.TaskRepository
}

func NewTaskServices(taskRepository domain.TaskRepository) *TaskServices {
	return &TaskServices{taskRepository}
}
