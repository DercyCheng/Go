package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func redisConnect() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "kwok2012",
		DB:       0, // use default DB
	})
	fmt.Println("connect success")
	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rdb)
}

func main() {
	redisConnect()
}
