package dao

type PrizeInterfaceR interface {
	LimitUnderActivityList(activityId int64, index int, limit int) []*PrizeModel
}

type PrizeInterfaceW interface {
	ChooseOneUsedPrize(activityId int64, level int8, idGt int64) (*PrizeModel, error)
	Insert(*PrizeModel) (int64, error)
	InsertBatch([]*PrizeModel) (int64, error)
	DeleteById(int64) bool
}

func GetPrizeServiceR() PrizeInterfaceR {
	return new(PrizeModel)
}

func GetPrizeServiceW() PrizeInterfaceW {
	return new(PrizeModel)
}
