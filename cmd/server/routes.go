package main

import (
	"github.com/hyuma/go-echo-jwt/app/frameworks/server/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Routes(e *echo.Echo) {
	// Visible Top page
	e.Static("/", "app/frameworks/server/public/index.html")

	e.POST("/login", handlers.PostUserLogin())

	// JWT が必要となる Group
	g := e.Group("/profile")
	g.Use(middleware.JWT([]byte("secret")))
	g.GET("", handlers.GetUserProfile())
}
