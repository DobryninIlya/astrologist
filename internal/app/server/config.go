package server

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	DatabaseURL string `toml:"database_url"`
	Token       string `toml:"token"`
	Debug       bool   `toml:"debug"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
