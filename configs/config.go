package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Config struct {
	MySQLDb DBConfig
}

func GetConfig() *Config {
	err := godotenv.Load()
	fmt.Println(os.Getenv("DB_HOST"))

	return &Config{
		MySQLDb: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "user"),
			Password: getEnv("DB_PASSWORD", "pass"),
			DBName:   getEnv("DB_NAME", "library"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
