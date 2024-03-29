package controllers

import (
	"chs/common"
	"chs/dao"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/kataras/iris"
)

var (
	msgHandler core.Handler
	msgServers map[string]*core.Server
	wechats    map[string]int64
)

func init() {
	mux := core.NewServeMux()
	msgHandler = mux
	mux.UseFunc(onStart)
	mux.DefaultMsgHandleFunc(defaultMsgHandler)
	mux.DefaultEventHandleFunc(defaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, textMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, menuClickEventHandler)
	msgServers = make(map[string]*core.Server)
	wechats = make(map[string]int64)
}

func Service(ctx iris.Context) {
	flag := ctx.Params().Get("flag")
	msgServer := getMsgServer(flag)
	if msgServer == nil {
		ctx.Application().Logger().Warn("Wechat service get mp server err wechat flag : %v", flag)
		return
	}
	query := ctx.Request().URL.Query()
	query.Add("flag", flag)
	msgServer.ServeHTTP(ctx.ResponseWriter(), ctx.Request(), query)
}

func getMsgServer(flag string) *core.Server {
	if service, ok := msgServers[flag]; ok == true {
		return service
	}
	wechat := dao.GetWechatServiceR().GetByFlag(flag)
	if wechat == nil {
		return nil
	}
	wechats[flag] = wechat.Id
	msgServer := core.NewServer("", wechat.Appid, wechat.Token, wechat.EncodingAesKey, msgHandler, nil)
	if wechat.NeedSaveMen != common.NO_VALUE {
		msgServers[flag] = msgServer
	}
	return msgServer
}
