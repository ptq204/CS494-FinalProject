package client

import (
	"github.com/spf13/viper"
)

const configFilePath = "../server/db/config"
const configFileName = "config"

type DatabaseConfig struct {
	Hostname string
	Port     int
	DBName   string
	Username string
	Password string
}

func LoadDatabaseConfig() error {
	viper.Reset()
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(configFilePath)
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
