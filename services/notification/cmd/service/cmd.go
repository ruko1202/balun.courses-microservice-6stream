package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ruko1202/balun.courses-microservice-6stream/notification/internal/app/service"
	"github.com/ruko1202/balun.courses-microservice-6stream/notification/internal/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	app := echo.New()
	initMiddlewares(app)

	{
		service.New(app.Group(""))
	}

	go func() {
		defer stop()
		slog.Info("Start server...")
		err := app.Start(config.Config.App.BuildHostPort())
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed to Start server", "error", err)
		}
	}()

	<-ctx.Done()
	slog.Info("Shutdown server...")
	if err := app.Shutdown(ctx); err != nil {
		slog.Error("failed to Shutdown server", "error", err)
	}

}

func initMiddlewares(app *echo.Echo) {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
}
