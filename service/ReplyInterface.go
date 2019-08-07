package _interface

import "chs/models"

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
	return new(models.ReplyModel)
}

func GetReplyServiceW() ReplyInterfaceW {
	return new(models.ReplyModel)
}
