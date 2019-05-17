package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func NewConnection(host, port, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})

	if _, err := client.Ping().Result(); err != nil {
		fmt.Println("Error connecting to cache", err)
		os.Exit(1)
	}

	return client
}
