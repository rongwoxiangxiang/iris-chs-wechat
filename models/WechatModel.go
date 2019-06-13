package models

import (
	"time"
)

type WechatModel struct {
	Id             int64 `xorm:"pk"`
	Gid            int64
	Name           string
	Appid          string
	Appsecret      string
	EncodingAesKey string
	Token          string
	Flag           string
	Type           int8
	Pass           int8
	SaveInput      int8
	NeedSaveMen    int8 //该公众号hander是否持久化
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (w *WechatModel) TableName() string {
	return "wechats"
}

func (w *WechatModel) Insert(wechat *WechatModel) (int64, error) {
	return Db.InsertOne(wechat)
}

func (w *WechatModel) DeleteById(id int64) bool {
	if id == 0 {
		return false
	}
	_, err := Db.Id(id).Unscoped().Delete(&WechatModel{})
	if err != nil {
		return false
	}
	return true
}

func (w *WechatModel) GetById(id int64) *WechatModel {
	if id != 0 {
		wechat := new(WechatModel)
		wechat.Id = id
		has, err := Db.Get(wechat)
		if !has || err != nil {
			return nil
		}
		return wechat
	}
	return nil
}

func (w *WechatModel) GetByAppid(appid string) *WechatModel {
	if appid == "" {
		return nil
	}
	wechat := new(WechatModel)
	has, err := Db.Where("appid = ?", appid).Get(wechat)
	if !has || err != nil {
		return nil
	}
	return wechat
}

func (w *WechatModel) GetByFlag(flag string) *WechatModel {
	if flag == "" {
		return nil
	}
	wechat := new(WechatModel)
	has, err := Db.Where("flag = ?", flag).Get(wechat)
	if !has || err != nil {
		return nil
	}
	return wechat
}

func (w *WechatModel) FindByGid(gid int64) (wechats []*WechatModel) {
	if gid == 0 {
		return nil
	}
	err := Db.Where("gid = ?", gid).Find(&wechats)
	if err != nil {
		return nil
	}
	return wechats
}
