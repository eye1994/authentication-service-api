package administrator

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/eye1994/authentication-service-api/repository"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// Profile with()
func Profile(ctx iris.Context) {
	token := ctx.Values().Get("jwt").(*jwt.Token)
	if token.Valid != true {
		ctx.JSON(context.Map{"error": "Invalid token"})
		return
	}

	ID, ok := token.Claims.(jwt.MapClaims)["ID"].(float64)
	if !ok {
		ctx.JSON(context.Map{"error": "Invalid token"})
		return
	}

	var administrator repository.Administrator
	search := &repository.Administrator{ID: uint(ID)}
	result := repository.DB.Where(search).First(&administrator)
	if result.Error != nil && result.RecordNotFound() {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(context.Map{"error": "Invalid token"})
		return
	}

	ctx.JSON(administrator)
}
