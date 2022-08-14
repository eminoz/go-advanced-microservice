package config

import (
	"github.com/joho/godotenv"
	"os"
)

var Config *Configuration

type Configuration struct {
	MongoDb   string `mapstructure:"MONGODB_URI"`
	Port      string `mapstructure:"PORT"`
	AppSecret string `mapstructure:"APP_SECRET"`
}

func SetupConfig() (err error) {
	config := godotenv.Load("./config/.env")
	if config != nil {
		return nil
	}
	configuration := &Configuration{
		MongoDb:   os.Getenv("MONGODB_URI"),
		Port:      os.Getenv("PORT"),
		AppSecret: os.Getenv("APP_SECRET"),
	}

	Config = configuration
	return
}

func GetConfig() *Configuration {
	return Config
}
