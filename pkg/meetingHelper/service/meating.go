package service

import (
	"github.com/rs/xid"
	"github.com/siller174/meetingHelper/pkg/logger"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/repository"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/structs"
)

type MeetingService struct {
	repository repository.RedisRepo
}

func NewMeetingService(repos repository.RedisRepo) *MeetingService {
	return &MeetingService{
		repository: repos,
	}
}

func (ms *MeetingService) Create() structs.Meeting {
	meeting := structs.Meeting{
		ID: generateID(),
	}
	logger.Debug("Create %+v", meeting)
	return meeting
}

func (ms *MeetingService) Put(meeting structs.Meeting) {
	//todo put to redis
	logger.Debug("Put %+v", meeting)
}

func (ms *MeetingService) Get(meeting structs.Meeting) structs.Meeting {
	//todo get from redis last url and time
	logger.Debug("Get %+v", meeting)
	return meeting
}

func (ms *MeetingService) Delete(meeting structs.Meeting) {
	//todo remove meeting from redis
	logger.Debug("Delete %+v", meeting)

}

// func (ms *MeetingService) History(meeting structs.Meeting) []structs.Meeting { TODO
func (ms *MeetingService) History(meeting structs.Meeting) {
	//todo get all urls from redis
	logger.Debug("Gey history %+v", meeting)
}

func generateID() string {
	guid := xid.New()
	return string(guid.String()[0:5])
}
