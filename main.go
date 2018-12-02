package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/eye1994/authentication-service-api/administrator"
	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func myHandler(ctx iris.Context) {
	ctx.JSON(context.Map{"message": "Welcome User Micro Service"})

}

// NewApp with()
func NewApp() *iris.Application {
	app := iris.New()

	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return utils.JwtSecretKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	app.Post("/administrator", administrator.Register)
	app.Post("/administrator/login", administrator.Login)
	app.Use(jwtHandler.Serve)
	app.Get("/administrator/profile", administrator.Profile)

	app.Run(iris.Addr("0.0.0.0:3002"))

	return app
}

func main() {
	defer repository.DB.Close()
	NewApp()
}
