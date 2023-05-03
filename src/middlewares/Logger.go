package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is a middleware that logs the request
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		fmt.Printf("%s %s %s %s %s\n",
			c.Request.Method,
			c.Request.RequestURI,
			c.Writer.Status(),
			latency,
			c.Request.RemoteAddr,
		)
	}
}
