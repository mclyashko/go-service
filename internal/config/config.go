package config

import "os"

type Config struct {
	Addr string
	Name string
}

func Load() Config {
	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	name := os.Getenv("APP_NAME")
	if name == "" {
		name = "go-service"
	}
	return Config{
		Addr: addr,
		Name: name,
	}
}
