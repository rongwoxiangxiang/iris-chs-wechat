package _interface

import "chs/models"

type WechatInterfaceR interface {
	GetById(int64) *models.WechatModel
	FindByGid(int64) []*models.WechatModel
	GetByAppid(string) *models.WechatModel
	GetByFlag(string) *models.WechatModel
}

type WechatInterfaceW interface {
	Insert(*models.WechatModel) (int64, error)
	DeleteById(int64) bool
}

func GetWechatServiceR() WechatInterfaceR {
	return new(models.WechatModel)
}

func GetWechatServiceW() WechatInterfaceW {
	return new(models.WechatModel)
}
