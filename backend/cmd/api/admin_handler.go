package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/services"
)

type adminHandler struct {
	userService *services.User
}

type postUserDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *adminHandler) postUser(c *gin.Context) {
	var userChanges postUserDto

	if err := c.ShouldBindWith(&userChanges, binding.JSON); err != nil {
		responseBadRequest(c)
		return
	}

	user, err := a.userService.ByEmail(userChanges.Email)

	if errors.Cause(err) == scores.ErrNotFound {
		user, err = a.userService.New(userChanges.Email, userChanges.Password, "user")

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
