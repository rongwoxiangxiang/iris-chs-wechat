package dao

type ActivityInterfaceR interface {
	GetById(int64) *ActivityModel
	LimitUnderWidList(index, limit, wid int) []*ActivityModel
}

type ActivityInterfaceW interface {
	Insert(*ActivityModel) (int64, error)
	Update(*ActivityModel) (int64, error)
	DeleteById(int64) bool
}

func GetActivityServiceR() ActivityInterfaceR {
	return new(ActivityModel)
}

func GetActivityServiceW() ActivityInterfaceW {
	return new(ActivityModel)
}
