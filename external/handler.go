package external

import (
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handlers with()
func Handlers(e *echo.Echo) {
	applicationGroup := e.Group("/api")
	applicationGroup.POST("/login", Login)

	authenticatedGroup := e.Group("/api/profile")
	authenticatedGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  utils.JwtSecretKey,
		TokenLookup: "header:Authorization",
	}))
	authenticatedGroup.Use(Middleware)
	authenticatedGroup.GET("", Profile)
}
