package _interface

import "chs/models"
import "chs/config"

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
	models.Db = config.GetDbR(models.APP_DB_READ)
	return new(models.ActivityModel)
}

func GetActivityServiceW() ActivityInterfaceW {
	models.Db = config.GetDbW(models.APP_DB_WRITE)
	return new(models.ActivityModel)
}
