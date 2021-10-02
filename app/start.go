package app

import (
	"context"
	"futsal-app/routes"
	"os"
	"os/signal"
	"time"

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

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
