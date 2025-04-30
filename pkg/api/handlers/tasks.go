package handlers

import (
	controllers "zocket-task/pkg/controller/interfaces"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskController controllers.TaskController
}

func NewTaskHandler(taskController controllers.TaskController) *TaskHandler {
	return &TaskHandler{
		taskController: taskController,
	}
}
func (t *TaskHandler) CreateTask(c *gin.Context) {

}
func (t *TaskHandler) GetAllTask(c *gin.Context) {

}
func (t *TaskHandler) AssignTask(c *gin.Context) {

}

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool { return true },
// }

// var clients = make(map[*websocket.Conn]bool)
// var broadcast = make(chan model.Task)

// func InitWebSocketHandler(r *gin.Engine) {
// 	r.GET("/ws", wsHandler)
// 	go controllers.BroadcastTaskUpdates(broadcast, clients)
// }

// func wsHandler(c *gin.Context) {
// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Println("Upgrade error:", err)
// 		return
// 	}
// 	clients[conn] = true
// }
