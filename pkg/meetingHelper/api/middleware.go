package api

import (
	"net/http"

	"github.com/siller174/meetingHelper/pkg/logger"
)

func MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri := r.RequestURI
		method := r.Method
		logger.Debug("Start request %v %v", method, uri)
		next.ServeHTTP(w, r)
		logger.Debug("Finish request %v %v", method, uri)
	})
}
