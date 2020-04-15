package manage

import (
	"github.com/siller174/meetingHelper/pkg/logger"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/structs"
	"github.com/siller174/meetingHelper/pkg/utils/http/response"
	"net/http"
)

const HealthRoute = "/health"

func HealthApi(health structs.Health) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		_, err := health.RedisClient.Ping().Result()
		if err == nil {
			_ = response.WriteJSON(w, http.StatusOK, []byte(`{"status": "UP"}`))
		} else {
			logger.Error("Redis in not available")
			_ = response.WriteJSON(w, http.StatusOK, []byte(`{"status": "DOWN"}`))
		}
	}
}
