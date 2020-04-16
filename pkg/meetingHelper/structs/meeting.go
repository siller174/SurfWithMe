package structs

import (
	"encoding/json"
	"github.com/siller174/meetingHelper/pkg/utils/http/errors"
	"time"
)

type Meeting struct {
	ID   string `json:"id"`
	Url  string `json:"url"`
	Time string `json:"time"`
}

func (meeting *Meeting) SetTime() {
	meeting.Time = time.Now().Format("2 Jan 2006 15:04:05")
}

func NewMeetingFromJSON(meetingJSON string) (*Meeting, error) {
	tempMeeting := &Meeting{}

	err := json.Unmarshal([]byte(meetingJSON), tempMeeting)
	if err != nil {
		return nil,  errors.NewBadRequest(err)
	}
	return tempMeeting, nil

}