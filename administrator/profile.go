package administrator

import (
	"net/http"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/labstack/echo"
)

// Profile with()
func Profile(c echo.Context) (err error) {
	administrator := c.Get("Administrator").(*repository.Administrator)
	return c.JSON(http.StatusOK, administrator)
}
