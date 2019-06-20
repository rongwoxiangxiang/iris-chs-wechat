package _interface

import "chs/models"

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
	return new(models.CheckinModel)
}

func GetCheckinServiceW() CheckinInterfaceW {
	return new(models.CheckinModel)
}
