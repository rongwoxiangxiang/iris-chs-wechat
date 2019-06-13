package _interface

import "chs/models"
import "chs/config"

type WechatUserInterfaceR interface {
	GetById(id int64) (*models.WechatUserModel, error)
	GetByWidAndOpenid(wid int64, openid string) (*models.WechatUserModel, error)
	LimitUnderWidList(index int, limit int) ([]*models.WechatUserModel, error)
}

type WechatUserInterfaceW interface {
	Insert(*models.WechatUserModel) (int64, error)
	Update(*models.WechatUserModel) (int64, error)
	DeleteById(int64) bool
}

func GetWechatUserServiceR() WechatUserInterfaceR {
	models.Db = config.GetDbR(models.APP_DB_READ)
	return new(models.WechatUserModel)
}

func GetWechatUserServiceW() WechatUserInterfaceW {
	models.Db = config.GetDbW(models.APP_DB_WRITE)
	return new(models.WechatUserModel)
}
