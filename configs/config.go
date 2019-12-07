package configs

import (
	"github.com/spf13/viper"
)

const configFilePath = "../configs"
const configFileName = "config"

type SocketConfig struct {
	Host string
	Port int
}

func LoadConfigs() error {
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(configFilePath)
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
