package handlers

import (
	"net/http"
	controllers "zocket-task/pkg/controller/interfaces"
	"zocket-task/pkg/utils/model"
	"zocket-task/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authController controllers.AuthController
}

func NewAuthHandler(authController controllers.AuthController) *AuthHandler {
	return &AuthHandler{
		authController: authController,
	}
}

func (a *AuthHandler) UserSignup(c *gin.Context) {

	var user model.UserDetails

	if err := c.ShouldBindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	userSignup, err := a.authController.UserSignUp(user)

	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "User could not signed up", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "User successfully signed up", userSignup, nil)
	c.JSON(http.StatusCreated, successRes)

}
func (a *AuthHandler) UserLogin(c *gin.Context) {
	var user model.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	userDetails, err := a.authController.UserLogin(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "User could not be logged in", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "User successfully logged in", userDetails, nil)
	c.JSON(http.StatusOK, successRes)

}
