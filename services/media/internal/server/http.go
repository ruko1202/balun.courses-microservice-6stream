package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/ruko1202/balun.courses-microservice-6stream/media/internal/app/healthcheck"
	"github.com/ruko1202/balun.courses-microservice-6stream/media/internal/app/swagger"
	"github.com/ruko1202/balun.courses-microservice-6stream/media/internal/config"
)

func initHTTPServer() *echo.Echo {
	s := echo.New()
	initMiddlewares(s)
	swagger.New(s.Group("/docs"))

	return s
}

func initStatusHTTPServer() *echo.Echo {
	s := echo.New()
	initMiddlewares(s)
	healthcheck.New(s.Group(""))

	return s
}

func initMiddlewares(app *echo.Echo) {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
}

func runHTTPServer(name string, restCfg config.RestConfig, s *echo.Echo) error {
	log.Info().Msgf("Start http '%s' server: '%s'", name, restCfg.BuildHostPort())
	err := s.Start(restCfg.BuildHostPort())
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return errors.Wrapf(err, "failed to start '%s' server", name)
	}

	return nil
}
