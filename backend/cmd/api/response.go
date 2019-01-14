package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/cmd/api/logger"
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

	cause := errors.Cause(err)

	if cause == scores.ErrNotFound {
		code = http.StatusNotFound
	} else if cause == scores.ErrorUnauthorized {
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
