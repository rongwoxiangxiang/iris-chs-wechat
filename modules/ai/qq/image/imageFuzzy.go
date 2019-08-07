package image

import (
	"chs/modules/ai/qq"
	"chs/modules/ai/qq/util"
	"log"
)

const imageFuzzyRequestUrl = "https://api.ai.qq.com/fcgi-bin/nlp/nlp_textchat"

type ImageFuzzy struct {
	ImageStruct *Image
}

var imageFuzzy *ImageFuzzy

func NewImageFuzzy(conf ...*qq.Configuration) *ImageFuzzy {
	if imageFuzzy == nil {
		if conf == nil {
			conf[0] = qq.DefaultConfiguration()
		}
		imageFuzzy = new(ImageFuzzy)
		imageFuzzy.ImageStruct = &Image{
			config: conf[0],
		}
	}
	return imageFuzzy
}

func (this *ImageFuzzy) ToMap() map[string]string {
	if this.ImageStruct.image == "" {
		log.Println("ImageFuzzy image err")
		return nil
	}
	config := this.ImageStruct.config.ToMap()
	config["image"] = this.ImageStruct.image
	return config
}

func (this *ImageFuzzy) Process(image string) *ImageFuzzy {
	this.ImageStruct.image = image
	return this
}

func (this *ImageFuzzy) Image() (float64, bool) {
	answer := new(Answer)
	requestBody := qq.GetRequestBody(this)
	err := util.HttpPostJSON(imageFuzzyRequestUrl, requestBody, answer)
	if err != nil || answer.ErrCode != 0 {
		log.Printf("QQ Ai ImageFuzzy request err :%v answer {%v}", err, answer)
		return 0, true
	}
	//TO FIX
	mp := answer.DataJson.(map[string]interface{})
	return mp["confidence"].(float64), mp["fuzzy"].(bool)
}
