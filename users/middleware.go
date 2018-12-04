package user

import (
	"net/http"
	"strconv"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/labstack/echo"
)

// Middleware with()
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, &utils.ErrorMessage{Error: "Internal server error"})
		}
		application := c.Get("Application").(*repository.Application)

		var users []*repository.User
		repository.DB.Model(application).Related(&users, "Users")

		var user *repository.User
		for _, item := range users {
			if item.ID == uint(id) {
				user = item
			}
		}

		if user == nil {
			return c.JSON(http.StatusNotFound, &utils.ErrorMessage{Error: "User not found"})
		}

		c.Set("ApplicationUser", user)
		return next(c)
	}
}
