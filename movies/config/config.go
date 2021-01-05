package config

import (
	"fmt"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

//LaunchMode - создание нового типа для простаты восприятия
type LaunchMode string

// Создание констант
const (
	LocalEnv LaunchMode = "local"
	ProdEnv LaunchMode = "prod"
)

//Config - структура конфигурационными переменными
type Config struct {
	Port string `envconfig:"PORT" default:"8080"`
}

//Load - загружает переменные окружения из файлов
func Load(launchMode LaunchMode, path string) (*Config, error) {
	switch launchMode {
	case LocalEnv:
		cfgPath := filepath.Join(path, fmt.Sprintf("%s.env", launchMode))

		err := godotenv.Load(cfgPath)
		if err != nil {
			return nil, fmt.Errorf("load .env config file: %w", err)
		}
	case ProdEnv:
		// любые переменные которые предоставляются окружением
	default:
		return nil, fmt.Errorf("unexpected LAUNCH_MODE: [%s]", launchMode)
	}

	config := new(Config)
	err := envconfig.Process("", config)
	if err != nil {
		return nil, fmt.Errorf("get config from env: %w", err)
	}

	return config, nil
}