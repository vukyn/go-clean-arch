package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Auth       AuthConfig
	PostgreSQL PostgreSQLConfig
}
type PostgreSQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type AuthConfig struct {
	JWTSecret string
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config file not found")
		}
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic("Unable to unmarshal config")
	}
	return &c
}
