package administrator

import (
	"net/http"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/labstack/echo"
)

// Login with()
func Login(c echo.Context) (err error) {
	params := new(repository.LoginAdministratorParams)
	if err = c.Bind(params); err != nil {
		return
	}

	var administrator repository.Administrator
	result := repository.DB.Where(&repository.Administrator{Email: params.Email}).First(&administrator)
	if result.Error != nil && result.RecordNotFound() {
		return c.JSON(http.StatusUnauthorized, &utils.ErrorMessage{Error: "Invalid email or password"})
	}

	ok, err := utils.ValidatePassword(params.Password, []byte(administrator.PasswordHash))
	if !ok {
		return c.JSON(http.StatusUnauthorized, &utils.ErrorMessage{Error: "Invalid email or password"})
	}

	tokenString, error := utils.SignJwt(int(administrator.ID))
	if error != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("Authorization", tokenString)
	return c.JSON(http.StatusOK, administrator)
}
