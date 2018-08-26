package main

import (
	"encoding/json"
	"os"
	"path"
	"sync"
)

var once sync.Once
var instance *Config

// GetConfig get config
func GetConfig() *Config {
	once.Do(func() {
		var config *Config
		pwd, _ := os.Getwd()
		path := path.Join(pwd, "config.json")
		configFile, err := os.Open(path)
		defer configFile.Close()

		if err != nil {
			return
		}

		jsonParser := json.NewDecoder(configFile)
		err = jsonParser.Decode(&config)
		if err != nil {
			return
		}

		instance = config
	})

	return instance
}

// Config config entry
type Config struct {
	APIKey    string `json:"apikey"`
	RateLimit int    `json:"ratelimit"`
}
