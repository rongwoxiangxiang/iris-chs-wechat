package bootstrap

import (
	"chs/common"
	"chs/config"
	"chs/router"
	"chs/util"
	"github.com/kataras/iris"
	"log"
)

type Application struct {
	Application *iris.Application
}

func App() *Application {
	application := &Application{Application: iris.New()}
	application.initRouter()
	return application
}

func (application *Application) Run() {
	err := application.Application.Run(application.runner(), application.configuration())
	if err != nil {
		log.Fatal("Init application run err : %v", err)
	}
}

func (application *Application) runner() iris.Runner {
	return iris.Addr(config.Conf.GetDefault("application.address", "8888").(string))
}

func (application *Application) configuration() iris.Configurator {
	serverConfigExist := util.CheckFileExist("server.yml")
	if serverConfigExist {
		return iris.WithConfiguration(iris.TOML("server.toml"))
	}
	return iris.WithConfiguration(iris.DefaultConfiguration())
}

func (application *Application) initRouter() {
	application.Application.Get("/check", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"version": config.Conf.GetDefault("application.version", "0.0.1")})
	})
	application.Application.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(common.ErrClientParams)
	})
	application.Application.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON(common.ErrUnKnow)
	})
	router.Routes(application.Application)
}
