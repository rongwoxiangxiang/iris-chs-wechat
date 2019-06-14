package Handlers

import (
	"chs/config"
	"chs/modules/ai/qq"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"log"
)

var Flag string

func DefaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func DefaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func DefaultTextMsgHandler(ctx *core.Context) {
	var (
		msg    *request.Text
		answer string
		resp   *response.Text
	)

	log.Printf("AI智能闲聊:\n%s\n", ctx.MsgPlaintext)
	msg = request.GetText(ctx.MixedMsg)
	answer = qq.NewNlpTextchat(qq.TomlConfiguration(config.Conf)).Question(msg.Content, msg.FromUserName).Answer()
	if answer != "" {
		resp = response.NewText(msg.MsgHeader.FromUserName, msg.MsgHeader.ToUserName, msg.MsgHeader.CreateTime, answer)
		if len(ctx.AESKey) == 0 {
			ctx.RawResponse(resp)
			return
		}
		ctx.AESResponse(resp, 0, "", nil)
		return
	}
	ctx.NoneResponse()
}
