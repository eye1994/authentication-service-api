package user

import (
	"github.com/eye1994/authentication-service-api/administrator"
	"github.com/eye1994/authentication-service-api/application"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handlers with()
func Handlers(e *echo.Echo) {
	users := e.Group("/application/:id/users")

	users.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  utils.JwtSecretKey,
		TokenLookup: "header:Authorization",
	}))
	users.Use(administrator.Middleware)
	users.Use(application.Middleware)
	users.GET("", Index)
	users.POST("", Create)

	rudGroup := users.Group("/:userId")
	rudGroup.Use(Middleware)
	rudGroup.GET("", Show)
	rudGroup.PUT("", Update)
	rudGroup.DELETE("", Delete)
}
