package config

var Prod = NewConfig()

type Config struct {
	Server         string `json:"server"`
	ServerResponse string `json:"server_response"`
}

func NewConfig() *Config {
	return &Config{
		Server:         "localhost:8888",
		ServerResponse: "http://localhost:8000",
	}
}
