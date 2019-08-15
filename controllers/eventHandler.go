package controllers

import (
	"chs/dao"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"log"
)

func menuClickEventHandler(ctx *core.Context) {
	log.Printf("收到菜单 click 事件:\n%s\n", ctx.MsgPlaintext)
	reply := dao.GetReplyServiceR().FindOne(&dao.ReplyModel{Wid: wxUser.Wid, ClickKey: ctx.MixedMsg.EventKey})
	if reply != nil && reply.Success != "" {
		resp := responseTextAndClick(reply, ctx.MixedMsg.MsgHeader)
		responseMsg(ctx, resp)
	}
}

func textMsgHandler(ctx *core.Context) {
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)
	msg := request.GetText(ctx.MixedMsg)
	reply := dao.GetReplyServiceR().FindOne(&dao.ReplyModel{Wid: wxUser.Wid, Alias: msg.Content})
	if reply != nil && reply.Success != "" {
		resp := responseTextAndClick(reply, msg.MsgHeader)
		responseMsg(ctx, resp)
	} else {
		defaultTextMsgHandler(ctx)
	}
}
