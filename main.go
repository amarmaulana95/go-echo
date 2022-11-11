package main

import (
	"github.com/amarmaulana95/go-echo/routes"
)

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":9001"))
}
