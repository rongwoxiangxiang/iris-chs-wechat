package _interface

import "chs/models"
import "chs/config"

type CheckinInterfaceR interface {
	ListByWid(wid int64) []*models.CheckinModel
	GetCheckinInfoByActivityIdAndWuid(activityId, wuid int64) (*models.CheckinModel, error)
	GetCheckinByActivityWuid(activityId, wuid int64) (*models.CheckinModel, error)
}

type CheckinInterfaceW interface {
	Insert(checkin *models.CheckinModel) (int64, error)
	Update(checkin *models.CheckinModel) (int64, error)
	DeleteById(id int64) bool
}

func GetCheckinServiceR() CheckinInterfaceR {
	models.Db = config.GetDbR(models.APP_DB_READ)
	return new(models.CheckinModel)
}

func GetCheckinServiceW() CheckinInterfaceW {
	models.Db = config.GetDbW(models.APP_DB_WRITE)
	return new(models.CheckinModel)
}
