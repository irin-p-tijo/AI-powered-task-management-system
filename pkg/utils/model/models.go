package model

import "time"

type UserDetails struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required" validate:"email"`
	Phone           string `json:"phone" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}
type TokenUsers struct {
	Users UserDetails
	Token string
}
type UserLogin struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required"`
}
type UserLoginResponse struct {
	Id       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
type Task struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	AssignedTo  string    `json:"assigned_to"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AssignTaskInput struct {
	TaskID     int    `json:"task_id"`
	AssignedTo string `json:"assigned_to"`
}
type Users struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	task     []Task
}
type PromptRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}
