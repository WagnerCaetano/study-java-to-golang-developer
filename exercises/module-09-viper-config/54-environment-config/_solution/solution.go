package envconfig

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port    int
	AppName string
}

func LoadFromEnv() (Config, error) {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		return Config{}, fmt.Errorf("PORT is required")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return Config{}, fmt.Errorf("invalid PORT: %w", err)
	}
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		return Config{}, fmt.Errorf("APP_NAME is required")
	}
	return Config{Port: port, AppName: appName}, nil
}
