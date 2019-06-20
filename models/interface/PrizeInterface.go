package _interface

import "chs/models"

type PrizeInterfaceR interface {
	LimitUnderActivityList(activityId int64, index int, limit int) []*models.PrizeModel
}

type PrizeInterfaceW interface {
	ChooseOneUsedPrize(activityId int64, level string, idGt int64) (*models.PrizeModel, error)
	Insert(*models.PrizeModel) (int64, error)
	InsertBatch([]*models.PrizeModel) (int64, error)
	DeleteById(int64) bool
}

func GetPrizeServiceR() PrizeInterfaceR {
	return new(models.PrizeModel)
}

func GetPrizeServiceW() PrizeInterfaceW {
	return new(models.PrizeModel)
}
