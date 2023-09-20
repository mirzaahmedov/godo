package config

import (
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	DatabaseURL    string `toml:"database_url"`
	BindingAddress string `toml:"binding_address"`
}

func Load(filePath string) (*Config, error) {
	c := &Config{}

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(bytes, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
