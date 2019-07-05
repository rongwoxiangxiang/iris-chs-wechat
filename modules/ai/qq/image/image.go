package image

import (
	"chs/modules/ai/qq"
)

type Image struct {
	config *qq.Configuration
	image  string //原始图片的base64编码数据（原图大小上限500KB）
}

type Answer struct {
	ErrCode  int      `json:"ret"`
	Msg      string   `json:"msg"`
	DataJson DataJson `json:"data"`
}

type DataJson interface{}
