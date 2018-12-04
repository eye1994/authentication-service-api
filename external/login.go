package external

import (
	"net/http"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/labstack/echo"
)

// Login with()
func Login(c echo.Context) (err error) {
	params := new(repository.LoginUserParams)
	if err = c.Bind(params); err != nil {
		return
	}

	publicToken := c.Request().Header.Get("ApplicationToken")
	var application repository.Application
	search := &repository.Application{PublicToken: publicToken}
	result := repository.DB.Where(search).First(&application)
	if result.Error != nil && result.RecordNotFound() {
		return c.JSON(http.StatusNotFound, &utils.ErrorMessage{Error: "Application not found"})
	}

	var users []*repository.User
	repository.DB.Model(application).Related(&users, "Users")

	var user *repository.User
	for _, item := range users {
		if item.Email == params.Email {
			user = item
		}
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, &utils.ErrorMessage{Error: "User or not found for the application"})
	}

	ok, err := utils.ValidatePassword(params.Password, []byte(user.PasswordHash))
	if !ok {
		return c.JSON(http.StatusUnauthorized, &utils.ErrorMessage{Error: "Invalid email or password"})
	}

	tokenString, error := utils.SignJwt(int(user.ID))
	if error != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("Authorization", tokenString)
	return c.JSON(http.StatusOK, user)
}
