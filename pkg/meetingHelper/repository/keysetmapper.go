package repository

import (
	"github.com/go-redis/redis"
	"github.com/siller174/meetingHelper/pkg/logger"
	"github.com/siller174/meetingHelper/pkg/utils/http/errors"
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
	request := repo.redisClient.SIsMember(key, value)
	res, err := request.Result()
	if err != nil {
		return false, errors.NewInternalErr(err)
	}
	if res {
		logger.Debug("Set %v contain %v.", key, value)
	} else {
		logger.Debug("Set %v is not contain %v.", key, value)
	}
	return res, nil
}

func (repo *KeySetMapper) Add(key string, value string) (bool, error) {
	logger.Debug("Put to set %v value %v", key, value)
	request := repo.redisClient.SAdd(key, value)
	res, err := request.Result()
	if err != nil {
		return false, errors.NewInternalErr(err)
	}
	if res == 1 {
		logger.Debug("Put to set %v value %v. Success. Return %v", key, value, res)
		return true, nil
	}
	return false, nil
}

func (repo *KeySetMapper) Remove(key string, value string) (bool, error) {
	logger.Debug("Remove from set %v value %v", key, value)
	request := repo.redisClient.SRem(key, value)
	res, err := request.Result()
	if err != nil {
		return false, errors.NewInternalErr(err)
	}
	if res == 1 {
		logger.Debug("Remove from set %v value %v. Success. Return %v", key, value, res)
		return true, nil
	}
	return false, nil
}
