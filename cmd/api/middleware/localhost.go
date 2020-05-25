package middleware

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores-backend/cmd/api/logger"
)

// LocalhostOnly middleware restricts routes to requests coming from private ip ranges
func LocalhostOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		host, _, err := net.SplitHostPort(c.Request.RemoteAddr)

		if err != nil {
			logger.Get(c).Warnf("cannot split hostport: %v", err)
			return
		}

		ip := net.ParseIP(host)

		if !ip.IsLoopback() && !privateIP(ip) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

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
