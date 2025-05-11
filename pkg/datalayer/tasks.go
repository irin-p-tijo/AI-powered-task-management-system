package datalayer

import (
	"time"
	"zocket-task/pkg/datalayer/interfaces"
	"zocket-task/pkg/utils/model"

	"gorm.io/gorm"
)

type TaskDL struct {
	DB *gorm.DB
}

func NewTaskDL(DB *gorm.DB) interfaces.TaskDL {
	return &TaskDL{DB}
}

func (td *TaskDL) CreateTask(task model.Task) (model.Task, error) {
	now := time.Now()
	query := `INSERT INTO tasks (title, description, created_by,assigned_to, created_at, updated_at) 
	          VALUES (?, ?,?, ?, ?, ?) RETURNING *`

	var createdTask model.Task
	err := td.DB.Raw(query, task.Title, task.Description, task.CreatedBy,task.AssignedTo, now, now).Scan(&createdTask).Error
	if err != nil {
		return model.Task{}, err
	}
	return createdTask, nil
}
//tbd
func (td *TaskDL) CheckTasks(task model.Task) (model.Task, error) {
	return model.Task{}, nil
}
//tbd
func (td *TaskDL) UpdateTasks(task model.Task) (model.Task, error) {
	return model.Task{}, nil
}
func (td *TaskDL) GetTasksByUser(email string) ([]model.Task, error) {
	var tasks []model.Task
	query := `select * from tasks where created_by=? ORDER BY created_at ASC`
	err := td.DB.Raw(query, email).Scan(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
func (td *TaskDL) AssignTask(taskID int, assignee string, assigner string) error {
	result := td.DB.Model(&model.Task{}).Where("id = ? AND created_by = ?", taskID, assigner).
		Updates(map[string]interface{}{
			"assigned_to": assignee,
			"updated_at":  time.Now(),
		})
	return result.Error
}
