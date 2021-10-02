package main

import (
	"futsal-app/app"
	"futsal-app/routes"

	"github.com/labstack/echo"
)

func main() {
	// disini panggil app.Start()
	// buat nge start server
	e := echo.New()
	app.Start()
	routes.HealthChecks()
	//http.Handle("/status", h.Handler())
	e.Start(":3000")
}
