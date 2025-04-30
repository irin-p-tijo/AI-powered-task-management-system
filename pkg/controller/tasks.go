package controller

import (
	"encoding/json"
	"log"
	controller "zocket-task/pkg/controller/interfaces"
	interfaces "zocket-task/pkg/datalayer/interfaces"
	"zocket-task/pkg/utils/model"

	"github.com/gorilla/websocket"
)

type TaskController struct {
	taskDataLayer interfaces.TaskDataLayer
}

func NewTaskController(data interfaces.TaskDataLayer) controller.TaskController {
	return &TaskController{
		taskDataLayer: data,
	}
}
func (u *TaskController) CreateTask(task model.TaskDetails) (model.Task, error) {
	return model.Task{}, nil
}
func BroadcastTaskUpdates(broadcast chan model.Task, clients map[*websocket.Conn]bool) {
	for {
		task := <-broadcast
		message, _ := json.Marshal(task)

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("WebSocket error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
