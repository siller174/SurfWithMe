package manage

import (
	"github.com/siller174/meetingHelper/pkg/utils/http/response"
	"net/http"
)

const HealthRoute = "/health"

func HealthApi(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	_ = response.WriteJSON(w, http.StatusOK, []byte(`{"status": "UP"}`))
}
