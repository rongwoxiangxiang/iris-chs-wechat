package ptu

import (
	"chs/config"
	"chs/modules/ai/qq"
	"log"
	"strconv"
)

const faceDecorationRequestUrl = "https://api.ai.qq.com/fcgi-bin/ptu/ptu_faceage"

type FaceDecoration struct {
	*Ptu
	decoration string
}

//type DataJson struct {
//	Image    string
//}

var faceDecoration *FaceDecoration

func NewFaceDecoration(conf ...*qq.Configuration) *FaceDecoration {
	if faceDecoration == nil {
		if conf == nil {
			conf[0] = qq.DefaultConfiguration()
		}
		faceDecoration = new(FaceDecoration)
		faceDecoration.Ptu = &Ptu{
			config: conf[0],
		}
	}
	return faceDecoration
}

func (this *FaceDecoration) ToMap() map[string]string {
	if this.image == "" || this.decoration == "" {
		config.Logger().Error("FaceDecoration image or decoration err")
		return nil
	}
	config := this.config.ToMap()
	config["image"] = this.image
	config["decoration"] = this.decoration
	return config
}

func (this *FaceDecoration) Process(decoration int, image string) *FaceDecoration {
	this.image = image
	this.decoration = strconv.Itoa(decoration)
	return this
}

func (this *FaceDecoration) Image() string {
	return getProcessedImg(faceDecorationRequestUrl, qq.GetRequestBody(this))
}
