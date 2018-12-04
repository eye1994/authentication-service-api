package application

import (
	"fmt"
	"net/http"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	validator "gopkg.in/validator.v2"
)

// Create with()
func Create(c echo.Context) (err error) {
	administrator := c.Get("Administrator").(*repository.Administrator)
	params := new(repository.ApplicationParams)
	if err = c.Bind(params); err != nil {
		return
	}

	err = validator.Validate(params)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.(validator.ErrorMap))
	}

	publicToken, err1 := uuid.NewUUID()
	privateToken, err2 := uuid.NewUUID()

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusInternalServerError, &utils.ErrorMessage{Error: "Internal server error"})
	}

	application := repository.Application{
		URL:         params.URL,
		Name:        params.Name,
		PublicToken: publicToken.String(),
		SecretToken: privateToken.String(),
	}

	result := repository.DB.Model(&administrator).Association("Applications").Append(&application)
	if result.Error != nil {
		fmt.Printf("\n%v\n", result.Error)
		return c.JSON(http.StatusUnprocessableEntity, result.Error)
	}

	return c.JSON(http.StatusOK, application)
}
