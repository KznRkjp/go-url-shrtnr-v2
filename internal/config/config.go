package config

var Prod = NewConfig()

type Config struct {
	Server         string `json:"server" env:"SERVER_ADDRESS"`
	ServerResponse string `json:"server_response" env:"BASE_URL"`
}

func NewConfig() *Config {
	return &Config{
		Server:         "localhost:8080",
		ServerResponse: "http://localhost:8080",
	}
}
