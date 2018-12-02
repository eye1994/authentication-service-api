package administrator

import (
	"github.com/eye1994/authentication-service-api/repository"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// Register with()
func Register(ctx iris.Context) {
	params := &repository.RegisterAdministratorParams{}
	err := ctx.ReadJSON(params)
	if err != nil {
		ctx.JSON(err.Error())
		return
	}

	if ok, err := params.Validate(); !ok {
		ctx.StatusCode(iris.StatusUnprocessableEntity)
		ctx.JSON(context.Map{"response": err})
		return
	}

	var result []repository.Administrator
	repository.DB.Where(&repository.Administrator{Email: params.Email}).Find(&result)
	if len(result) > 0 {
		ctx.JSON(context.Map{"error": "Email address is taken by another user"})
		return
	}

	administrator, err := params.ToModel()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(context.Map{"error": err.Error()})
		return
	}

	db := repository.DB.Create(&administrator)
	if db.Error != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(context.Map{"error": db.Error.Error()})
		return
	}

	ctx.JSON(administrator)
}
