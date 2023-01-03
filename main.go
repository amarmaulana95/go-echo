package main

import (
	"net/http"
	"os"

	"github.com/amarmaulana95/go-echo/config"
	k "github.com/amarmaulana95/go-echo/middleware"
	"github.com/amarmaulana95/go-echo/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// var err error
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	db := config.DatabaseConnect()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	k.WebSecurityConfig(e)

	routes.NewRoute(db, e)

	e.Static("/static", "static")

	// e.Logger.Fatal(e.Start(":8080"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
