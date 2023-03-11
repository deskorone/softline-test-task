package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

// Config - конфиг приложения
type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Token    TokenConfig    `yaml:"token"`
	Server   Server         `yaml:"server"`
}

// DatabaseConfig - конфиг базы дынных
type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"db_name"`
}

// TokenConfig - конфигурация для токена авторизации
type TokenConfig struct {
	SecretWord string        `yaml:"secret_word"`
	Expired    time.Duration `yaml:"expired"`
}

// Server - конфиг сервера
type Server struct {
	Port int `yaml:"port"`
}

// GetConfig - функция считывающая конфиг из файла конфигурации
func GetConfig() Config {
	f, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d := yaml.NewDecoder(f)

	config := Config{}
	if err := d.Decode(&config); err != nil {
		log.Fatal(err)
	}

	return config
}
