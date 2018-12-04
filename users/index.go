package user

import (
	"net/http"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/labstack/echo"
)

// Index with()
func Index(c echo.Context) (err error) {
	application := c.Get("Application").(*repository.Application)
	var users []*repository.User
	repository.DB.Model(&application).Related(&users, "Users")
	return c.JSON(http.StatusOK, users)
}
