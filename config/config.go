package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Src             string `env:"MANYMOTD_SRC,default=$HOME/.manymotd"`
	RefreshInterval string `env:"MANYMOTD_REFRESH_INTERVAL"`
}

func GetConfig() Config {
	ctx := context.Background()

	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		panic(err)
	}
	log.Printf("Loaded configuration: %#v", c)
	return c
}
