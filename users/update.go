package user

import (
	"net/http"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/labstack/echo"
	"gopkg.in/validator.v2"
)

// Update with()
func Update(c echo.Context) (err error) {
	user := c.Get("ApplicationUser").(*repository.User)

	params := new(repository.UserParams)
	if err = c.Bind(params); err != nil {
		return
	}

	err = validator.Validate(params)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.(validator.ErrorMap))
	}

	userUpdate, err := params.ToModel()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	updateResult := repository.DB.Model(&user).Updates(userUpdate)
	if updateResult.Error != nil {
		return c.JSON(http.StatusUnprocessableEntity, updateResult.Error)
	}

	return c.JSON(http.StatusOK, user)
}
