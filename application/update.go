package application

import (
	"github.com/eye1994/authentication-service-api/repository"
	"github.com/labstack/echo"
	"gopkg.in/validator.v2"
	"net/http"
)

// Update with()
func Update(c echo.Context) (err error) {
	application := c.Get("Application").(*repository.Application)

	params := new(repository.ApplicationParams)
	if err = c.Bind(params); err != nil {
		return
	}

	err = validator.Validate(params)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.(validator.ErrorMap))
	}

	updateResult := repository.DB.Model(&application).Updates(params)
	if updateResult.Error != nil {
		return c.JSON(http.StatusUnprocessableEntity, updateResult.Error)
	}

	return c.JSON(http.StatusOK, application)
}
