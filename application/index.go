package application

import (
	"net/http"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/labstack/echo"
)

// Index with()
func Index(c echo.Context) (err error) {
	administrator := c.Get("Administrator").(*repository.Administrator)
	var applications []*repository.Application
	repository.DB.Model(&administrator).Related(&applications, "Applications")
	return c.JSON(http.StatusOK, applications)
}
