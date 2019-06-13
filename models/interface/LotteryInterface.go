package _interface

import "chs/models"
import "chs/config"

type LotteryInterfaceR interface {
	List(wid, activityId int64) []*models.LotteryModel
}

type LotteryInterfaceW interface {
	Insert(*models.LotteryModel) (int64, error)
	DeleteById(int64) bool
	Luck(wid, activityId int64) (*models.LotteryModel, error)
}

func GetLotteryServiceR() LotteryInterfaceR {
	models.Db = config.GetDbR(models.APP_DB_READ)
	return new(models.LotteryModel)
}

func GetLotteryServiceW() LotteryInterfaceW {
	models.Db = config.GetDbW(models.APP_DB_WRITE)
	return new(models.LotteryModel)
}
