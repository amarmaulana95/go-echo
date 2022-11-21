package main

import (
	"net/http"

	"github.com/amarmaulana95/go-echo/config"
	"github.com/amarmaulana95/go-echo/middleware"
	"github.com/amarmaulana95/go-echo/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := config.DatabaseConnect()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	middleware.WebSecurityConfig(e)

	routes.NewRoute(db, e)

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":8080"))
}
