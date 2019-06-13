package bootstrap

import (
	"chs/common"
	"chs/config"
	"chs/router"
	"chs/util"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/pelletier/go-toml"
	"log"
)

type Application struct {
	Application *iris.Application
	Config      *toml.Tree
	Debug       bool
}

func App() *Application {
	application := &Application{Application: iris.New()}
	application.initConfig()
	application.isDebug()
	application.initRouter()
	application.initLogger()
	application.initDatabase()

	return application
}

func (application *Application) Run() {
	err := application.Application.Run(application.runner(), application.configuration())
	if err != nil {
		log.Fatal("Init application run err : %v", err)
	}
}

func (application *Application) runner() iris.Runner {
	return iris.Addr(application.Config.GetDefault("application.address", "8888").(string))
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
		ctx.JSON(iris.Map{"version": application.Config.GetDefault("application.version", "0.0.1")})
	})
	application.Application.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(common.ErrClientParams)
	})
	application.Application.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON(common.ErrUnKnow)
	})
	router.Routes(application.Application)
}

func (application *Application) initLogger() {
	application.Application.Use(logger.New(logger.DefaultConfig()))
}

func (application *Application) initConfig() {
	//TODO
	config.InitConfig()
	application.Config = config.Conf
}

func (application *Application) initDatabase() {
	sources := application.Config.Get("source")
	if sources == nil {
		log.Println("Init application orm failed: database source null")
		return
	}
	config.SetDbDebug(application.Debug)
	config.InitStoreDb(sources.(*toml.Tree))
}

func (application *Application) isDebug() {
	debug := application.Config.Get("application.debug")
	if debug != nil {
		application.Debug = debug.(bool)
	}
}
