package app

import (
	"github.com/labstack/echo"
)

// file ini untuk init semua keperluan api
// seperti koneksi database, routing, dll

func Start() {
	// panggil function ConnectToDB
	ConnectToDB()
	e := echo.New()
	e.Start(":3000")

}
