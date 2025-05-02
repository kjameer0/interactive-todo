package todo

import (
	"encoding/json"
	"errors"
	"os"
)

//load config

// config struct
type Config struct {
	ShowComplete bool   `json:"show-complete"`
	SavePath     string `json:"save-path"`
}

func (a *App) SaveConfig() error {
	data, err := json.Marshal(a.Config)
	if err != nil {
		return err
	}
	os.WriteFile(a.configPath, []byte(data), 0644)
	return nil
}
func (a *App) LoadConfig() (*Config, error) {
	if a.configPath == "" {
		return &Config{}, errors.New("No config path provided")
	}
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
