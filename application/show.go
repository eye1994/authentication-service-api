package application

import (
	"github.com/eye1994/authentication-service-api/repository"
	"github.com/labstack/echo"
	"net/http"
)

// Show with()
func Show(c echo.Context) (err error) {
	application := c.Get("Application").(*repository.Application)
	return c.JSON(http.StatusOK, application)
}
