package repository

import (
	"github.com/go-redis/redis"
	"github.com/siller174/meetingHelper/pkg/logger"
)

type KeySetMapper struct {
	redisClient *redis.Client
}

func NewKeySetMapper(client *redis.Client) *KeySetMapper {
	return &KeySetMapper{
		redisClient: client,
	}
}

func (repo *KeySetMapper) IsMember(key string, value string) (bool, error) {
	res := repo.redisClient.SIsMember(key, value)
	isMember, err := res.Result()
	if err != nil {
		return false, err
	}
	logger.Debug("Set %v contain %v.", key, value)
	return isMember, nil
}

func (repo *KeySetMapper) Add(key string, value string) (bool, error) {
	logger.Debug("Put to set %v value %v", key, value)
	put := repo.redisClient.SAdd(key, value)
	res, err := put.Result()
	if err != nil {
		return false, err
	}
	if res == 1 {
		logger.Debug("Put to set %v value %v. Success. Return %v", key, value, res)
		return true, nil
	}
	return false, nil
}

func (repo *KeySetMapper) Remove(key string, value string) (bool, error) {
	logger.Debug("Remove from set %v value %v", key, value)
	put := repo.redisClient.SRem(key, value)
	res, err := put.Result()
	if err != nil {
		return false, err
	}
	if res == 1 {
		logger.Debug("Remove from set %v value %v. Success. Return %v", key, value, res)
		return true, nil
	}
	return false, nil
}
