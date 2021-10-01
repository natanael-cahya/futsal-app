package main

import (
	"futsal-app/app"

	"github.com/labstack/echo"
)

func main() {
	// disini panggil app.Start()
	// buat nge start server
	e := echo.New()
	app.Start()
	e.Start(":3000")
}
