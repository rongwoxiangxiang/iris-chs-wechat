package _interface

import "chs/models"

type LotteryInterfaceR interface {
	List(wid, activityId int64) []*models.LotteryModel
}

type LotteryInterfaceW interface {
	Insert(*models.LotteryModel) (int64, error)
	DeleteById(int64) bool
	Luck(wid, activityId int64) (*models.LotteryModel, error)
}

func GetLotteryServiceR() LotteryInterfaceR {
	return new(models.LotteryModel)
}

func GetLotteryServiceW() LotteryInterfaceW {
	return new(models.LotteryModel)
}
