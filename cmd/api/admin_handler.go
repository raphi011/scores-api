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
	var user postUserDto

	if err := c.ShouldBindWith(&user, binding.JSON); err != nil {
		responseBadRequest(c)
		return
	}

	_, err := a.userService.ByEmail(user.Email)

	if err == scores.ErrorNotFound {
		a.userService.New(user.Email, user.Password)
	} else if err != nil {
		responseErr(c, err)
	}

}

func (a *adminHandler) getUsers(c *gin.Context) {
	users, err := a.userService.All()

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, users)
}
