package _interface

import "chs/models"
import "chs/config"

type RecordInterfaceR interface {
	GetById(int64) *models.RecordModel
	LimitUnderWidList(wid int64, index int, limit int) []*models.RecordModel
	LimitUnderWuidList(wuid int64, index int, limit int) []*models.RecordModel
}

type RecordInterfaceW interface {
	Insert(*models.RecordModel) (int64, error)
}

func GetRecordServiceR() RecordInterfaceR {
	models.Db = config.GetDbR(models.APP_DB_READ)
	return new(models.RecordModel)
}

func GetRecordServiceW() RecordInterfaceW {
	models.Db = config.GetDbW(models.APP_DB_WRITE)
	return new(models.RecordModel)
}
