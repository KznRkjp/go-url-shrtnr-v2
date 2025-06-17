package flags

import (
	"flag"
	"fmt"

	"github.com/KznRkjp/go-url-shrtnr-v2/internal/config"
	"github.com/caarlos0/env/v6"
)

func ParseFlags(cfg *config.Config) {
	err := env.Parse(cfg)
	if err != nil {
		fmt.Println("Failed to parse environment variables: " + err.Error())
	}
	flag.StringVar(&cfg.Server, "a", cfg.Server, "Server address to listen on")
	flag.StringVar(&cfg.ServerResponse, "b", cfg.ServerResponse, "Server response URL for shortened links")
	flag.Parse()
}
