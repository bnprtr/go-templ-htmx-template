package views

import (
	"app/templates/components"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		components.Layout(Index()).Render(c.Request().Context(), c.Response())
		return nil
	})
	e.POST("/ping", Ping)
}

func Ping(c echo.Context) error {
	components.Ping().Render(c.Request().Context(), c.Response())
	return nil
}
