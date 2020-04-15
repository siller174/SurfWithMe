package structs

import "github.com/go-redis/redis"

type Health struct {
	RedisClient *redis.Client
}
