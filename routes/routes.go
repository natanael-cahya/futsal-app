package routes

// ini file untuk define semua routing
// yang nantinya akan merefer ke masing masing route domain
import (
	"net/http"

	"github.com/labstack/echo"
)

func API(e *echo.Echo) {
	e.GET("health", func(c echo.Context) error {

		return c.JSON(http.StatusOK, map[string]string{"Status": "Normal"})

	})
}
