package api

import (
	"net/http"

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

// func initRouters(config config.App) *mux.Router {
// 	redis := repository.New(config.Redis)
// 	keyList := repository.NewKeyListMapper(redis)
// 	keySet := repository.NewKeySetMapper(redis)

// }
