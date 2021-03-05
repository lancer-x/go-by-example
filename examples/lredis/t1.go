package main

import (
	"github.com/go-redis/redis"
	"time"
)

var redisdb *redis.Client


func main()  {
	initRedis()
	tSet()
}

func initRedis() error {
	redisdb = redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               "localhost:6379",
		Dialer:             nil,
		OnConnect:          nil,
		Password:           "",
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})
	_, err := redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func tSet()  {
	redisdb.Set("aaaaaa", 100, 1000 * time.Second)
}




