package main

import (
	"futsal-app/app"
	"futsal-app/routes"
)

func main() {
	// disini panggil app.Start()
	// buat nge start server
	app.Start()
	routes.HealthChecks()

}
