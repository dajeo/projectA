package config

import "github.com/spf13/viper"

var config *viper.Viper

func Init(configName string) {
	config = viper.New()
	config.SetConfigType("toml")
	config.SetConfigName(configName)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err := config.ReadInConfig()
	if err != nil {
		return
	}
}

func GetConfig() *viper.Viper {
	return config
}
