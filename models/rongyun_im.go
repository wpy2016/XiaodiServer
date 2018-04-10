package models

import (
	"github.com/astaxie/beego/httplib"
	"strconv"
	"time"
	"crypto/sha1"
	"io"
	"fmt"
	"encoding/json"
	"strings"
	"errors"
	"math/rand"
)

type RongYunTokenResp struct {
	Code   int    `json:"code"`
	UserId string `json:"userId"`
	Token  string `json:"token"`
}

type RongyunUser struct {
	AppKey    string
	AppSecret string
}

var rongyunUser = RongyunUser{AppKey: "x18ywvqfxbczc", AppSecret: "NLo09CCOpBee0r"}

func GetRongyunToken(userId, name, portraitUri string) (string, error) {
	if ( userId == "") {
		return "", errors.New("Paramer 'userId' is required");
	}

	if ( name == "") {
		return "", errors.New("Paramer 'name' is required");
	}

	if ( portraitUri == "") {
		return "", errors.New("Paramer 'portraitUri' is required");
	}

	destinationUrl := "http://api.cn.ronghub.com/user/getToken.json"
	req := httplib.Post(destinationUrl)
	fillHeader(req, rongyunUser.AppKey, rongyunUser.AppSecret)
	req.Param("userId", userId)
	req.Param("name", name)
	req.Param("portraitUri", portraitUri)
	byteData, err := req.Bytes()
	if err != nil {
		return "", err
	}
	strData := string(byteData)
	var ret = RongYunTokenResp{}
	err = JsonParse(strData, &ret)
	if nil != err {
		return "", err
	}
	return ret.Token, nil

}

//本地生成签名
//Signature (数据签名)计算方法：将系统分配的 App Secret、Nonce (随机数)、
//Timestamp (时间戳)三个字符串按先后顺序拼接成一个字符串并进行 SHA1 哈希计算。如果调用的数据签名验证失败，接口调用会返回 HTTP 状态码 401。
func getSignature(appSecret string) (nonce, timestamp, signature string) {
	nonceInt := rand.Int()
	nonce = strconv.Itoa(nonceInt)
	timeInt64 := time.Now().Unix()
	timestamp = strconv.FormatInt(timeInt64, 10)
	h := sha1.New()
	io.WriteString(h, appSecret+nonce+timestamp)
	signature = fmt.Sprintf("%x", h.Sum(nil))
	return
}

//API签名
func fillHeader(req *httplib.BeegoHTTPRequest, appKey, appSecret string) {
	nonce, timestamp, signature := getSignature(appSecret)
	req.Header("App-Key", appKey)
	req.Header("Nonce", nonce)
	req.Header("Timestamp", timestamp)
	req.Header("Signature", signature)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
}

func JsonParse(jsonStr string, v interface{}) error {
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	err := dec.Decode(v)
	if err != nil {
		return err
	}
	return nil
}
