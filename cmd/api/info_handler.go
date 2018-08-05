package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type infoHandler struct {
}

func (h *infoHandler) version(c *gin.Context) {
	jsonn(c, http.StatusOK, version, "")
}
