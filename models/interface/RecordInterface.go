package _interface

import "chs/models"

type RecordInterfaceR interface {
	GetById(int64) *models.RecordModel
	LimitUnderWidList(wid int64, index int, limit int) []*models.RecordModel
	LimitUnderWuidList(wuid int64, index int, limit int) []*models.RecordModel
}

type RecordInterfaceW interface {
	Insert(*models.RecordModel) (int64, error)
}

func GetRecordServiceR() RecordInterfaceR {
	return new(models.RecordModel)
}

func GetRecordServiceW() RecordInterfaceW {
	return new(models.RecordModel)
}
