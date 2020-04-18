package response

import (
	"net/http"

	"github.com/siller174/meetingHelper/pkg/logger"
)

func WriteJSON(w http.ResponseWriter, code int, message []byte) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, HEAD")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0") // TODO CHECK
	w.Header().Set("Vary", "Accept-Encoding") // TODO CHECK
	w.WriteHeader(code)
	writeBytes, err := w.Write(message)
	if err != nil {
		logger.Error("Could not write json message in response")
		return err
	}
	if writeBytes == 0 {
		logger.Error("Wrote empty message in response")
	}

	return nil
}

func Empty(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}
