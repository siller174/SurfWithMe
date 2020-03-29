package structs

import "time"

type Meeting struct {
	ID   string `json:"id"`
	Url  string `json:"url"`
	Time string `json:"time"`
}

func (meeting *Meeting) SetTime()  {
	meeting.Time = time.Now().Format("2 Jan 2006 15:04:05")
}