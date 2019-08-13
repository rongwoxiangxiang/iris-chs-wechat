package dao

type RecordInterfaceR interface {
	GetById(int64) *RecordModel
	LimitUnderWidList(wid int64, index int, limit int) []*RecordModel
	LimitUnderWuidList(wuid int64, index int, limit int) []*RecordModel
}

type RecordInterfaceW interface {
	Insert(*RecordModel) (int64, error)
}

func GetRecordServiceR() RecordInterfaceR {
	return new(RecordModel)
}

func GetRecordServiceW() RecordInterfaceW {
	return new(RecordModel)
}
