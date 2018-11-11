package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/repo/sqlite"
	"golang.org/x/oauth2"
)

type adminHandler struct {
	userRepository  scores.UserService
	groupRepository *sqlite.GroupRepository
	conf            *oauth2.Config
}

func (a *adminHandler) createPlayer(c *gin.Context) {

}
