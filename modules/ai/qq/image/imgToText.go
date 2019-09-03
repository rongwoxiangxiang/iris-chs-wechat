package image

import (
	"chs/config"
	"chs/modules/ai/qq"
	"chs/modules/ai/qq/util"
)

const imageToTextRequestUrl = "https://api.ai.qq.com/fcgi-bin/nlp/nlp_textchat"

type ImageToText struct {
	ImageStruct *Image
	SessionId   string
}

var imageToText *ImageToText

func NewImageToText(conf ...*qq.Configuration) *ImageToText {
	if imageToText == nil {
		if conf == nil {
			conf[0] = qq.DefaultConfiguration()
		}
		imageToText = new(ImageToText)
		imageFuzzy.ImageStruct = &Image{
			config: conf[0],
		}
	}
	return imageToText
}

func (this *ImageToText) ToMap() map[string]string {
	if this.ImageStruct.image == "" || this.SessionId == "" {
		config.Logger().Error("ImageToText image or SessionId empty")
		return nil
	}
	config := this.ImageStruct.config.ToMap()
	config["image"] = this.ImageStruct.image
	config["session_id"] = this.SessionId
	return config
}

func (this *ImageToText) Process(sessionId, image string) *ImageToText {
	this.ImageStruct.image = image
	this.SessionId = sessionId
	return this
}

func (this *ImageToText) Image() string {
	answer := new(Answer)
	requestBody := qq.GetRequestBody(this)
	err := util.HttpPostJSON(imageFuzzyRequestUrl, requestBody, answer)
	if err != nil || answer.ErrCode != 0 {
		config.Logger().Errorf("QQ Ai ImageToText request err :%v answer {%v}", err, answer)
		return ""
	}
	return answer.DataJson.(map[string]interface{})["text"].(string)
}
