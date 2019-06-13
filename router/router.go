package router

import (
	"chs/controllers"
	"github.com/kataras/iris"
)

func Routes(app *iris.Application) {
	app.Any("/service/{flag:string}", controllers.Service)
}
