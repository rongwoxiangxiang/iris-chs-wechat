package dao

type WechatInterfaceR interface {
	GetById(int64) *WechatModel
	FindByGid(int64) []*WechatModel
	GetByAppid(string) *WechatModel
	GetByFlag(string) *WechatModel
}

type WechatInterfaceW interface {
	Insert(*WechatModel) (int64, error)
	DeleteById(int64) bool
}

func GetWechatServiceR() WechatInterfaceR {
	return new(WechatModel)
}

func GetWechatServiceW() WechatInterfaceW {
	return new(WechatModel)
}
