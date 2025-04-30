package domain

import "time"

type Tasks struct {
	Id          int       `json:"id" gorm:"primary key,not null"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AssignedTo  string    `json:"assigned_to"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
