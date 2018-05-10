package main

import (
	"github.com/astaxie/beego/httplib"
	"crypto/md5"
	"encoding/hex"
	"time"
	"strconv"
)

const appid = "2015063000000001"
const key = "12345678"

func langDetect(src string) string {
	var json map[string]interface{}
	req := httplib.Post("http://fanyi.baidu.com/langdetect")
	req.Param("query", src).ToJSON(&json)
	return json["lan"].(string)
}

func translate(query, from, to string) string {
	req := httplib.Get("http://api.fanyi.baidu.com/api/trans/vip/translate")
	req.Param("q", query)
	req.Param("from", from)
	req.Param("to", to)
	req.Param("appid", appid)
	salt, sign := getSign(query);
	req.Param("salt", salt)
	req.Param("sign", sign)
	var json map[string]interface{}
	req.ToJSON(&json)
	return json["trans_result"].([]interface{})[0].(map[string]interface{})["dst"].(string)
}

func getSign(query string) (string, string) {
	salt := strconv.FormatInt(time.Now().Unix(), 10)
	sign := md5Hex(appid + query + salt + key)
	return salt, sign
}

func md5Hex(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
