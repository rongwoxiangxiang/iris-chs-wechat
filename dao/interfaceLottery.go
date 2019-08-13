package dao

type LotteryInterfaceR interface {
	List(wid, activityId int64) []*LotteryModel
}

type LotteryInterfaceW interface {
	Insert(*LotteryModel) (int64, error)
	DeleteById(int64) bool
	Luck(wid, activityId int64) (*LotteryModel, error)
}

func GetLotteryServiceR() LotteryInterfaceR {
	return new(LotteryModel)
}

func GetLotteryServiceW() LotteryInterfaceW {
	return new(LotteryModel)
}
