package ptu

import (
	"chs/config"
	"chs/modules/ai/qq"
	"strconv"
)

const imgfilterRequestUrl = "https://api.ai.qq.com/fcgi-bin/ptu/ptu_imgfilter"             //对原图进行滤镜特效处理，更适合人物图片
const visionImgfilterRequestUrl = "https://api.ai.qq.com/fcgi-bin/vision/vision_imgfilter" //更适合风景图片

type Imgfilter struct {
	*Ptu
	filter string
}

var imgfilter *Imgfilter

func NewImgfilter(conf ...*qq.Configuration) *Imgfilter {
	if imgfilter == nil {
		if conf == nil {
			conf[0] = qq.DefaultConfiguration()
		}
		imgfilter = new(Imgfilter)
		imgfilter.Ptu = &Ptu{
			config: conf[0],
		}
	}
	return imgfilter
}

func (this *Imgfilter) ToMap() map[string]string {
	if this.filter == "" || this.image == "" {
		config.Logger().Error("Imgfilter filter or image err")
		return nil
	}
	config := this.config.ToMap()
	config["image"] = this.image
	config["cosmetic"] = this.filter
	return config
}

func (this *Imgfilter) Process(filter int, image string) *Imgfilter {
	this.image = image
	this.filter = strconv.Itoa(filter)
	return this
}

func (this *Imgfilter) Image() string {
	return getProcessedImg(imgfilterRequestUrl, qq.GetRequestBody(this))
}

func (this *Imgfilter) VisionImage() string {
	return getProcessedImg(visionImgfilterRequestUrl, qq.GetRequestBody(this))
}
