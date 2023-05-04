package middlewares

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// Logger is a middleware that logs the request
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		redis := redisClient()
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
		err := redis.Set(context.Background(), "request", fmt.Sprintf("", c.Request.Method, c.Request.RequestURI, c.Writer.Status(), latency), 1*time.Hour).Err()
		if err != nil {
			fmt.Println(err.Error())
			// panic(err)
		}
	}
}

// create a redis client
func redisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:4000", //redis port
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println(pong)
	}
	return client
}
