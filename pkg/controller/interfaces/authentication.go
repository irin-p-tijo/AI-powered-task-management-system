package interfaces

import "zocket-task/pkg/utils/model"

type AuthController interface {
	UserSignUp(user model.UserDetails) (model.TokenUsers, error)
	UserLogin(user model.UserLogin) (model.TokenUsers, error)
}
