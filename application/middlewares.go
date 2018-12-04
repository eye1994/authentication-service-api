package application
import (
	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// Middleware with()
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, &utils.ErrorMessage{Error: "Internal server error"})
		}
		administrator := c.Get("Administrator").(*repository.Administrator)

		var applications []*repository.Application
		repository.DB.Model(administrator).Related(&applications, "Applications")

		var application *repository.Application
		for _, item := range applications {
			if item.ID == uint(id) {
				application = item
			}
		}

		if application == nil {
			return c.JSON(http.StatusNotFound, &utils.ErrorMessage{Error: "Application not found"})
		}

		c.Set("Application", application)
		return next(c)
	}
}