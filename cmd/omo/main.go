package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"gitlab.com/oramaz/omo/internal/app"
	"gitlab.com/oramaz/omo/internal/app/config"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := config.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Start(config); err != nil {
		log.Fatal(err)
	}
}
