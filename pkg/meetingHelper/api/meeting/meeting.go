package meeting

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/siller174/meetingHelper/pkg/logger"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/service"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/structs"
	"github.com/siller174/meetingHelper/pkg/utils/converter"
	"github.com/siller174/meetingHelper/pkg/utils/http/errors"
	"github.com/siller174/meetingHelper/pkg/utils/http/errors/handler"
	"github.com/siller174/meetingHelper/pkg/utils/http/response"
	"net/http"
)

const api = "/api/v1/meeting"
const RouteCreate = api + "/create"
const RouteGet = api
const RoutePut = api
const RouteHistory = api + "/history"
const RouteDelete = api
const RouteOptions = api

func Create(service *service.MeetingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		meeting, err := service.Create()
		if err != nil || meeting == nil {
			handler.Handle(w, err)
			return
		}
		logger.Debug("Create %v+", meeting)
		writeMeetingResponse(w, *meeting)
	}
}

func Get(service *service.MeetingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		meeting := decodeMeeting(w, r)
		if meeting == nil {
			return
		}

		isMember := isMember(w, service, meeting)
		if !isMember {
			return
		}

		result, err := service.Get(meeting)
		if err != nil {
			handler.Handle(w, err)
			return
		}
		logger.Debug("Get %v+", meeting)
		writeMeetingResponse(w, *result)
	}
}

func Put(service *service.MeetingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		meeting := decodeMeeting(w, r)
		if meeting == nil {
			return
		}

		isMember := isMember(w, service, meeting)
		if !isMember {
			return
		}

		meeting.SetTime()
		err := service.Put(meeting)
		if err != nil {
			handler.Handle(w, err)
			return
		}
		logger.Debug("Put %v+", meeting)
		response.Empty(w)
	}
}

func IsMember(service *service.MeetingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		meeting := decodeMeeting(w, r)
		if meeting == nil {
			return
		}
		isMember := isMember(w, service, meeting)
		if isMember {
			response.Empty(w)
		}
	}
}

func History(service *service.MeetingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		meeting := decodeMeeting(w, r)
		if meeting == nil {
			return
		}
		history, err := service.History(meeting)
		if err != nil {
			handler.Handle(w, err)
			return
		}
		meetingJSON, err := converter.StructToJsonByte(history)
		err = response.WriteJSON(w, http.StatusOK, meetingJSON)
	}
}

func Delete(service *service.MeetingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		meeting := decodeMeeting(w, r)
		if meeting == nil {
			return
		}

		isMember := isMember(w, service, meeting)
		if !isMember {
			return
		}

		err := service.Delete(meeting)
		if err != nil {
			handler.Handle(w, err)
			return
		}
		logger.Debug("Delete %v+", meeting)
		response.Empty(w)
	}
}

func isMember(w http.ResponseWriter, service *service.MeetingService, meeting *structs.Meeting) bool {
	isMember, err := service.IsMember(meeting)
	if err != nil {
		handler.Handle(w, err)
		return false
	}
	if !isMember {
		response.NotFound(w)
		return false
	}
	return true
}

func writeMeetingResponse(w http.ResponseWriter, meeting structs.Meeting) {
	meetingJSON, err := converter.StructToJsonByte(meeting)
	if err != nil {
		handler.Handle(w, err)
	}
	err = response.WriteJSON(w, http.StatusOK, meetingJSON)
}

func decodeMeeting(w http.ResponseWriter, r *http.Request) *structs.Meeting {
	var meeting structs.Meeting
	err := json.NewDecoder(r.Body).Decode(&meeting)
	if err != nil {
		handler.Handle(w, errors.NewBadRequest(err))
		return nil
	}
	return &meeting
}

func getMeetingID(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["meetingID"]
}
