package datasource

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var Redis *redis.Client

func NewRedisClient() error {
	address := viper.GetString("redis.address")
	password := viper.GetString("redis.password")

	Redis = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0, // use default DB
	})

	_, err := Redis.Ping().Result()

	if err != nil {
		return err
	}

	fmt.Printf("%v \n", "redis ready!")
	return nil
}
