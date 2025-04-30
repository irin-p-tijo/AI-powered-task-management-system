package interfaces

import "zocket-task/pkg/utils/model"

type TaskController interface {
	CreateTask(task model.Task) (model.Task, error)
	GetTasksByUser(email string) ([]model.Task, error)
	AssignTask(taskID int, assignee string, assigner string) error
}
