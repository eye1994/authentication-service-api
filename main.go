package main

import (
	"github.com/eye1994/authentication-service-api/administrator"
	"github.com/eye1994/authentication-service-api/application"
	"github.com/eye1994/authentication-service-api/repository"
	user "github.com/eye1994/authentication-service-api/users"
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

	administrator.Handlers(e)
	application.Handlers(e)
	user.Handlers(e)

	return e
}

func main() {
	defer repository.DB.Close()
	app := EchoHandler()
	app.Start(":3002")
}
