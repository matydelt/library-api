package configs

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", //redis port
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
