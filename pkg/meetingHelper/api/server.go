package api

import (
	"github.com/siller174/meetingHelper/pkg/common/repository"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/service"
	"net/http"

	"github.com/siller174/meetingHelper/pkg/meetingHelper/api/meeting"
	"github.com/siller174/meetingHelper/pkg/utils/http/errors/handler"

	"github.com/gorilla/mux"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/api/manage"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/config"
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
	redis := repository.New(config.Redis)
	keyList := repository.NewKeyListMapper(redis)
	keySet := repository.NewKeySetMapper(redis)
	meetingService := service.NewMeetingService(keyList, keySet)
	errorHandler := handler.NewHandler(config.DevMode)
	middleware := newMeetingMiddleWare(meetingService, errorHandler)
	meetingApi := meeting.NewMeetingApi(meetingService, errorHandler)
	healthApi := manage.NewHealthApi(redis)

	router := mux.NewRouter()
	router.HandleFunc(manage.HealthRoute, healthApi.Handle()).Methods(http.MethodGet)
	router.HandleFunc(meeting.RouteCreate, meetingApi.Create()).Methods(http.MethodPost)
	router.HandleFunc(meeting.RouteGet, meetingApi.Get()).Methods(http.MethodGet)
	router.HandleFunc(meeting.RoutePut, meetingApi.Put()).Methods(http.MethodPut)
	router.HandleFunc(meeting.RouteDelete, meetingApi.Delete()).Methods(http.MethodDelete)
	router.HandleFunc(meeting.RouteHistory, meetingApi.History()).Methods(http.MethodGet)
	router.HandleFunc(meeting.RouteOptions, meetingApi.IsMember()).Methods(http.MethodOptions)
	router.Use(AddCors,middleware.MiddleWare)
	return router
}
