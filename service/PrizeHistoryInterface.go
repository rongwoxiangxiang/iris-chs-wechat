package _interface

import "chs/models"

type PrizeHistoryInterfaceR interface {
	GetByActivityWuId(activityId, wuid int64) (*models.PrizeHistoryModel, error)
	LimitUnderActivityList(activityId int64, index int, limit int) []*models.PrizeHistoryModel
}

type PrizeHistoryInterfaceW interface {
	Insert(*models.PrizeHistoryModel) (int64, error)
	DeleteById(int64) bool
}

func GetPrizeHistoryServiceR() PrizeHistoryInterfaceR {
	return new(models.PrizeHistoryModel)
}

func GetPrizeHistoryServiceW() PrizeHistoryInterfaceW {
	return new(models.PrizeHistoryModel)
}
