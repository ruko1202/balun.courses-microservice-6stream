package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ruko1202/balun.courses-microservice-6stream/user-profile/internal/config/build"
)

func (a *app) Version(c echo.Context) error {
	return c.JSON(http.StatusOK, build.GetAppVersion())
}
