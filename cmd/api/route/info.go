package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InfoHandler handles the info routes.
func InfoHandler(version string) Info {
	return Info{version: version}
}

// Info wraps the dependencies of the InfoHandler.
type Info struct {
	version string
}

// GetVersion serves the version of the API.
func (h *Info) GetVersion(c *gin.Context) {
	response(c, http.StatusOK, h.version)
}
