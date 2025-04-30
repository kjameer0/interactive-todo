package main

import (
	"encoding/json"
	"os"
)

//load config

// config struct
type Config struct {
	ShowComplete bool `json:"show-complete"`
}

func (a *app) saveConfig() error {
	data, err := json.Marshal(a.config)
	if err != nil {
		return err
	}
	os.WriteFile(a.configPath, []byte(data), 0644)
	return nil
}
func (a *app) loadConfig() (*Config, error) {
	data, err := os.ReadFile(a.configPath)
	if err != nil {
		return &Config{}, err
	}
	var c Config
	err = json.Unmarshal(data, &c)
	if err != nil {
		return &Config{}, err
	}
	return &c, nil
}
