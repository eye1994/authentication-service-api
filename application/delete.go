package application

import (
	"net/http"
	"strconv"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/labstack/echo"
)

// Delete with()
func Delete(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.ErrorMessage{Error: "Internal server error"})
	}

	application := c.Get("Application").(*repository.Application)
	repository.DB.Delete(application, &repository.Application{ID: uint(id)})

	return c.JSON(http.StatusOK, map[string]string{})
}
