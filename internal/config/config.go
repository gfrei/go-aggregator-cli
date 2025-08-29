package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	configFile = "/.gatorconfig.json"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func New() Config {
	return Config{}
}

func (c *Config) SetUser(user string) error {
	cfg, err := Read()

	if err != nil {
		return err
	}

	cfg.CurrentUserName = user

	err = writeConfig(cfg)

	if err != nil {
		return err
	}

	return nil
}

func getConfigPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("UserHomeDir not found")
	}

	fullPath := homePath + configFile

	return fullPath, nil
}

func writeConfig(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	path, err := getConfigPath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	empty := Config{}

	path, err := getConfigPath()
	if err != nil {
		return empty, fmt.Errorf("error getting path: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return empty, fmt.Errorf("error reading file: %v", err)
	}

	var parsed Config
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		return empty, fmt.Errorf("could not parse json")
	}

	return parsed, nil
}
