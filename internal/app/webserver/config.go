package webserver

import "github.com/yakovasavr/sql_connection/internal/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:	":8000",
		Store:		store.NewConfig(),
	}
}