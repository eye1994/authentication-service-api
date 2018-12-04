package user

import (
	"fmt"
	"net/http"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/labstack/echo"
	validator "gopkg.in/validator.v2"
)

// Create with()
func Create(c echo.Context) (err error) {
	application := c.Get("Application").(*repository.Application)
	params := new(repository.UserParams)
	if err = c.Bind(params); err != nil {
		return
	}

	err = validator.Validate(params)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.(validator.ErrorMap))
	}

	user, err := params.ToModel()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result := repository.DB.Model(&application).Association("Users").Append(user)
	if result.Error != nil {
		fmt.Printf("\n%v\n", result.Error)
		return c.JSON(http.StatusUnprocessableEntity, result.Error)
	}

	return c.JSON(http.StatusOK, user)
}
