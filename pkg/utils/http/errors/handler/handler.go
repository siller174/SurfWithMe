package handler

import (
	"github.com/siller174/meetingHelper/pkg/logger"
	"github.com/siller174/meetingHelper/pkg/utils/http/errors"
	"github.com/siller174/meetingHelper/pkg/utils/http/response"
	"net/http"
)

type Handler struct {
	devMode bool
}

func NewHandler(devMode bool) *Handler {
	return &Handler{
		devMode:devMode,
	}
}

func (handler *Handler) Handle(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	var innerErr error
	var status int
	var res string

	switch typeError := err.(type) {
	case errors.HTTPError:
		status = typeError.GetStatus()
		res = typeError.ToResponse()
	default:
		defErr := errors.NewInternalErr(err)
		res = defErr.ToResponse()
		status = defErr.Status
	}
	logger.Error(err.Error())

	if handler.devMode {
		innerErr = response.WriteJSON(w, status, []byte(res))

		if innerErr != nil {
			logger.Error(err.Error())
		}
	}
}
