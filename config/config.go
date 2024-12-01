package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DB struct {
		Driver   string
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode config: %s", err)
	}
}
