package dao

type ReplyInterfaceR interface {
	GetById(int64) *ReplyModel
	FindOne(*ReplyModel) *ReplyModel
	LimitUnderWidList(wid int64, index int, limit int) []*ReplyModel
}

type ReplyInterfaceW interface {
	Insert(*ReplyModel) (int64, error)
	ChangeDisabledByWidActivityId(wid, activityId int64, disabled int8) bool
	Update(*ReplyModel) (int64, error)
	DeleteById(int64) bool
}

func GetReplyServiceR() ReplyInterfaceR {
	return new(ReplyModel)
}

func GetReplyServiceW() ReplyInterfaceW {
	return new(ReplyModel)
}
