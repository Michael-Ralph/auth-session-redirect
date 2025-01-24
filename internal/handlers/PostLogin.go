package handlers

import (
	"net/http"

	"github.com/Michael-Ralph/auth-session-redirect/internal/templates"
	"github.com/labstack/echo/v4"
)

func Postlogin(c echo.Context) error {
	h := templates.Home()
	// Redirect the user to "/home"
	if err := c.Redirect(http.StatusSeeOther, "/home"); err != nil {
		return err
	}
	// Render the home template
	return Render(c, http.StatusOK, h)
}
