package flags

import (
	"flag"

	"github.com/KznRkjp/go-url-shrtnr-v2/internal/config"
)

func ParseFlags(cfg *config.Config) {
	flag.StringVar(&cfg.Server, "a", cfg.Server, "Server address to listen on")
	flag.StringVar(&cfg.ServerResponse, "b", cfg.ServerResponse, "Server response URL for shortened links")
	flag.Parse()
}
