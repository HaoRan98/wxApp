package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mozillazg/request"
	"log"
	"net/http"
)

type QRCode struct {
	Bean 		string	`json:"bean"`
	ReturnCode 	int 	`json:"return_code"`
	Err
}

func GetQRCode() (*QRCode,error) {
	url := "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=%s"
	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, AccessToken)

	// 创建http get请求
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Json = map[string]interface{}{
		"access_token ":AccessToken,
		"path":"page/index/index",
	}
	resp,err := req.Post(url)
	if err != nil {
	}
	defer resp.Body.Close()

	wxResp := QRCode{}
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
