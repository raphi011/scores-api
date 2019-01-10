package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Metric middleware logs how long a request took
func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		fmt.Printf("Request %s took %s\n", c.Request.URL.String(), time.Since(start))
	}
}
