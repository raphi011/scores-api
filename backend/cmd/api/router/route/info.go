package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InfoHandler(version string) Info {
	return Info{version: version}
}

type Info struct {
	version string
}

func (h *Info) GetVersion(c *gin.Context) {
	response(c, http.StatusOK, h.version)
}
