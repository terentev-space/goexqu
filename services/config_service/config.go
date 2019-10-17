package config_service

import (
	"encoding/json"
	"goexqu/helpers"
	"io/ioutil"
	"os"
)

type Config struct {
	QueueService struct {
		Type       string `json:"type"`
		Connection string `json:"connection"`
		Name       string `json:"name"`
	} `json:"queue_service"`
}

const ConfigFileName = "config.json"

func LoadConfig() *Config {
	configFile, err := os.Open(ConfigFileName)
	helpers.CheckError(err, "File `"+ConfigFileName+"` not found!")
	defer configFile.Close()
	configData, err := ioutil.ReadAll(configFile)
	helpers.CheckError(err, "Can not read config!")
	config := new(Config)
	err = json.Unmarshal(configData, &config)
	helpers.CheckError(err, "Wrong config!")
	return config
}

func DefaultConfig() *Config {
	config := new(Config)
	config.QueueService.Type = "rabbitmq"
	config.QueueService.Connection = "amqp://guest:guest@localhost:5672/"
	config.QueueService.Name = "goexqu"
	return config
}
