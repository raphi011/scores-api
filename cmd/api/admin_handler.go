package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores/db/sqlite"
	"golang.org/x/oauth2"
)

type adminHandler struct {
	userRepository  *sqlite.UserRepository
	groupRepository *sqlite.GroupRepository
	conf         *oauth2.Config
}

func (a *adminHandler) createPlayer(c *gin.Context) {

}
