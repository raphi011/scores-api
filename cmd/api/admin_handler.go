package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/raphi011/scores"
)

type adminHandler struct {
	userService *scores.UserService
}

type postUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *adminHandler) postUser(c *gin.Context) {
	var userChanges postUserDto

	if err := c.ShouldBindWith(&userChanges, binding.JSON); err != nil {
		responseBadRequest(c)
		return
	}

	user, err := a.userService.ByEmail(userChanges.Email)

	if err == scores.ErrorNotFound {
		user, err = a.userService.New(userChanges.Email, userChanges.Password)

		if err != nil {
			responseErr(c, err)
			return
		}

		response(c, http.StatusCreated, user)
	} else if err != nil {
		responseErr(c, err)
		return
	}

	err = a.userService.SetPassword(user.ID, userChanges.Password)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, user)
}

func (a *adminHandler) getUsers(c *gin.Context) {
	users, err := a.userService.All()

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, users)
}
