package interfaces

import "zocket-task/pkg/utils/model"

type AuthDataLayer interface {
	CheckUserAvailability(email string) bool
	UserSignUp(user model.UserDetails) (model.UserDetails, error)
	UserLogin(user model.UserDetails) (model.UserDetails, error)
	FindUserByEmail(user model.UserLogin) (model.UserLoginResponse, error)
}
