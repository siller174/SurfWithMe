package meeting

import (
	"encoding/json"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/structs"
	"github.com/siller174/meetingHelper/pkg/utils/http/errors"
	"github.com/siller174/meetingHelper/pkg/utils/http/errors/handler"
	"github.com/siller174/meetingHelper/pkg/utils/http/response"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/service"
)

const api = "/api/v1/meeting"
const RouteCreate = api + "/create"
const RouteGet = api
const RoutePut = api
const RouteHistory = api + "/history"
const RouteDelete = api

func Create(service *service.MeetingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		meeting := service.Create()
		writeMeetingResponse(w, r, meeting)

	}
}

func Get(service *service.MeetingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		meeting := decodeMeeting(w, r)
		if meeting == nil {
			return
		}
		result := service.Get(*meeting)
		writeMeetingResponse(w, r, result)
	}
}

func Put(service *service.MeetingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		meeting := decodeMeeting(w, r)
		if meeting == nil {
			return
		}
		service.Put(*meeting)
		response.Empty(w)
	}
}

func History(service *service.MeetingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func Delete(service *service.MeetingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		meeting := decodeMeeting(w, r)
		if meeting == nil {
			return
		}
		service.Delete(*meeting)
		response.Empty(w)
	}
}

func writeMeetingResponse(w http.ResponseWriter, r *http.Request, meeting structs.Meeting) {
	bytes, err := json.Marshal(meeting)
	if err != nil {
		handler.Handle(w, errors.NewInternalErr(err))
	}
	err = response.WriteJSON(w, http.StatusOK, bytes)
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
