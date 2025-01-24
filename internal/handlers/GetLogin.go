package handlers

import (
	"net/http"

	"github.com/Michael-Ralph/auth-session-redirect/internal/templates"
	"github.com/labstack/echo/v4"
)

func GetLogin(c echo.Context) error {
	p := templates.Login("Login")
	return Render(c, http.StatusOK, p)
}
