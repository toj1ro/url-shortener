package config

import (
	"flag"
	"os"
)

type Config struct {
	ServerAddress string
	BaseURL       string
}

func NewConfig() Config {
	var c Config
	c.ServerAddress = *flag.String("a", os.Getenv("SERVER_ADDRESS"), "Адрес запуска HTTP-сервера")
	c.BaseURL = *flag.String("b", os.Getenv("BASE_URL"), "Базовый адрес")
	return c
}
