package main

import (
	"flag"
	"log"

	"github.com/mirzaahmedov/godo/internal/config"
	"github.com/mirzaahmedov/godo/internal/store/postgres"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "configs/local.toml", "path to the config file")
}
func main() {
	flag.Parse()

	c, err := config.Load(configPath)
	if err != nil {
		log.Fatal("could not load the config file: \n", err)
	}

	store := postgres.NewStore(c.DatabaseURL)
	err = store.Open()
	if err != nil {
		log.Fatal("Could not connect to the database: \n", err)
	}
	defer store.Close()

	log.Println(c.BindingAddress, c.DatabaseURL)
}
