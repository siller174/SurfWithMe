package service

import (
	"encoding/json"
	"fmt"
	"github.com/rs/xid"
	"github.com/siller174/meetingHelper/pkg/logger"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/repository"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/structs"
	"github.com/siller174/meetingHelper/pkg/utils/converter"
)

const activeSession = "active_session"
const disabeSession = "disable_session"


type MeetingService struct {
	keyList *repository.KeyListMapper
	keySet *repository.KeySetMapper
}

func NewMeetingService(keyList *repository.KeyListMapper, keySet *repository.KeySetMapper) *MeetingService {
	return &MeetingService{
		keyList: keyList,
		keySet: keySet,
	}
}

func (ms *MeetingService) Create() (*structs.Meeting, error) {
	ID := generateID()
	meeting := structs.Meeting{
		ID: ID,
	}
	res, err := ms.keySet.Add(activeSession, ID)
	if err != nil {
		logger.Error("Could not save to %v in Redis. Err: %v",activeSession, err)
		return  nil, err
	}
	if !res {
		return nil, fmt.Errorf("Could not save to %v in Redis", meeting, activeSession)
	}
	return &meeting, nil
}

func (ms *MeetingService) IsMember(meeting *structs.Meeting) (bool, error) {
	isMember, err := ms.keySet.IsMember(activeSession, meeting.ID)
	if err != nil {
		return false, err
	}
	return isMember, err
}

func (ms *MeetingService) Put(meeting *structs.Meeting) error {
	meetingJson, err := converter.StructToJsonString(meeting)
	if err != nil {
		return err
	}
	err = ms.keyList.Put(meeting.ID, meetingJson)
	return err
}

func (ms *MeetingService) Get(meeting *structs.Meeting) (*structs.Meeting, error) {
	resultMeeting := structs.Meeting{}
	meetingJson, err := ms.keyList.GetLast(meeting.ID)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(*meetingJson), &resultMeeting)
	if err != nil {
		return nil, err
	}
	return &resultMeeting, nil
}

func (ms *MeetingService) History(meeting *structs.Meeting) (*[]structs.Meeting, error) {
	history, err := ms.keyList.GetAll(meeting.ID)
	if err != nil {
		return nil, err
	}
	historyMeetings := make([]structs.Meeting, len(*history))

	for _, meetingJson := range *history {
		tempMeeting := structs.Meeting{}
		err = json.Unmarshal([]byte(meetingJson), &tempMeeting)
		if err != nil {
			return nil, err
		}
		historyMeetings = append(historyMeetings, tempMeeting)
	}
	return &historyMeetings, nil
}


func (ms *MeetingService) Delete(meeting *structs.Meeting) error {
	res, err := ms.keySet.Remove(activeSession, meeting.ID)
	if err != nil {
		logger.Error("Could not remove %v in Redis. Err: %v", activeSession, err)
		return err
	}
	if !res {
		return fmt.Errorf("Could not remove %v from %v Redis", meeting, activeSession)
	}
	res, err = ms.keySet.Add(disabeSession, meeting.ID)
	if err != nil {
		logger.Error("Could not save to %v in Redis. Err: %v", disabeSession, err)
		return  err
	}
	if !res {
		return fmt.Errorf("Could not save %v to %v Redis",meeting,  disabeSession)
	}
	return nil

}


func generateID() string {
	guid := xid.New()
	return string(guid.String()[12:])
}
