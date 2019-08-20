package qq

import (
	"chs/modules/ai/qq/util"
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type AiServer interface {
	ToMap() map[string]string
}

type Answer struct {
	ErrCode  int      `json:"ret"`
	Msg      string   `json:"msg"`
	DataJson DataJson `json:"data"`
}

type DataJson interface{}

func GetRequestBody(ai AiServer) string {
	params := ai.ToMap()
	app_key, ok := params["app_key"]
	if ok != true {
		return ""
	}
	delete(params, "app_key")
	params["time_stamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	params["nonce_str"] = util.GetRandomString(16)

	query := url.Values{}
	for key, value := range params {
		query.Add(key, value)
	}
	buf := make([]byte, 0, 120)
	buf = append(buf, (query.Encode())...)
	requestBody := string(buf)

	buf = append(buf, "&app_key="...)
	buf = append(buf, app_key...)
	hashsum := md5.Sum(buf)
	requestBody += "&sign=" + strings.ToUpper(hex.EncodeToString(hashsum[:]))
	return requestBody
}
