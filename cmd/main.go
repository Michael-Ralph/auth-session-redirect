package main

import (
	"github.com/Michael-Ralph/auth-session-redirect/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  // Logs each HTTP request
	e.Use(middleware.Recover()) // Recovers from panics

	// Routes
	e.GET("/index", handlers.GetLogin)
	e.GET("/home", handlers.Gethome)
	e.POST("/home", handlers.Postlogin)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
