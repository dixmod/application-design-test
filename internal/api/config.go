package api

type Config struct {
	Port string `json:"port"`
}

func NewConfig() *Config {
	return &Config{
		Port: "8080",
	}
}
