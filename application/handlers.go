package application

import (
	"github.com/eye1994/authentication-service-api/administrator"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handlers with()
func Handlers(e *echo.Echo) {
	applicationGroup := e.Group("/application")

	applicationGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  utils.JwtSecretKey,
		TokenLookup: "header:Authorization",
	}))
	applicationGroup.Use(administrator.Middleware)

	applicationGroup.GET("", Index)
	applicationGroup.POST("", Create)

	rudGroup := applicationGroup.Group("/:id")
	rudGroup.Use(Middleware)
	rudGroup.GET("", Show)
	rudGroup.PUT("", Update)
	rudGroup.DELETE("", Delete)
}
