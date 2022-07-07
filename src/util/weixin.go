package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"

	"blog-1.0/config"
	. "blog-1.0/util/log"
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
)

// WeixinSession 微信Session
type WeixinSession struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errMsg"`
}

// WeixinToken 微信access_token
type WeixinToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"` // wtf
}

type WeixinMsgSecCheck struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type QQMsgSecCheck struct {
	ErrCode int    `json:"errCode"` // 和微信的 key 有大小写区分
	ErrMsg  string `json:"errMsg"`
}

const (
	// https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	WeixinSessionURI             = "https://api.weixin.qq.com/sns/jscode2session"
	WeixinGetTokenURI            = "https://api.weixin.qq.com/cgi-bin/token"
	WeixinMsgSecCheckURI         = "https://api.weixin.qq.com/wxa/msg_sec_check"
	WeixinImgSecCheckURI         = "https://api.weixin.qq.com/wxa/img_sec_check"
	WeixinGetWxaCodeUnlimitedURI = "https://api.weixin.qq.com/wxa/getwxacodeunlimit"
)

func unpad(s []byte) []byte {
	size := len(s)
	skip := int(s[size-1])
	return s[:size-skip]
}

// AES-128-CBC
func WXBizDataDeCrypt(str string, iv string, key string) ([]byte, error) {
	// 都要b64解密
	dstr, _ := base64.StdEncoding.DecodeString(str)
	div, _ := base64.StdEncoding.DecodeString(iv)
	dkey, _ := base64.StdEncoding.DecodeString(key)

	block, err := aes.NewCipher(dkey)
	if err != nil {
		return nil, err
	}

	decrypter := cipher.NewCBCDecrypter(block, div)

	decrypter.CryptBlocks(dstr, dstr)

	return unpad(dstr), nil
}

// RequestGetJSON GET请求并返回json
func RequestGetJSON(url string, param req.Param, v interface{}) error {
	r, err := req.Get(url, param)
	if err != nil {
		return err
	}

	err = r.ToJSON(v)
	if err != nil {
		return err
	}

	return nil
}

// GetWeixinAccessToken 获得微信access_token
func GetWeixinAccessToken() (WeixinToken, error) {
	ret := WeixinToken{}
	appInfo := config.C.Wechat
	param := req.Param{
		"appid":      appInfo.AppID,
		"secret":     appInfo.AppSecret,
		"grant_type": "client_credential",
	}
	url := WeixinGetTokenURI
	err := RequestGetJSON(url, param, &ret)
	return ret, err
}

// GetWeixinSession 获得微信Session
func GetWeixinSession(code string) (WeixinSession, error) {
	ret := WeixinSession{}
	param := req.Param{
		"appid":      config.C.Wechat.AppID,
		"secret":     config.C.Wechat.AppSecret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	err := RequestGetJSON(WeixinSessionURI, param, &ret)

	return ret, err
}

// postWeixinMsgSecCheck 敏感内容检查API
func PostWeixinMsgSecCheck(accessToken, content string) (bool, error) {
	ret := WeixinMsgSecCheck{}

	param := req.Param{
		"access_token": accessToken,
	}

	request := map[string]string{
		"content": content,
	}

	r, err := req.Post(WeixinMsgSecCheckURI, param, req.BodyJSON(request))
	if err != nil {
		return false, err
	}

	err = r.ToJSON(&ret)
	if err != nil {
		return false, err
	}
	if ret.ErrCode == 87014 {
		return false, nil
	}
	if ret.ErrCode != 0 {
		return false, errors.New("PostWeixinMsgSecCheck failed: " + strconv.Itoa(ret.ErrCode) + " " + ret.ErrMsg)
	}
	return true, nil
}

// postWeixinImgSecCheck 敏感图片检查API
// 不知道为什么使用req框架失败
func PostWeixinImgSecCheck(accessToken string, file multipart.File, fileName string) (bool, error) {
	ret := WeixinMsgSecCheck{}

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	fw, _ := w.CreateFormFile("media", fileName) // 这里的uploadFile必须和服务器端的FormFile-name一致
	io.Copy(fw, file)
	w.Close()

	resp, err := http.Post(fmt.Sprintf("%s?access_token=%s", WeixinImgSecCheckURI, accessToken),
		w.FormDataContentType(), buf)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll Error", err)
	}
	err = jsoniter.Unmarshal(body, &ret)
	if err != nil {
		return false, err
	}
	if ret.ErrCode == 87014 {
		return false, nil
	}
	if ret.ErrCode != 0 {
		return false, errors.New("PostWeixinImgSecCheck failed: " + strconv.Itoa(ret.ErrCode) + " " + ret.ErrMsg)
	}
	return true, nil
}

func RefreshWeixinAccessToken() {
	t, err := GetWeixinAccessToken()
	if err != nil {
		Logger.Errorf("Error on updating access_token from request: %s", err.Error())
		return
	}

	if t.ErrCode != 0 {
		Logger.Errorf("Error on updating access_token from weixin: %s", t.ErrMsg)
		return
	}

	//err = model.SetWeixinAccessToken(t.AccessToken)
	//if err != nil {
	//	Logger.Errorf("Error on updating access_token from redis: %s", err.Error())
	//	return
	//}

	Logger.Printf("access_token has been refreshed")
}

// GetWxaCodeUnlimited 获得小程序码
func GetWxaCodeUnlimited(accessToken, page, scene string) ([]byte, error) {
	body := map[string]interface{}{
		"scene": scene,
		"page":  page,
	}

	url := fmt.Sprintf("%s?access_token=%s", WeixinGetWxaCodeUnlimitedURI, accessToken)

	r, err := req.Post(url, req.BodyJSON(&body))
	if err != nil {
		return nil, err
	}

	if r.Response().Header.Get("Content-Type") != "image/jpeg" {
		var errResult map[string]interface{}
		err = r.ToJSON(&errResult)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(errResult["errmsg"].(string))
	}

	return ioutil.ReadAll(r.Response().Body)
}
