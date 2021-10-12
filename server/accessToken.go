package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"wxApp/pkg/setting"
	"wxApp/pkg/waitgroup"
)

//  接口调用凭证
var AccessToken string
var err error

// 登陆接口凭证消息体
type WXTokenResp struct {
	AccessToken	string	`json:"access_token"`
	ExpiresIn 	int		`json:"expires_in"`
	ErrCode 	int		`json:"errcode"`
	ErrMsg 		string 	`json:"errmsg"`
}

func GetAccessToken() (string,error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	url = fmt.Sprintf(url , setting.WxAppSetting.AppID,setting.WxAppSetting.AppSecret)

	// 创建http get请求
	resp,err := http.Get(url)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXTokenResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return "", err
	}

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return "", errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}else {
		return wxResp.AccessToken , nil
	}
}

// 定时任务
func TimeAccessToken() {

	waitgroup.WG.Add(1)
	go func() {
		defer waitgroup.WG.Done()
		ticker := time.NewTicker(time.Hour * 1)
		for range ticker.C {
			AccessToken , err = GetAccessToken()
			if err != nil {
				log.Printf("accessToken获取失败：",err)
			}
			log.Println(AccessToken)
		}
	}()
}
