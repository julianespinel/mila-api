package admin

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
)

func AddRoutes(router router.Party) router.Party {
	router.Get("/ping", pong)
	return router
}

func pong(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	responseBody := map[string]string{"message": "pong"}
	ctx.JSON(responseBody)
}
