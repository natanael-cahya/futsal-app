package app

import (
	"futsal-app/routes"

	"github.com/labstack/echo"
)

// file ini untuk init semua keperluan api
// seperti koneksi database, routing, dll

func Start() {
	// panggil function ConnectToDB
	ConnectToDB()
	e := echo.New()

	routes.API(e)

	go func() {
		if err := e.Start(":3000"); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()

}
