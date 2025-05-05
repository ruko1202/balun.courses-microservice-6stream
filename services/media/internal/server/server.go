package server

import (
	"context"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ruko1202/balun.courses-microservice-6stream/media/internal/config"
)

type Server struct {
	cfg config.App

	httpServer       *echo.Echo
	statusHTTPServer *echo.Echo

	netListener net.Listener
	grpcServer  *grpc.Server
}

func NewServer(ctx context.Context, cfg config.App) (*Server, error) {
	gmux := runtime.NewServeMux()

	httpServer := initHTTPServer()
	httpServer.Any("/*", echo.WrapHandler(gmux))

	grpcServer, err := initGRPCServer(ctx, gmux, cfg.Grpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize gRPC server")
	}

	statusHTTPServer := initStatusHTTPServer()

	return &Server{
		cfg:              cfg,
		httpServer:       httpServer,
		statusHTTPServer: statusHTTPServer,
		grpcServer:       grpcServer,
	}, nil
}

func (s *Server) Run() {
	go func() {
		err := runHTTPServer("gateway", &s.cfg.Rest, s.httpServer)
		if err != nil {
			log.Error().Err(err).Msg("failed to Start server")
		}
	}()

	go func() {
		err := runHTTPServer("status", &s.cfg.Status, s.statusHTTPServer)
		if err != nil {
			log.Error().Err(err).Msg("failed to Start server")
		}
	}()

	go func() {
		reflection.Register(s.grpcServer)

		l, err := net.Listen("tcp", s.cfg.Grpc.BuildHostPort())
		if err != nil {
			log.Fatal().Err(err).Msg("failed to listen")
		}
		s.netListener = l

		log.Info().Msgf("Start grpc server: %s", s.cfg.Grpc.BuildHostPort())
		if err := s.grpcServer.Serve(l); err != nil {
			log.Fatal().Err(err).Msg("Failed running gRPC server")
		}
	}()
}

func (s *Server) Shutdown(ctx context.Context) {
	log.Info().Msg("Shutdown http server...")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("failed to shutdown http server")
	}
	log.Info().Msgf("http server shutdown")

	log.Info().Msg("Shutdown grpc server...")
	s.netListener.Close()
	s.grpcServer.GracefulStop()

	log.Info().Msgf("grpc server shutdown")
}
