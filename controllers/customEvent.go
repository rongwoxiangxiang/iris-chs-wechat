package controllers

import (
	"chs/common"
	"chs/dao"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"strconv"
	"strings"
	"time"
)

func responseTextAndClick(reply *dao.ReplyModel, header core.MsgHeader) (msg *response.Text) {
	var replyMsg string
	switch reply.Type {
	case dao.REPLY_TYPE_TEXT:
		replyMsg = reply.Success
	case dao.REPLY_TYPE_CODE:
		replyMsg = doReplyCode(reply)
	case dao.REPLY_TYPE_CHECKIN:
		replyMsg = doReplyCheckin(reply)
	case dao.REPLY_TYPE_LUCKY:
		replyMsg = doReplyLucky(reply)
	}
	return response.NewText(
		header.FromUserName,
		header.ToUserName,
		header.CreateTime,
		replyMsg)
	return
}

func doReplyCode(reply *dao.ReplyModel) string {
	history, _ := dao.GetPrizeHistoryServiceR().GetByActivityWuId(reply.ActivityId, wxUser.Id)
	if history.Code != "" {
		return strings.Replace(reply.Success, "%code%", history.Code, 1)
	}
	prize, err := (dao.GetPrizeServiceW().ChooseOneUsedPrize(reply.ActivityId, dao.PRIZE_LEVEL_DEFAULT, 0))

	if err == common.ErrDataUnExist {
		return reply.NoPrizeReturn
	}
	if prize.Code != "" {
		_, err := dao.GetPrizeHistoryServiceW().Insert(&dao.PrizeHistoryModel{ActivityId: reply.ActivityId, Wuid: wxUser.Id, Code: prize.Code})
		if err != nil {
			return dao.PLEASE_TRY_AGAIN
		}
		return strings.Replace(reply.Success, "%code%", prize.Code, 1)
	}
	return dao.PLEASE_TRY_AGAIN
}

func doReplyCheckin(reply *dao.ReplyModel) string {
	checkin, err := dao.GetCheckinServiceR().GetCheckinByActivityWuid(reply.ActivityId, wxUser.Wid)
	if err != nil {
		return dao.CHECK_FAIL
	}
	lastCheckinDate := checkin.Lastcheckin.Format("2006-01-02")
	if lastCheckinDate == time.Now().Format("2006-01-02") {
		return strings.
			NewReplacer("%liner%", strconv.FormatInt(checkin.Liner, 10), "%total%", strconv.FormatInt(checkin.Total, 10)).
			Replace(reply.Success)
	}
	if lastCheckinDate == time.Now().Add(-24*time.Hour).Format("2006-01-02") { //连续签到
		checkin.Liner = checkin.Liner + 1
	} else { //重置连续签到数
		checkin.Liner = 1
	}
	checkin.Total = checkin.Total + 1
	checkin.Lastcheckin = time.Now()
	_, err = dao.GetCheckinServiceW().Update(checkin)
	if err != nil {
		return dao.CHECK_FAIL
	}
	return strings.
		NewReplacer("%liner%", strconv.FormatInt(checkin.Liner, 10), "%total%", strconv.FormatInt(checkin.Total, 10)).
		Replace(reply.Success)
}

func doReplyLucky(reply *dao.ReplyModel) string {
	activity := dao.GetActivityServiceR().GetById(reply.ActivityId)
	now := time.Now()
	if activity == nil || activity.TimeStarted.IsZero() || activity.TimeEnd.IsZero() {
		return dao.ACTIVITY_DATA_UNDEFINE
	} else if now.Before(activity.TimeStarted) {
		return dao.ACTIVITY_DATE_BEFORE
	} else if now.After(activity.TimeEnd) {
		return dao.ACTIVITY_DATE_AFTER
	}
	history, _ := dao.GetPrizeHistoryServiceR().GetByActivityWuId(reply.ActivityId, wxUser.Id)
	if activity.ActivityType == dao.ACTIVITY_TYPE_LUCK_CHECKIN { //签到抽奖，验证签到条件是否满足
		if history.Prize != "" {
			return strings.NewReplacer("%prize%", history.Prize, "%code%", history.Code).Replace(reply.Success)
		}
		checkin, err := dao.GetCheckinServiceR().GetCheckinInfoByActivityIdAndWuid(activity.RelativeId, wxUser.Id)
		if err != nil {
			return reply.Fail
		}
		if need, _ := strconv.ParseInt(activity.Extra, 10, 64); checkin.Total < need {
			return reply.Fail
		}
	} else if activity.ActivityType == dao.ACTIVITY_TYPE_LUCK_EVERYDAY { //每日抽奖，验证今日是否已经获取
		if history.CreatedAt.Format("2006-01-02") == time.Now().Format("2006-01-02") {
			return strings.NewReplacer("%prize%", history.Prize, "%code%", history.Code).Replace(reply.Success)
		}
	} else {
		if history.Prize != "" {
			return strings.NewReplacer("%prize%", history.Prize, "%code%", history.Code).Replace(reply.Success)
		}
	}

	luck, err := dao.GetLotteryServiceW().Luck(reply.Wid, reply.ActivityId)
	if err == common.ErrLuckFinal {
		return common.ErrLuckFinal.Msg
	} else if err == common.ErrDataUnExist {
		return reply.NoPrizeReturn
	} else if err != nil {
		return common.ErrLuckFail.Msg
	}
	var prize *dao.PrizeModel
	if luck.FirstCodeId != 0 { //表示存在礼包码
		prize, err = dao.GetPrizeServiceW().ChooseOneUsedPrize(reply.ActivityId, luck.Level, luck.FirstCodeId)
		if err == common.ErrDataUnExist {
			return dao.PLEASE_TRY_AGAIN
		}
	}
	if luck.Name != "" {
		dao.GetPrizeHistoryServiceW().Insert(&dao.PrizeHistoryModel{
			ActivityId: reply.ActivityId,
			Wuid:       wxUser.Id,
			Prize:      luck.Name,
			Code:       prize.Code,
			Level:      luck.Level,
		})
	}
	return strings.NewReplacer("%prize%", luck.Name, "%code%", prize.Code).Replace(reply.Success)
}
