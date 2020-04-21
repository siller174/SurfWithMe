package meeting

import (
	"github.com/gorilla/context"
	"github.com/siller174/meetingHelper/pkg/common"
	"github.com/siller174/meetingHelper/pkg/logger"
	"github.com/siller174/meetingHelper/pkg/meetingHelper"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/structs"
	"github.com/siller174/meetingHelper/pkg/utils/converter"
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

type Meeting struct {
	service      *common.MeetingService
	errorHandler *handler.Handler
}

func NewMeetingApi(service *common.MeetingService, errorHandler *handler.Handler) *Meeting {
	return &Meeting{
		service:      service,
		errorHandler: errorHandler,
	}
}

func (meeting *Meeting) Create() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		mtg, err := meeting.service.Create()
		if err != nil {
			meeting.errorHandler.Handle(w, err)
			return
		}
		logger.Debug("Create %v+", mtg)
		writeMeetingResponse(w, r, *mtg)
	}
}

func (meeting *Meeting) Get() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		mtg := getMeeting(r)
		mtg, err := meeting.service.Get(mtg)
		if err != nil {
			meeting.errorHandler.Handle(w, err)
			return
		}
		logger.Debug("Get %v+", mtg)
		writeMeetingResponse(w, r, *mtg)
	}
}

func (meeting *Meeting) Put() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		mtg := getMeeting(r)
		mtg.SetTime()
		err := meeting.service.Put(mtg)
		if err != nil {
			meeting.errorHandler.Handle(w, err)
			return
		}
		logger.Debug("Put %v+", mtg)
		response.Empty(w)
	}
}

func (meeting *Meeting) IsMember() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		meeting := getMeeting(r)
		if meeting != nil {
			response.Empty(w)
		}
	}
}

func (meeting *Meeting) History() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		mtg := getMeeting(r)
		history, err := meeting.service.History(mtg)
		if err != nil {
			meeting.errorHandler.Handle(w, err)
			return
		}
		meetingJSON, err := converter.StructToJsonByte(history)
		response.WriteJSON(w, http.StatusOK, meetingJSON)
	}
}

func (meeting *Meeting) Delete() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		mtg := getMeeting(r)
		err := meeting.service.Delete(mtg)
		if err != nil {
			meeting.errorHandler.Handle(w, err)
			return
		}
		logger.Debug("Delete %v+", mtg)
		response.Empty(w)
	}
}

func getMeeting(r *http.Request) *structs.Meeting {
	if rv := context.Get(r, meetingHelper.UnitContext); rv != nil {
		return rv.(*structs.Meeting)
	}
	return nil
}

func writeMeetingResponse(w http.ResponseWriter, r *http.Request, meeting structs.Meeting) {
	meetingJSON, err := converter.StructToJsonByte(meeting)
	if err != nil {
		logger.Error("Could not write meeting %v to response", meeting)
	}
	err = response.WriteJSON(w, http.StatusOK, meetingJSON)
	if err != nil {
		logger.Error("Could not write meeting %v to response", meeting)
	}
}
