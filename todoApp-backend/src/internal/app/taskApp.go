package app

import "todoApp-backend/src/internal/domain"

type TaskServices struct {
	taskRepository domain.TaskRepository
}

func NewTaskServices(taskRepository domain.TaskRepository) *TaskServices {
	return &TaskServices{taskRepository}
}
