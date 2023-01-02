package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amarmaulana95/go-echo/config"
	"github.com/amarmaulana95/go-echo/middleware"
	"github.com/amarmaulana95/go-echo/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	var err error
	e := echo.New()

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	db := config.DatabaseConnect()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	middleware.WebSecurityConfig(e)

	routes.NewRoute(db, e)

	e.Static("/static", "static")

	// e.Logger.Fatal(e.Start(":8080"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
