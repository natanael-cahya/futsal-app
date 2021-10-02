package routes

// ini file untuk define semua routing
// yang nantinya akan merefer ke masing masing route domain
import (
	"net/http"
	"time"

	"github.com/hellofresh/health-go/v4"
	healthMysql "github.com/hellofresh/health-go/v4/checks/mysql"
)

func HealthChecks() {
	h, _ := health.New()
	h.Register(health.Config{
		Name:      "Mysql-Checks",
		Timeout:   time.Second * 2,
		SkipOnErr: false,
		Check: healthMysql.New(healthMysql.Config{
			DSN: "root@tcp(127.0.0.1:3306)/db_futsal?charset=utf8mb4",
		}),
	})
	http.Handle("/status", h.Handler())
	http.ListenAndServe(":3000", nil)
}
