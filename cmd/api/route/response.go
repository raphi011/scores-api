package route

import (
	"encoding/json"
	"net/http"
	"strconv"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores-api"
	"github.com/raphi011/scores-api/cmd/api/logger"
)

func responseBadRequest(c *gin.Context) {
	writeResponse(c, http.StatusBadRequest, nil, "Bad request")
}

func responseErr(c *gin.Context, err error) {
	code, message := extractErrorInformation(err)

	writeResponse(c, code, nil, message)

	if code >= 500 {
		// log server errors
		logger.Get(c).Warnf("an error occured: %v", err)
	}
}

func extractErrorInformation(err error) (code int, message string) {
	code = http.StatusInternalServerError
	message = err.Error()

	if errors.Is(err, scores.ErrNotFound) {
		code = http.StatusNotFound
	} else if errors.Is(err, scores.ErrorUnauthorized) {
		code = http.StatusUnauthorized
	}

	if code == http.StatusInternalServerError {
		// redact server errors for security reasons
		message = ""
	}

	return
}

func response(c *gin.Context, code int, data interface{}) {
	writeResponse(c, code, data, "")
}

func responseNoContent(c *gin.Context) {
	writeResponse(c, http.StatusNoContent, nil, "")
}

func writeResponse(c *gin.Context, code int, data interface{}, message string) {
	out, _ := json.Marshal(gin.H{
		"status":  code,
		"message": message,
		"data":    data,
	})

	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Writer.Header().Set("Content-Length", strconv.Itoa(len(out)))

	c.Writer.WriteHeader(code)
	c.Writer.Write(out)

}
