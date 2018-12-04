package external

import (
	"net/http"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/labstack/echo"
)

// Profile with()
func Profile(c echo.Context) (err error) {
	user := c.Get("ApplicationUser").(*repository.User)
	return c.JSON(http.StatusOK, user)
}
