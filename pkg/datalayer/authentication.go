package datalayer

import (
	"errors"
	"fmt"
	"zocket-task/pkg/datalayer/interfaces"
	"zocket-task/pkg/utils/model"

	"gorm.io/gorm"
)

type UserAuthenticationDL struct {
	DB *gorm.DB
}

func NewUserAuthenticationDL(DB *gorm.DB) interfaces.AuthDataLayer {
	return &UserAuthenticationDL{DB}
}

func (c *UserAuthenticationDL) CheckUserAvailability(email string) bool {

	var count int
	query := fmt.Sprintf("select count(*) from users where email='%s'", email)
	if err := c.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0

}

func (c *UserAuthenticationDL) UserSignUp(user model.UserDetails) (model.UserDetails, error) {

	var userDetails model.UserDetails
	err := c.DB.Raw(`INSERT INTO users (name, email, phone, password) VALUES ($1, $2, $3, $4) RETURNING id, name, email, phone`, user.Name, user.Email, user.Phone, user.Password).Scan(&userDetails).Error

	if err != nil {
		return model.UserDetails{}, err
	}

	return userDetails, nil
}
func (c *UserAuthenticationDL) UserLogin(user model.UserDetails) (model.UserDetails, error) {

	var userResponse model.UserDetails
	err := c.DB.Save(&userResponse).Error
	return userResponse, err

}

func (c *UserAuthenticationDL) FindUserByEmail(user model.UserLogin) (model.UserLoginResponse, error) {

	var userDetails model.UserLoginResponse

	err := c.DB.Raw(`
		select * from users where email = ? and blocked = false	`, user.Email).Scan(&userDetails).Error

	if err != nil {
		return model.UserLoginResponse{}, errors.New("error checking user details")
	}

	return userDetails, nil

}
