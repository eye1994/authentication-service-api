package administrator

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/labstack/echo"
)

// Middleware with()
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		if token.Valid != true {
			return c.String(http.StatusUnauthorized, "Invalid token")
		}

		ID, ok := token.Claims.(jwt.MapClaims)["ID"].(float64)
		if !ok {
			return c.JSON(http.StatusUnauthorized, &utils.ErrorMessage{Error: "Invalid token"})
		}

		var administrator repository.Administrator
		search := &repository.Administrator{ID: uint(ID)}
		result := repository.DB.Where(search).First(&administrator)
		if result.Error != nil && result.RecordNotFound() {
			return c.JSON(http.StatusUnauthorized, &utils.ErrorMessage{Error: "Invalid token"})
		}

		c.Set("Administrator", &administrator)

		return next(c)
	}
}
