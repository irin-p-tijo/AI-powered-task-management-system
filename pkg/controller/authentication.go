package controller

import (
	"errors"
	"log"

	controller "zocket-task/pkg/controller/interfaces"
	interfaces "zocket-task/pkg/datalayer/interfaces"
	"zocket-task/pkg/utils/auth"
	"zocket-task/pkg/utils/model"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	authDataLayer interfaces.AuthDataLayer
}

func NewAuthController(data interfaces.AuthDataLayer) controller.AuthController {
	return &AuthController{
		authDataLayer: data,
	}
}

func (u *AuthController) UserSignUp(user model.UserDetails) (model.TokenUsers, error) {
	userExists := u.authDataLayer.CheckUserAvailability(user.Email)
	if userExists {
		return model.TokenUsers{}, errors.New("User already existing")
	}

	if user.Password != user.ConfirmPassword {
		return model.TokenUsers{}, errors.New("password does not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Println("Error hashing password:", err)
		return model.TokenUsers{}, err
	}

	user.Password = string(hashedPassword)

	userData, err := u.authDataLayer.UserSignUp(user)
	if err != nil {
		return model.TokenUsers{}, errors.New("Error in storing data")
	}

	token, err := auth.GenerateJWT(user.Email)
	if err != nil {
		log.Println("Error generating JWT token:", err)
		return model.TokenUsers{}, err
	}
	var userDetails model.UserDetails
	err = copier.Copy(&userDetails, &userData)
	if err != nil {
		return model.TokenUsers{}, err
	}
	return model.TokenUsers{
		Users: userDetails,
		Token: token,
	}, nil
}
func (u *AuthController) UserLogin(user model.UserLogin) (model.TokenUsers, error) {
	userExists := u.authDataLayer.CheckUserAvailability(user.Email)
	if !userExists {
		return model.TokenUsers{}, errors.New("the user does not exist")
	}

	userDetails, err := u.authDataLayer.FindUserByEmail(user)
	if err != nil {
		return model.TokenUsers{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(user.Password))
	if err != nil {
		return model.TokenUsers{}, errors.New("password incorrect")
	}
	var userData model.UserDetails
	err = copier.Copy(&userDetails, &userDetails)
	if err != nil {
		return model.TokenUsers{}, err
	}
	token, err := auth.GenerateJWT(user.Email)
	if err != nil {
		log.Println("Error generating JWT token:", err)
		return model.TokenUsers{}, err
	}
	return model.TokenUsers{
		Users: userData,
		Token: token,
	}, nil

}
