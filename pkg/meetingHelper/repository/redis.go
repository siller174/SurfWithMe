package repository

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/siller174/meetingHelper/pkg/logger"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/config"
)

func New(redisConfig config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.DB})

	_, err := client.Ping().Result()
	if err != nil {
		logger.Error("Redis is not available")
		os.Exit(-1)
	}
	return client
}
