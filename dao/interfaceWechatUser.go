package dao

type WechatUserInterfaceR interface {
	GetById(id int64) (*WechatUserModel, error)
	GetByWidAndOpenid(wid int64, openid string) (*WechatUserModel, error)
	LimitUnderWidList(index int, limit int) ([]*WechatUserModel, error)
}

type WechatUserInterfaceW interface {
	Insert(*WechatUserModel) (int64, error)
	Update(*WechatUserModel) (int64, error)
	DeleteById(int64) bool
}

func GetWechatUserServiceR() WechatUserInterfaceR {
	return new(WechatUserModel)
}

func GetWechatUserServiceW() WechatUserInterfaceW {
	return new(WechatUserModel)
}
