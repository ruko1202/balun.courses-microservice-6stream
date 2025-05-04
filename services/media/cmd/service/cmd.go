package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"github.com/ruko1202/balun.courses-microservice-6stream/media/internal/config"
	"github.com/ruko1202/balun.courses-microservice-6stream/media/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	s, err := server.NewServer(ctx, config.Config.App)
	if err != nil {
		stop()
		log.Fatal().Err(err) //nolint:gocritic
	}

	s.Run()

	<-ctx.Done()
	log.Info().Msg("Shutdown server...")
	s.Shutdown(ctx)

}
