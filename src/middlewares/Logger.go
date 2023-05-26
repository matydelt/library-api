package middlewares

import (
	"context"
	"fmt"
	"library/src/configs"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is a middleware that logs the request
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		redis := configs.RedisConnect()
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		// fmt.Println(c.W)
		err := redis.Set(
			context.Background(),
			fmt.Sprint("", c.Request.Method, c.Request.RequestURI, c.Writer.Status(), latency),
			fmt.Sprint("", c.Request.Method, c.Request.RequestURI, c.Writer.Status()),
			2*time.Minute,
		).Err()
		if err != nil {
			// fmt.Println(err.Error())
			panic(err)
		}
	}
}
