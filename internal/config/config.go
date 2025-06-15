package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
	"runtime"
)

type (
	Config struct {
		AppName      string   `env:"APP_NAME"`
		AppPort      string   `env:"APP_PORT"`
		JWTSecretKey string   `env:"JWT_SECRET_KEY"`
		Database     Database `envPrefix:"DB_"`
	}

	Database struct {
		Host     string `env:"HOST"`
		User     string `env:"USER"`
		Password string `env:"PASSWORD"`
		Name     string `env:"NAME"`
		Port     string `env:"PORT"`
	}
)

func LoadConfig() *Config {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../../")

	// Load file .env dari root folder
	err := godotenv.Load(filepath.Join(basePath, ".env"))
	if err != nil {
		log.Println("No .env file found in root, continuing with system env")
	}

	var cfg Config
	if err = env.Parse(&cfg); err != nil {
		return nil
	}
	return &cfg
}
