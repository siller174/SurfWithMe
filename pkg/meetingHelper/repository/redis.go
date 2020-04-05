package repository

import (
	"fmt"
	"github.com/siller174/meetingHelper/pkg/utils/http/errors"
	"os"

	"github.com/go-redis/redis"
	"github.com/siller174/meetingHelper/pkg/logger"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/config"
)

type KeyListMapper struct {
	redisClient *redis.Client
}

func New(redisConfig config.Redis) KeyListMapper {
	res := &KeyListMapper{
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

func (repo *KeyListMapper) Ping() bool {
	_, err := repo.redisClient.Ping().Result()
	if err != nil {
		logger.Error("Redis is not available")
		return false
	}
	return true
}

func (repo *KeyListMapper) Put(key string, value string) error {
	logger.Debug("Put key %v value %v", key, value)
	put := repo.redisClient.RPush(key, value)
	i, err := put.Result()
	if err != nil {
		return err
	}
	logger.Debug("Put key %v value %v. Success. Return %v", key, value, i)
	return nil
}

func (repo *KeyListMapper) GetLast(key string) (*string, error) {
	logger.Debug("GetLast %v", key)
	get := repo.redisClient.LRange(key, -1, -1)
	meetingJson, err := get.Result()
	if err != nil {
		return nil, err
	}
	if len(meetingJson) < 1 {
		return nil,  errors.NewNotFound(fmt.Errorf("get values from %v not found", key))
	}
	logger.Debug("GetLast key %v. Success. Return %v", key, meetingJson)
	return &meetingJson[0], nil
}

func (repo *KeyListMapper) GetAll(key string) (*[]string, error) {
	logger.Debug("GetAll by %v", key)
	lRange := repo.redisClient.LRange(key, 0, -1)
	meetingJson, err := lRange.Result()
	if err != nil {
		return nil, err
	}
	if len(meetingJson) < 1 {
		return nil, errors.NewNotFound(fmt.Errorf("history from %v not found", key))
	}
	logger.Debug("GetAll by %v. Success. Return %v", key, meetingJson)
	return &meetingJson, nil
}