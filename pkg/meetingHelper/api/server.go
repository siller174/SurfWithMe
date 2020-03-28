package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/api/manage"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/api/meeting"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/config"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/repository"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/service"
)

func New(appConfig config.App) *http.Server {
	router := initRouters(appConfig)

	srv := http.Server{
		Handler:      router,
		Addr:         appConfig.Server.Port,
		WriteTimeout: appConfig.Server.WriteTimeout,
		ReadTimeout:  appConfig.Server.ReadTimeout,
	}
	return &srv
}

func initRouters(config config.App) *mux.Router {
	meetingService := service.NewMeetingService(repository.New(config.Redis))
	router := mux.NewRouter()
	router.HandleFunc(manage.HealthRoute, manage.HealthApi).Methods(http.MethodGet)
	router.HandleFunc(meeting.RouteCreate, meeting.Create(meetingService)).Methods(http.MethodPost)
	router.HandleFunc(meeting.RouteGet, meeting.Get(meetingService)).Methods(http.MethodGet)
	router.HandleFunc(meeting.RoutePut, meeting.Put(meetingService)).Methods(http.MethodPut)
	router.HandleFunc(meeting.RouteDelete, meeting.Delete(meetingService)).Methods(http.MethodDelete)
	router.HandleFunc(meeting.RouteHistory, meeting.History(meetingService)).Methods(http.MethodGet)
	router.Use(MiddleWare)
	return router
}
