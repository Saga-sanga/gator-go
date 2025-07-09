package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	dat, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("Error reading file: %s", err)
	}

	var config Config
	if err := json.Unmarshal(dat, &config); err != nil {
		return Config{}, fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error reading Home dir: %v", err)
	}
	path := filepath.Join(homeDir, configFileName)

	return path, nil
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Error marshalling json: %v", err)
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}

	return nil
}

func (c *Config) SetUser(user_name string) error {
	c.CurrentUserName = user_name
	return write(*c)
}
