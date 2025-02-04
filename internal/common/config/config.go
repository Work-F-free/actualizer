package config

import "os"

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBPassword string
	DBUsername string
	DBSslMode  string
}

type Config struct {
	Port string
	DBConfig
}

var AppConfig *Config

func LoadConfig() {

	dbCfg := DBConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBSslMode:  os.Getenv("DB_SSL_MODE"),
	}

	AppConfig = &Config{
		Port:     os.Getenv("PORT"),
		DBConfig: dbCfg,
	}
}
