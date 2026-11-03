package envconfig

import "errors"

// Config holds application settings from the environment.
type Config struct {
	Port    int
	AppName string
}

var errInvalidConfig = errors.New("invalid config")

// LoadFromEnv reads PORT and APP_NAME from the environment.
func LoadFromEnv() (Config, error) {
	// FIX ME: read PORT (int) and APP_NAME (non-empty string)
	return Config{}, errInvalidConfig
}
