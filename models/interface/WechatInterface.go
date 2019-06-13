package _interface

import "chs/models"
import "chs/config"

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
	models.Db = config.GetDbR(models.APP_DB_READ)
	return new(models.WechatModel)
}

func GetWechatServiceW() WechatInterfaceW {
	models.Db = config.GetDbW(models.APP_DB_WRITE)
	return new(models.WechatModel)
}
