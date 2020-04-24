package service

import (
	"github.com/go-redis/redis"
	"github.com/siller174/meetingHelper/pkg/common/repository"
	"github.com/siller174/meetingHelper/pkg/logger"
)

type NotificationManager struct {
	keyListMapper *repository.KeyListMapper
	keySetMapper  *repository.KeySetMapper
}

func NewNotificationManager(keyListMapper *repository.KeyListMapper, keySetMapper *repository.KeySetMapper) *NotificationManager {
	return &NotificationManager{
		keyListMapper: keyListMapper,
		keySetMapper:  keySetMapper,
	}
}

func (notificationManager *NotificationManager) Subcribe(meetingID string) <-chan *redis.Message {
	logger.Debug("Subcribe on %v", meetingID)
	channelMessages := notificationManager.keyListMapper.Subscribe(meetingID)
	return channelMessages
}
