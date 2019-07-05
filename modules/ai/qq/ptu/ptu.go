package ptu

import (
	"chs/modules/ai/qq"
	"chs/modules/ai/qq/util"
	"log"
)

type Ptu struct {
	config *qq.Configuration
	image  string //原始图片的base64编码数据（原图大小上限500KB）
}

type Answer struct {
	ErrCode  int      `json:"ret"`
	Msg      string   `json:"msg"`
	DataJson DataJson `json:"data"`
}

type DataJson struct {
	Image string `json:"image"`
}

func getProcessedImg(requestUrl, requestBody string) string {
	answer := new(Answer)
	err := util.HttpPostJSON(requestUrl, requestBody, answer)
	if err != nil || answer.ErrCode != 0 {
		log.Printf("QQ Ai Img request err :%v Image {%v}", err, answer)
		return ""
	}
	return answer.DataJson.Image
}
