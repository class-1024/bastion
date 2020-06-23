package service

import (
	"bastion/pkg/datasource"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/medivhzhan/weapp/v2"
	"log"
	"time"
)

type WeApp struct {
	AppName   string
	AppId     string
	AppSecret string
}

var client *resty.Client

func init() {
	client = resty.New()
}

var baseUrl = "https://api.weixin.qq.com/sns/jscode2session"

func getJsCode2SessionUrl(appId, secret, code string) string {
	return fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		baseUrl, appId, secret, code)
}

type SessionKeyRes struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
}

const WEAPPSK = "weapp_sk_"

func (w WeApp) GetSessionKey(code string) (*SessionKeyRes, error) {
	url := getJsCode2SessionUrl(w.AppId, w.AppSecret, code)
	session := SessionKeyRes{}

	//client.SetDebug(true)
	res, err := client.R().
		Get(url)
	if err != nil {
		return nil, err
	}

	// 微信设置头部application/json无效 自己解析
	err = json.Unmarshal(res.Body(), &session)
	if err != nil {
		log.Println("Unmarshal 失败")
		return nil, err
	}

	fmt.Printf("%s \n", session)
	if session.SessionKey == "" || session.Openid == "" {
		return nil, errors.New(res.String())
	}

	// 缓存SessionKey
	datasource.Redis.Set(WEAPPSK+session.Openid, session.SessionKey, 7*24*time.Hour)

	return &session, nil
}

func (w WeApp) DecryptUserInfo(Openid, rawData, encryptedData, signature, iv string) (*weapp.UserInfo, error) {
	sk, err := datasource.Redis.Get(WEAPPSK + Openid).Result()
	if err != nil {
		return nil, fmt.Errorf("redis没有SessionKey %w", err)
	}

	res, err := weapp.DecryptUserInfo(sk, rawData, encryptedData, signature, iv)
	if err != nil {
		return nil, err
	}

	fmt.Printf("DecryptUserInfo: %v \n", res)
	return res, nil
}
