package _interface

import "chs/models"
import "chs/config"

type ReplyInterfaceR interface {
	GetById(int64) *models.ReplyModel
	FindOne(*models.ReplyModel) *models.ReplyModel
	LimitUnderWidList(wid int64, index int, limit int) []*models.ReplyModel
}

type ReplyInterfaceW interface {
	Insert(*models.ReplyModel) (int64, error)
	ChangeDisabledByWidActivityId(wid, activityId int64, disabled int8) bool
	Update(*models.ReplyModel) (int64, error)
	DeleteById(int64) bool
}

func GetReplyServiceR() ReplyInterfaceR {
	models.Db = config.GetDbR(models.APP_DB_READ)
	return new(models.ReplyModel)
}

func GetReplyServiceW() ReplyInterfaceW {
	models.Db = config.GetDbW(models.APP_DB_WRITE)
	return new(models.ReplyModel)
}
