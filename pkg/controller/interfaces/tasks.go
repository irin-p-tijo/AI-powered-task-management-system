package interfaces

import "zocket-task/pkg/utils/model"

type TaskController interface {
	CreateTask(task model.TaskDetails) (model.Task, error)
}
