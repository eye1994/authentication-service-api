package administrator

import (
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handlers with()
func Handlers(e *echo.Echo) {
	administratorGroup := e.Group("/administrator")

	administratorGroup.POST("/login", Login)
	administratorGroup.POST("", Register)

	authenticatedGroup := e.Group("/administrator/profile")
	authenticatedGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  utils.JwtSecretKey,
		TokenLookup: "header:Authorization",
	}))
	authenticatedGroup.Use(Middleware)
	authenticatedGroup.GET("", Profile)
}
