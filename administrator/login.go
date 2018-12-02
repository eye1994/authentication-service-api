package administrator

import (
	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// Login with()
func Login(ctx iris.Context) {
	params := &repository.LoginAdministratorParams{}
	err := ctx.ReadJSON(params)
	if err != nil {
		ctx.JSON(err.Error())
		return
	}

	var administrator repository.Administrator
	result := repository.DB.Where(&repository.Administrator{Email: params.Email}).First(&administrator)
	if result.Error != nil && result.RecordNotFound() {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(context.Map{"error": "Invalid email or password"})
		return
	}

	ok, err := utils.ValidatePassword(params.Password, []byte(administrator.PasswordHash))
	if !ok {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(context.Map{"error": "Invalid email or password"})
		return
	}

	tokenString, error := utils.SignJwt(int(administrator.ID))
	if error != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(error.Error())
		return
	}

	ctx.Header("Auhtnetication", tokenString)
	ctx.JSON(administrator)
}
