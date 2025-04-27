package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type app struct{}

func New(gr *echo.Group) {
	a := &app{}

	for route, handler := range map[*echo.Route]struct {
		handler    echo.HandlerFunc
		middleware []echo.MiddlewareFunc
	}{
		&echo.Route{Path: "/liveness", Name: "Liveness", Method: http.MethodGet}: {
			handler: a.Liveness,
		},
		&echo.Route{Path: "/readiness", Name: "Readiness", Method: http.MethodGet}: {
			handler: a.Readiness,
		},
		&echo.Route{Path: "/version", Name: "Version", Method: http.MethodGet}: {
			handler: a.Version,
		},
	} {
		gr.Add(route.Method, route.Path, handler.handler, handler.middleware...)
	}
}
