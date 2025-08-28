package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	empty := Config{}

	homePath, err := os.UserHomeDir()
	if err != nil {
		return empty, fmt.Errorf("UserHomeDir not found")
	}

	fullPath := homePath + "/.gatorconfig.json"

	contentBytes, err := os.ReadFile(fullPath)
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
