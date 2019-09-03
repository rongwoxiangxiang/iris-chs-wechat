package ptu

import (
	"chs/config"
	"chs/modules/ai/qq"
	"log"
)

const faceAgeRequestUrl = "https://api.ai.qq.com/fcgi-bin/ptu/ptu_faceage"

type FaceAge struct {
	*Ptu
}

//type DataJson struct {
//	Image    string
//}

var faceAge *FaceAge

func NewFaceAge(conf ...*qq.Configuration) *FaceAge {
	if faceAge == nil {
		if conf == nil {
			conf[0] = qq.DefaultConfiguration()
		}
		faceAge = new(FaceAge)
		faceAge.Ptu = &Ptu{
			config: conf[0],
		}
	}
	return faceAge
}

func (this *FaceAge) ToMap() map[string]string {
	if this.image == "" {
		config.Logger().Error("FaceAge cosmetic or image err")
		return nil
	}
	config := this.config.ToMap()
	config["image"] = this.image
	return config
}

func (this *FaceAge) Process(image string) *FaceAge {
	this.image = image
	return this
}

func (this *FaceAge) Image() string {
	return getProcessedImg(faceAgeRequestUrl, qq.GetRequestBody(this))
}
