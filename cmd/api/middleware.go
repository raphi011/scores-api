package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raphi011/scores"
)

func metricMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		fmt.Printf("Request %s took %s\n", c.Request.URL.String(), time.Since(start))
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger(c)
		session := sessions.Default(c)
		userID := session.Get("user-id")

		if userID == nil {
			log.Print("unauthorized")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user-id", userID)

		log = log.WithFields(logrus.Fields{"user-id": userID})

		c.Set("log", log)

		c.Next()
	}
}

func adminMiddlware(userService *scores.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger(c)
		session := sessions.Default(c)
		userID := session.Get("user-id").(uint)

		if !userService.HasRole(userID, "admin") {
			log.Print("unauthorized")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func loggerMiddleware(log logrus.FieldLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log = log.WithFields(logrus.Fields{
			"method":     c.Request.Method,
			"url":        c.Request.URL.String(),
			"ip":         c.Request.RemoteAddr,
			"user-agent": c.Request.UserAgent(),
			"request-id": uuid.New().String(),
		})

		c.Set("log", log)

		c.Next()
	}
}

func privateIP(ip net.IP) bool {
	_, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
	_, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
	_, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")

	private := private24BitBlock.Contains(ip) || private20BitBlock.Contains(ip) || private16BitBlock.Contains(ip)

	return private
}

func localhostOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		host, _, err := net.SplitHostPort(c.Request.RemoteAddr)

		if err != nil {
			// this really should not happen
		}

		ip := net.ParseIP(host)

		if !ip.IsLoopback() && !privateIP(ip) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
