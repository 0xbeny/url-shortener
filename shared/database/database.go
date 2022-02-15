package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Database struct {
	Client *redis.Client
}

var (
	Ctx      = context.TODO()
	database *Database
)

func Connect(addr string) {
	client := redis.NewClient(&redis.Options{
		DB:   0,
		Addr: addr,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		panic(err)
	}
	fmt.Printf("[connected] redis connection is available on %s\n", addr)
	database = &Database{
		Client: client,
	}
}
func GetInstance() *Database {
	return database
}
