package database

import (
	"pkgs/config"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func InitRedis() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.host") + ":" + config.GetString("redis.port"),
		Password: config.GetString("redis.password"),
		DB:       config.GetInt("redis.db"),
	})

	return RDB.Ping(RDB.Context()).Err()
}
