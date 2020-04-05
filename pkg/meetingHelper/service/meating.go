package service

import (
	"encoding/json"
	"github.com/rs/xid"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/repository"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/structs"
	"github.com/siller174/meetingHelper/pkg/utils/converter"
)

type MeetingService struct {
	repository repository.KeyListMapper
}

func NewMeetingService(repos repository.KeyListMapper) *MeetingService {
	return &MeetingService{
		repository: repos,
	}
}

func (ms *MeetingService) Create() *structs.Meeting {
	meeting := structs.Meeting{
		ID: generateID(),
	}
	return &meeting

}

func (ms *MeetingService) Put(meeting *structs.Meeting) error {
	meetingJson, err := converter.StructToJsonString(meeting)
	if err != nil {
		return err
	}
	err = ms.repository.Put(meeting.ID, meetingJson)
	return err
}

func (ms *MeetingService) Get(meeting *structs.Meeting) (*structs.Meeting, error) {
	resultMeeting := structs.Meeting{}
	meetingJson, err := ms.repository.GetLast(meeting.ID)
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
	history, err := ms.repository.GetAll(meeting.ID)
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


func (ms *MeetingService) Delete(meeting *structs.Meeting) {
	//todo remove meeting from redis
	//ms.repository.Delete(meeting)

}


func generateID() string {
	guid := xid.New()
	return string(guid.String()[0:5])
}
