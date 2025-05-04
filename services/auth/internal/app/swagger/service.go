package swagger

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ruko1202/swaggerui"

	doc "github.com/ruko1202/balun.courses-microservice-6stream/auth"
)

type app struct{}

func New(gr *echo.Group) {
	a := &app{}

	for route, handler := range map[*echo.Route]struct {
		handler    echo.HandlerFunc
		middleware []echo.MiddlewareFunc
	}{
		&echo.Route{Path: "", Name: "Swagger docs", Method: http.MethodGet}: {
			handler: a.swaggerMainPage,
		},
		&echo.Route{Path: "/template/*", Name: "Template Service", Method: http.MethodGet}: {
			handler: a.swaggerServicePage("/docs/template"),
		},
	} {
		gr.Add(route.Method, route.Path, handler.handler, handler.middleware...)
	}
}

func (a *app) swaggerMainPage(c echo.Context) error {
	return c.HTML(http.StatusOK, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Swagger docs</title>
</head>
<body>
    <h1>MetriQA swagger links</h1>
    <ul>
        <li><a href="/docs/template/">TemplateService</a></li>
    </ul>
</body>
</html>`)
}

func (a *app) swaggerServicePage(prefix string) func(c echo.Context) error {
	return echo.WrapHandler(
		http.StripPrefix(prefix, swaggerui.Handler(doc.SwaggerDocService)),
	)
}
