package notification

import (
	"github.com/siller174/meetingHelper/pkg/logger"
	"github.com/siller174/meetingHelper/pkg/meetingHelper/config"
	"github.com/siller174/meetingHelper/pkg/notification/api"
	"github.com/spf13/pflag"
)

func main() {
	configPath := pflag.String("config-path", "./notification.properties", "Path to config file")
	pflag.Parse()
	appConfig := config.New(*configPath)
	server := api.New(appConfig)
	logger.Fatal("App was close", server.ListenAndServe())
}
