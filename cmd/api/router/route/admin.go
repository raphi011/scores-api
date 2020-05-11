package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/services"
)

// AdminHandler is the constructor for the Admin routes handler.
func AdminHandler(userService *services.User) Admin {
	return Admin{userService: userService}
}

// Admin wraps the dependencies of the AdminHandler.
type Admin struct {
	userService *services.User
}

type postUserDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// PostUser handles the user post route that allows an admin to change
// details of a user.
func (a *Admin) PostUser(c *gin.Context) {
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

// GetUsers returns all current users.
func (a *Admin) GetUsers(c *gin.Context) {
	users, err := a.userService.All()

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, users)
}
