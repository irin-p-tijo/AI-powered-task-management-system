package domain

type Users struct {
	ID       int    `json:"id" gorm:"primary key,not null"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8,max=20"`
	Phone    string `json:"phone"`
}
