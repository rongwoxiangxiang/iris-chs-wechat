package controllers

import (
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/kataras/iris"
)

type ServiceController struct {
	Ctx iris.Context
}

var (
	msgHandler core.Handler
	msgServers map[string]*core.Server
)

func init() {
	mux := core.NewServeMux()
	msgHandler = mux
	msgServers = make(map[string]*core.Server)
}

func Service(ctx iris.Context) {

}
