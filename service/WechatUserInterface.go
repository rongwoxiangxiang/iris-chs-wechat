package _interface

import "chs/models"

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
	return new(models.WechatUserModel)
}

func GetWechatUserServiceW() WechatUserInterfaceW {
	return new(models.WechatUserModel)
}
