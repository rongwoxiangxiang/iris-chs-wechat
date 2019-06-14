package controllers

import (
	"chs/controllers/Handlers"
	"chs/models/interface"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/kataras/iris"
	"iris/common"
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
	mux.DefaultMsgHandleFunc(Handlers.DefaultMsgHandler)
	mux.DefaultEventHandleFunc(Handlers.DefaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, Handlers.DefaultTextMsgHandler)
	msgServers = make(map[string]*core.Server)
}

func Service(ctx iris.Context) {
	flag := ctx.Params().Get("flag")
	msgServer := getMsgServer(flag)
	if msgServer == nil {
		ctx.Application().Logger().Warn("Wechat service get mp server err wechat flag : %v", flag)
		return
	}
	msgServer.ServeHTTP(ctx.ResponseWriter(), ctx.Request(), nil)
}

func getMsgServer(flag string) *core.Server {
	if service, ok := msgServers[flag]; ok == true {
		return service
	}
	wechat := _interface.GetWechatServiceR().GetByFlag(flag)
	if wechat == nil {
		return nil
	}
	msgServer := core.NewServer("", wechat.Appid, wechat.Token, wechat.EncodingAesKey, msgHandler, nil)
	if wechat.NeedSaveMen != common.NO_VALUE {
		msgServers[flag] = msgServer
	}

	return msgServer
}
