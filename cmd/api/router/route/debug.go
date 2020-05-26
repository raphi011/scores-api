package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/raphi011/scores-api"
	"github.com/raphi011/scores-api/services"
)

// DebugHandler is the constructor for the debug routes handler.
func DebugHandler(userService *services.User) Debug {
	return Debug{userService: userService}
}

// Debug wraps the dependencies of the DebugHandler.
type Debug struct {
	userService *services.User
}

// PostCreateAdmin creates a user with the admin role
// for debugging purposes.
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
