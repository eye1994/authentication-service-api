package main

import (
	"github.com/eye1994/authentication-service-api/administrator"
	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func myHandler(ctx iris.Context) {
	ctx.JSON(context.Map{"message": "Welcome User Micro Service"})

}

// EchoHandler with()
func EchoHandler() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/administrator", administrator.Register)
	e.POST("/administrator/login", administrator.Login)

	g := e.Group("/administrator/profile")
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  utils.JwtSecretKey,
		TokenLookup: "header:Authorization",
	}))
	g.GET("", administrator.Profile)

	return e
}

func main() {
	defer repository.DB.Close()
	app := EchoHandler()
	app.Start(":3002")
}
