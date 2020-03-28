package config

import (
	"github.com/spf13/viper"
)

func Read(configPath string, rawVal interface{}) error {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	err := viper.Unmarshal(rawVal)
	if err != nil {
		return err
	}
	return nil
}
