package main

import (
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user-id")

		if userID == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user-id", userID)
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

func localAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		host, _, err := net.SplitHostPort(c.Request.RemoteAddr)

		if err != nil {
			log.Print(err)
		}

		ip := net.ParseIP(host)

		if !ip.IsLoopback() && !privateIP(ip) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
