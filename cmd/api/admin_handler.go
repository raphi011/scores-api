package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores"
)

type adminHandler struct {
	userService *scores.UserService
}

func (a *adminHandler) createPlayer(c *gin.Context) {

}

func (a *adminHandler) getUsers(c *gin.Context) {
	users, err := a.userService.All()

	if err != nil {
		jsonn(c, http.StatusInternalServerError, nil, "")
	}

	jsonn(c, http.StatusOK, users, "")
}

func (a *adminHandler) postUser(c *gin.Context) {

}
