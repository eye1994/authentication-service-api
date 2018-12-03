package administrator

import (
	"net/http"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/labstack/echo"
)

// Register with()
func Register(c echo.Context) (err error) {
	params := new(repository.RegisterAdministratorParams)
	if err = c.Bind(params); err != nil {
		return
	}

	if ok, err := params.Validate(); !ok {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	var result []repository.Administrator
	repository.DB.Where(&repository.Administrator{Email: params.Email}).Find(&result)
	if len(result) > 0 {
		return c.JSON(http.StatusUnprocessableEntity, &utils.ErrorMessage{Error: "Email address is taken by another user"})
	}

	administrator, err := params.ToModel()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	db := repository.DB.Create(&administrator)
	if db.Error != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, administrator)
}
