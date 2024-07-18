package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Redis    RedisConfig
}

type DatabaseConfig struct {
	DSN string
}

type ServerConfig struct {
	Address string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local" // default to local if no environment is set
	}

	viper.SetConfigName("config." + env)
	viper.AddConfigPath("config/environments")
	viper.SetConfigType("ini")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading config file:", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Error unmarshaling config:", err)
	}
	return &cfg
}
