package repository

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/siller174/meetingHelper/pkg/logger"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/config"
)

type RedisRepo struct {
	redisClient *redis.Client
}

func New(redisConfig config.Redis) RedisRepo {
	res := &RedisRepo{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     redisConfig.Address,
			Password: redisConfig.Password,
			DB:       redisConfig.DB}),
	}
	if !res.Ping() {
		logger.Error("Shutdown")
		os.Exit(-1)
	}
	return *res
}

func (db *RedisRepo) Ping() bool {
	_, err := db.redisClient.Ping().Result()
	if err != nil {
		logger.Error("Redis is not available")
		return false
	}
	return true
}
