package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DB_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func Read() Config {
	gator_config_path, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error getting config file path:", err)
		return Config{}
	}

	data, err := os.ReadFile(gator_config_path)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		fmt.Println("exiting with code 1 at config reading")
		os.Exit(1)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		fmt.Println("exiting with code 1 at config parsing")
		os.Exit(1)
	}

	return config
}

func (cfg Config) SetUser(user_name string) error {
	cfg.Current_user_name = user_name
	fmt.Printf("Logging in: %s\n", user_name)
	return write(cfg)
}

func getConfigFilePath() (string, error) {
	home_path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error getting user home directory: %v", err)
	}

	return fmt.Sprintf("%s/%s", home_path, configFileName), nil
}

func write(cfg Config) error {
	gator_config_path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Error getting config file path: %v", err)
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Error marshaling config data: %v", err)
	}

	err = os.WriteFile(gator_config_path, data, 0644)
	if err != nil {
		return fmt.Errorf("Error writing config file: %v", err)
	}

	return nil
}
