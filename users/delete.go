package user

import (
	"net/http"
	"strconv"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/labstack/echo"
)

// Delete with()
func Delete(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.ErrorMessage{Error: "Internal server error"})
	}

	user := c.Get("ApplicationUser").(*repository.User)
	repository.DB.Delete(user, &repository.User{ID: uint(id)})

	return c.JSON(http.StatusOK, map[string]string{})
}
