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

func getConfigPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("UserHomeDir not found")
	}

	fullPath := homePath + configFile

	return fullPath, nil
}

func Read() (Config, error) {
	empty := Config{}

	path, err := getConfigPath()
	if err != nil {
		return empty, fmt.Errorf("error getting path: %v", err)
	}

	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return empty, fmt.Errorf("error reading file: %v", err)
	}

	var parsed Config
	err = json.Unmarshal(contentBytes, &parsed)
	if err != nil {
		return empty, fmt.Errorf("could not parse json")
	}

	return parsed, nil
}
