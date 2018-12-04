package user

import (
	"net/http"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/labstack/echo"
)

// Show with()
func Show(c echo.Context) (err error) {
	user := c.Get("ApplicationUser").(*repository.User)
	return c.JSON(http.StatusOK, user)
}
