package manage

import (
	"github.com/go-redis/redis"
	"github.com/siller174/meetingHelper/pkg/logger"
	"github.com/siller174/meetingHelper/pkg/utils/http/response"
	"net/http"
)

const HealthRoute = "/health"

type Health struct {
	redisClient *redis.Client
}

func NewHealthApi (client *redis.Client) *Health {
	return &Health{
		redisClient:client,
	}
}

func (health *Health) Handle() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		_, err := health.redisClient.Ping().Result()
		if err == nil {
			_ = response.WriteJSON(w, http.StatusOK, []byte(`{"status": "UP"}`))
		} else {
			logger.Error("Redis in not available")
			_ = response.WriteJSON(w, http.StatusOK, []byte(`{"status": "DOWN"}`))
		}
	}
}
