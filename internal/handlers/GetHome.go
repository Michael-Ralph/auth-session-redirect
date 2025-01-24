package handlers

import (
	"net/http"

	"github.com/Michael-Ralph/auth-session-redirect/internal/templates"
	"github.com/labstack/echo/v4"
)

func Gethome(c echo.Context) error {
	p := templates.Home()
	return Render(c, http.StatusOK, p)
}
