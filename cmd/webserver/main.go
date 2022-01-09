package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/yakovasavr/sql_connection/internal/app/webserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/webserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := webserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := webserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
