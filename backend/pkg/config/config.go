package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DbPassword    string `mapstructure:"MONGO_INITDB_ROOT_PASSWORD"`
	DbUsername    string `mapstructure:"MONGO_INITDB_ROOT_USERNAME"`
	Frontend      string `mapstructure:"FRONTEND"`
	MongoDatabase string `mapstructure:"MONGO_DATABASE"`
	MongoUri      string `mapstructure:"MONGO_URI"`
	Port          string `mapstructure:"PORT"`
}

func LoadConfig(path string) Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("go")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.AutomaticEnv()

	config := Config{}
	err = viper.Unmarshal(&config)

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return config
}

var AppConfig = LoadConfig(".")
