package main

import (
	"chs/util"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestWxReply(t *testing.T) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := util.GetRandomString(8)
	openid := "ol0Xm1aADAb4H-s30_RPjmcmU96g"
	signature := sign("chens", timestamp, nonce)

	requestUrl := fmt.Sprintf("http://127.0.0.1:8888/service/chensss?signature=%s&timestamp=%s&nonce=%s&openid=%s", signature, timestamp, nonce, openid)
	requestStr := "<xml><ToUserName><![CDATA[gh_008302fc091b]]></ToUserName>" +
		"<FromUserName><![CDATA[ol0Xm1aADAb4H-s30_RPjmcmU96g]]></FromUserName>" +
		"<CreateTime>1560924810</CreateTime>" +
		"<MsgType><![CDATA[text]]></MsgType>" +
		"<Content><![CDATA[QQ]]></Content>" +
		"<MsgId>22347068330549091</MsgId>" +
		"</xml>"
	http.NewRequest("POST", requestUrl, strings.NewReader(requestStr))
}

func sign(token, timestamp, nonce string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce}
	strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce))
	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}
