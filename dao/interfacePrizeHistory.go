package dao

type PrizeHistoryInterfaceR interface {
	GetByActivityWuId(activityId, wuid int64) (*PrizeHistoryModel, error)
	LimitUnderActivityList(activityId int64, index int, limit int) []*PrizeHistoryModel
}

type PrizeHistoryInterfaceW interface {
	Insert(*PrizeHistoryModel) (int64, error)
	DeleteById(int64) bool
}

func GetPrizeHistoryServiceR() PrizeHistoryInterfaceR {
	return new(PrizeHistoryModel)
}

func GetPrizeHistoryServiceW() PrizeHistoryInterfaceW {
	return new(PrizeHistoryModel)
}
