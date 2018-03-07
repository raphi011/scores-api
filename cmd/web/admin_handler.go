package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores/sqlite"
	"golang.org/x/oauth2"
)

type adminHandler struct {
	userService  *sqlite.UserService
	groupService *sqlite.GroupService
	conf         *oauth2.Config
}

func (a *adminHandler) createPlayer(c *gin.Context) {

}
