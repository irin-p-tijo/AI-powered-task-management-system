package datalayer

import (
	"zocket-task/pkg/datalayer/interfaces"
	"zocket-task/pkg/utils/model"

	"gorm.io/gorm"
)

type TaskAuthenticationDL struct {
	DB *gorm.DB
}

func NewTaskAuthenticationDL(DB *gorm.DB) interfaces.AuthDataLayer {
	return &UserAuthenticationDL{DB}
}

func (c *UserAuthenticationDL) CreateTask(task model.TaskDetails) (model.Task, error) {
	return model.Task{}, nil
}
