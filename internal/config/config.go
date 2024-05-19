package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port uint16 `env:"PORT"`
}

func LoadConfig() *Config {
	viper.AutomaticEnv()

	return &Config{
		Port: viper.GetUint16("PORT"),
	}
}
