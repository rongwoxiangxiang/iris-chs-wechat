package dao

type CheckinInterfaceR interface {
	ListByWid(wid int64) []*CheckinModel
	GetCheckinInfoByActivityIdAndWuid(activityId, wuid int64) (*CheckinModel, error)
	GetCheckinByActivityWuid(activityId, wuid int64) (*CheckinModel, error)
}

type CheckinInterfaceW interface {
	Insert(checkin *CheckinModel) (int64, error)
	Update(checkin *CheckinModel) (int64, error)
	DeleteById(id int64) bool
}

func GetCheckinServiceR() CheckinInterfaceR {
	return new(CheckinModel)
}

func GetCheckinServiceW() CheckinInterfaceW {
	return new(CheckinModel)
}
