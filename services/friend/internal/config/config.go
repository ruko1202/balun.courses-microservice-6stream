package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
)

var (
	Config config
)

type config struct {
	App App
}

type App struct {
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port string `env:"POST" envDefault:"8080"`
}

func (a *App) BuildHostPort() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

func init() {
	err := env.Parse(&Config)
	if err != nil {
		log.Fatalf("env parse failed: %s", err)
	}
}
