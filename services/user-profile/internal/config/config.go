package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/keepalive"
)

var (
	Config config
)

type config struct {
	App App
}

type RestConfig interface {
	BuildHostPort() string
}

type App struct {
	Grpc   Grpc
	Rest   Rest
	Status Status
}

// Rest - contains parameter rest json connection.
type Rest struct {
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port string `env:"POST" envDefault:"8000"`
}

func (a *Rest) BuildHostPort() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

// Status - contains parameter rest json connection.
type Status struct {
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port string `env:"POST" envDefault:"8001"`
}

func (a *Status) BuildHostPort() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

// Grpc - contains parameter address grpc.
type Grpc struct {
	Host             string `env:"HOST" envDefault:"0.0.0.0"`
	Port             string `env:"POST" envDefault:"8002"`
	ServerParameters keepalive.ServerParameters
}

func (a *Grpc) BuildHostPort() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

func init() {
	err := env.Parse(&Config)
	if err != nil {
		log.Fatal().Err(err).Msg("env parse failed")
	}

	Config.App.Grpc.ServerParameters = keepalive.ServerParameters{
		MaxConnectionIdle: time.Duration(1) * time.Minute,
		Timeout:           time.Duration(30) * time.Second,
		MaxConnectionAge:  time.Duration(1) * time.Minute,
		Time:              time.Duration(1) * time.Minute,
	}

}
