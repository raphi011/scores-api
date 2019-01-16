package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/services"
)

type debugHandler struct {
	userService *services.User
}

func (a *debugHandler) postCreateAdmin(c *gin.Context) {
	testEmail := "admin@scores.network"
	testPassword := "test123"

	_, err := a.userService.ByEmail(testEmail)

	if errors.Cause(err) == scores.ErrNotFound {
		_, err = a.userService.New(testEmail, testPassword)

		if err != nil {
			responseErr(c, err)
			return
		}
	} else if err != nil {
		responseErr(c, err)
		return
	}

	responseNoContent(c)
}
