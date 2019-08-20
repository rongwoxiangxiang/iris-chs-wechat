package controllers

import (
	"chs/config"
	"chs/dao"
	"chs/modules/ai/qq"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"github.com/chanxuehong/wechat/mp/user"
	"log"
	"time"
)

var (
	wxUser *dao.WechatUserModel
)

func onStart(ctx *core.Context) {
	setWechatUser(ctx.QueryParams.Get("flag"), ctx.MixedMsg.FromUserName)
	if wxUser == nil {
		panic("onStart.setWechatUser err")
	}
	go func() {
		updateWxUserInfo(wxUser)
		insertRecord(wxUser, ctx.MixedMsg)
	}()
}

func defaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func defaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func subscribeHandler(ctx *core.Context) {
	log.Printf("收到关注事件:\n%s\n", ctx.MsgPlaintext)
	reply := dao.GetReplyServiceR().FindOne(&dao.ReplyModel{Wid: wxUser.Wid, Alias: string(ctx.MixedMsg.EventType)})
	if reply.Success != "" {
		resp := response.NewText(ctx.MixedMsg.FromUserName, ctx.MixedMsg.ToUserName, ctx.MixedMsg.CreateTime, reply.Success)
		responseMsg(ctx, resp)
	}
}

func unsubscribeHandler(ctx *core.Context) {
	log.Printf("收到取关事件:\n%s\n", ctx.MsgPlaintext)
	//TODO
	ctx.NoneResponse()
}

func defaultTextMsgHandler(ctx *core.Context) {
	log.Printf("AI智能闲聊:\n%s\n", ctx.MsgPlaintext)
	msg := request.GetText(ctx.MixedMsg)
	answer := qq.NewNlpTextchat(qq.TomlConfiguration(config.Conf)).Question(msg.Content, msg.FromUserName).Answer()
	if answer != "" {
		resp := response.NewText(msg.MsgHeader.FromUserName, msg.MsgHeader.ToUserName, msg.MsgHeader.CreateTime, answer)
		responseMsg(ctx, resp)
		return
	}
	ctx.NoneResponse()
}

func responseMsg(ctx *core.Context, response *response.Text) {
	if len(ctx.AESKey) == 0 {
		ctx.RawResponse(response)
		return
	}
	ctx.AESResponse(response, 0, "", nil)
}

/**
 * 插入用户操作
 */
func insertRecord(wechatUser *dao.WechatUserModel, msg *core.MixedMsg) {
	dao.GetRecordServiceW().Insert(&dao.RecordModel{
		Wid:     wechatUser.Wid,
		Wuid:    wechatUser.Id,
		Type:    string(msg.MsgType) + string(msg.EventType),
		Content: msg.Content + msg.EventKey + msg.MediaId,
	})
}

/**
 * 获取本地微信用户信息，并异步更新远程信息
 */
func setWechatUser(flag, openId string) {
	wid := wechats[flag]
	userWx, _ := dao.GetWechatUserServiceR().GetByWidAndOpenid(wid, openId)
	if userWx == nil {
		userWx = &dao.WechatUserModel{Openid: openId, Wid: wid}
		_, err := dao.GetWechatUserServiceW().Insert(userWx)
		if err != nil {
			userWx = nil
		}
	}
	wxUser = userWx
}

/**
 * 按条件拉取微信用户信息，尽量异步执行
 */
func updateWxUserInfo(wechatUser *dao.WechatUserModel) {
	if wechatUser != nil && (wechatUser.Nickname == "" || wechatUser.UpdatedAt.IsZero() || time.Now().After(wechatUser.UpdatedAt.Add(24*time.Hour))) {
		wechat := dao.GetWechatServiceR().GetById(wechatUser.Wid)
		accessTokenServer := core.NewDefaultAccessTokenServer(wechat.Appid, wechat.Appsecret, nil)
		userInfo, err := user.Get(core.NewClient(accessTokenServer, nil), wechatUser.Openid, "")
		if err == nil {
			dao.GetWechatUserServiceW().Update(&dao.WechatUserModel{
				Id:         wechatUser.Id,
				Nickname:   userInfo.Nickname,
				Sex:        userInfo.Sex,
				Province:   userInfo.Province,
				City:       userInfo.City,
				Country:    userInfo.Country,
				Language:   userInfo.Language,
				Headimgurl: userInfo.HeadImageURL,
				UpdatedAt:  time.Now(),
			})
		}

	}
}
