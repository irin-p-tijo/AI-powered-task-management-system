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
type TaskDetails struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	AssignedUserID int       `json:"assigned_user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Status    string `json:"status"` // pending, in_progress, completed
	UpdatedBy string `json:"updated_by"`
}
