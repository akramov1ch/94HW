package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	viper.AutomaticEnv()
}
