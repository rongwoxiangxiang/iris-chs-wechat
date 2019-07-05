package ptu

import (
	"chs/modules/ai/qq"
	"log"
	"strconv"
)

const faceStickerRequestUrl = "https://api.ai.qq.com/fcgi-bin/ptu/ptu_facecosmetic"

type FaceSticker struct {
	*Ptu
	sticker string
}

var faceSticker *FaceSticker

func NewFaceSticker(conf ...*qq.Configuration) *FaceSticker {
	if faceSticker == nil {
		if conf == nil {
			conf[0] = qq.DefaultConfiguration()
		}
		faceSticker = new(FaceSticker)
		faceSticker.Ptu = &Ptu{
			config: conf[0],
		}
	}
	return faceSticker
}

func (this *FaceSticker) ToMap() map[string]string {
	if this.sticker == "" || this.image == "" {
		log.Println("FaceSticker sticker or image err")
		return nil
	}
	config := this.config.ToMap()
	config["image"] = this.image
	config["sticker"] = this.sticker
	return config
}

func (this *FaceSticker) Process(sticker int, image string) *FaceSticker {
	this.image = image
	this.sticker = strconv.Itoa(sticker)
	return this
}

func (this *FaceSticker) Image() string {
	return getProcessedImg(faceStickerRequestUrl, qq.GetRequestBody(this))
}
