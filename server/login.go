package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"wxApp/pkg/setting"
)

// 登陆返回消息体
type WXLoginResp struct {
	OpenId 		string		`json:"openid"`
	SessionKey 	string		`json:"session_key"`
	UnionId 	string		`json:"unionid"`
	Err
}

// 这个函数以 code 作为输入, 返回调用微信接口得到的对象指针和异常情况
func WXLogin(code string) (*WXLoginResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, setting.WxAppSetting.AppID, setting.WxAppSetting.AppSecret, code)

	// 创建http get请求
	resp,err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	log.Println(decoder)

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}

	return &wxResp, nil
}



