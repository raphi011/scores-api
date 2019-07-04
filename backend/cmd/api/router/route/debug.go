package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/services"
)

func DebugHandler(userService *services.User) Debug {
	return Debug{userService: userService}
}

type Debug struct {
	userService *services.User
}

func (a *Debug) PostCreateAdmin(c *gin.Context) {
	testEmail := "admin@scores.network"
	testPassword := "test123"

	_, err := a.userService.ByEmail(testEmail)

	if errors.Cause(err) == scores.ErrNotFound {
		_, err = a.userService.New(testEmail, testPassword, "admin")

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
