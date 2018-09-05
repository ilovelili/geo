// Package config configuration related
package config

import (
	"encoding/json"
	"os"
	"path"
	"sync"
)

var instance *Config
var once sync.Once

// GetConfig get config defined in config.json
func GetConfig() (config *Config) {
	once.Do(func() {
		var config *Config
		pwd, _ := os.Getwd()
		path := path.Join(pwd, "config.json")
		configFile, err := os.Open(path)
		defer configFile.Close()
		if err != nil {
			panic(err)
		}

		jsonParser := json.NewDecoder(configFile)
		if err = jsonParser.Decode(&config); err != nil {
			return
		}

		instance = config
	})

	return instance
}

// Config config entry
type Config struct {
	APIKey    string `json:"apikey"`
	RateLimit int    `json:"rate_limit"`
}

// GetRateLimit get rate limit
func (c *Config) GetRateLimit() int {
	if c.RateLimit == 0 {
		return 3
	}

	return c.RateLimit
}
