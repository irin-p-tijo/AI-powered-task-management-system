package controller

import (
	"errors"
	controller "zocket-task/pkg/controller/interfaces"
	interfaces "zocket-task/pkg/datalayer/interfaces"
	"zocket-task/pkg/utils/model"
)

type TaskController struct {
	taskDataLayer interfaces.TaskDL
	authDL        interfaces.AuthDataLayer
}

func NewTaskController(data interfaces.TaskDL, authData interfaces.AuthDataLayer) controller.TaskController {
	return &TaskController{
		taskDataLayer: data,
		authDL:        authData,
	}
}
func (tc *TaskController) CreateTask(task model.Task) (model.Task, error) {
	newTask, err := tc.taskDataLayer.CreateTask(task)
	if err != nil {
		return model.Task{}, errors.New("Could not create task")
	}
	return newTask, nil
}
func (tc *TaskController) GetTasksByUser(email string) ([]model.Task, error) {
	tasks, err := tc.taskDataLayer.GetTasksByUser(email)
	if err != nil {
		return []model.Task{}, err
	}
	return tasks, nil
}

func (tc *TaskController) AssignTask(taskID int, assignee string, assigner string) error {
	return tc.taskDataLayer.AssignTask(taskID, assignee, assigner)
}
