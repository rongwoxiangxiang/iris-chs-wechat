package _interface

import "chs/models"
import "chs/config"

type PrizeHistoryInterfaceR interface {
	GetByActivityWuId(activityId, wuid int64) (*models.PrizeHistoryModel, error)
	LimitUnderActivityList(activityId int64, index int, limit int) []*models.PrizeHistoryModel
}

type PrizeHistoryInterfaceW interface {
	Insert(*models.PrizeHistoryModel) (int64, error)
	DeleteById(int64) bool
}

func GetPrizeHistoryServiceR() PrizeHistoryInterfaceR {
	models.Db = config.GetDbR(models.APP_DB_READ)
	return new(models.PrizeHistoryModel)
}

func GetPrizeHistoryServiceW() PrizeHistoryInterfaceW {
	models.Db = config.GetDbW(models.APP_DB_WRITE)
	return new(models.PrizeHistoryModel)
}
