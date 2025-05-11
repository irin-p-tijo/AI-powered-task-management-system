package handlers

import (
	"context"
	"net/http"
	"zocket-task/pkg/config"
	controllers "zocket-task/pkg/controller/interfaces"
	"zocket-task/pkg/utils/model"
	"zocket-task/pkg/utils/response"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

type TaskHandler struct {
	taskController controllers.TaskController
	config         config.Config
}

func NewTaskHandler(taskController controllers.TaskController, cfg config.Config) *TaskHandler {
	return &TaskHandler{
		taskController: taskController,
		config:         cfg,
	}
}

//create task

func (t *TaskHandler) CreateTask(c *gin.Context) {

	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	emailVal, exists := c.Get("email")
	if !exists {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized: email not found in context", nil, "Email not present in token context")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	task.CreatedBy = emailVal.(string)
	task.AssignedTo = emailVal.(string)
	taskDetails, err := t.taskController.CreateTask(task)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "User could not signed up", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "User successfully signed up", taskDetails, nil)
	c.JSON(http.StatusCreated, successRes)
}

// track tasks
func (t *TaskHandler) TrackTask(c *gin.Context) {
	emailVal, exists := c.Get("email")
	if !exists {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized: email not found in context", nil, "Email not present in token context")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	tasks, err := t.taskController.GetTasksByUser(emailVal.(string))
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "User could not signed up", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the task data is retrived", tasks, nil)
	c.JSON(http.StatusOK, successRes)
}

// assign tasks
func (t *TaskHandler) AssignTask(c *gin.Context) {
	var req model.AssignTaskInput
	if err := c.ShouldBindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	assignerEmail, exists := c.Get("email")
	if !exists {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized: email not found in context", nil, "Email not present in token context")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	err := t.taskController.AssignTask(req.TaskID, req.AssignedTo, assignerEmail.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the task data is retrived", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (h *TaskHandler) GetTaskSuggestions(c *gin.Context) {
	var prompt model.PromptRequest
	if err := c.ShouldBindJSON(&prompt); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid prompt format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	suggestions, err := h.generateSuggestions(prompt.Prompt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "Failed to get suggestions", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Task suggestions generated successfully", suggestions, nil)
	c.JSON(http.StatusOK, successRes)
}

func (h *TaskHandler) generateSuggestions(prompt string) (string, error) {
	client := openai.NewClient(h.config.OpenAIAPIKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
