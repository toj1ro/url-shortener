package config

import (
	"fmt"
	"os"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresConfig() PostgresConfig {
	var c PostgresConfig
	c.Host = os.Getenv("POSTGRES_HOST")
	c.Port = os.Getenv("POSTGRES_PORT")
	c.User = os.Getenv("POSTGRES_USER")
	c.Password = os.Getenv("POSTGRES_PASSWORD")
	c.DBName = os.Getenv("POSTGRES_DBNAME")
	c.SSLMode = os.Getenv("POSTGRES_SSLMODE")
	return c
}

func (c PostgresConfig) ConnString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}
