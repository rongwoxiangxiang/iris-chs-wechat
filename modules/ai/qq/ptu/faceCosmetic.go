package ptu

import (
	"chs/modules/ai/qq"
	"log"
	"strconv"
)

const faceCosmeticRequestUrl = "https://api.ai.qq.com/fcgi-bin/ptu/ptu_facecosmetic"

type FaceCosmetic struct {
	*Ptu
	cosmetic string
}

//type DataJson struct {
//	Image    string
//}

var faceCosmetic *FaceCosmetic

func NewFaceCosmetic(conf ...*qq.Configuration) *FaceCosmetic {
	if faceCosmetic == nil {
		if conf == nil {
			conf[0] = qq.DefaultConfiguration()
		}
		faceCosmetic = new(FaceCosmetic)
		faceCosmetic.Ptu = &Ptu{
			config: conf[0],
		}
	}
	return faceCosmetic
}

func (this *FaceCosmetic) ToMap() map[string]string {
	if this.cosmetic == "" || this.image == "" {
		log.Println("FaceCosmetic cosmetic or image err")
		return nil
	}
	config := this.config.ToMap()
	config["image"] = this.image
	config["cosmetic"] = this.cosmetic
	return config
}

func (this *FaceCosmetic) Process(cosmetic int, image string) *FaceCosmetic {
	this.image = image
	this.cosmetic = strconv.Itoa(cosmetic)
	return this
}

func (this *FaceCosmetic) Image() string {
	return getProcessedImg(faceCosmeticRequestUrl, qq.GetRequestBody(this))
}
