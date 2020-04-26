package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/context"
	"github.com/siller174/meetingHelper/pkg/meetingHelper"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/api/meeting"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/service"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/structs"
	"github.com/siller174/meetingHelper/pkg/utils/http/errors"
	"github.com/siller174/meetingHelper/pkg/utils/http/errors/handler"
	"github.com/siller174/meetingHelper/pkg/utils/http/response"

	"github.com/siller174/meetingHelper/pkg/logger"
)

type MeetingMiddleWare struct {
	service      *service.MeetingService
	errorHandler *handler.Handler
}

func newMeetingMiddleWare(service *service.MeetingService, errorHandler *handler.Handler) *MeetingMiddleWare {
	return &MeetingMiddleWare{
		service:      service,
		errorHandler: errorHandler,
	}
}

func (meetingMiddleWare *MeetingMiddleWare) MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri := r.RequestURI
		method := r.Method
		logger.Debug("Start request %v %v", method, uri)
		defer logger.Debug("Finish request %v %v", method, uri)
		if uri != meeting.RouteCreate {
			mtg, err := decodeMeeting(r)
			if err != nil {
				meetingMiddleWare.errorHandler.Handle(w, errors.NewBadRequest(err))
				return
			}
			isMem, err := isMember(meetingMiddleWare.service, mtg)
			if err != nil {
				meetingMiddleWare.errorHandler.Handle(w, err)
				return
			}
			if isMem {
				context.Set(r, meetingHelper.UnitContext, mtg)
			} else {
				response.NotFound(w)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func decodeMeeting(r *http.Request) (*structs.Meeting, error) {
	var mtg structs.Meeting
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn("Could not get body from request")
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &mtg)
	if err != nil {
		logger.Warn("Could not decode meeting. Body %v", string(bodyBytes))
		return nil, err
	}
	return &mtg, nil
}

func isMember(service *service.MeetingService, meeting *structs.Meeting) (bool, error) {
	isMember, err := service.IsMember(meeting)
	if err != nil {
		return false, err
	}
	return isMember, nil
}
