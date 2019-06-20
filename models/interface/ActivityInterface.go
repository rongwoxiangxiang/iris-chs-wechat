package _interface

import "chs/models"

type ActivityInterfaceR interface {
	GetById(int64) *models.ActivityModel
	LimitUnderWidList(index, limit, wid int) []*models.ActivityModel
}

type ActivityInterfaceW interface {
	Insert(*models.ActivityModel) (int64, error)
	Update(*models.ActivityModel) (int64, error)
	DeleteById(int64) bool
}

func GetActivityServiceR() ActivityInterfaceR {
	return new(models.ActivityModel)
}

func GetActivityServiceW() ActivityInterfaceW {
	return new(models.ActivityModel)
}
