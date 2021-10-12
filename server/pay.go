package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mozillazg/request"
	"log"
	"net/http"
	"wxApp/pkg/setting"
	v1 "wxApp/routers/api/v1"
)

type Pay struct {
	PrepayID	string	`json:"prepay_id"`
	Err
}

type OrderInfo struct {
	Amount 	struct{
		Currency		string	`json:"currency"`
		PayerCurrency	string	`json:"payer_currency"`
		PayerTotal		int		`json:"payer_total"`
		Total			int		`json:"total"`
	}
	Appid	string	`json:"appid"`
	Attach	string	`json:"attach"`
	BankType	string	`json:"bank_type"`
	Mchid		string	`json:"mchid"`
	OutTradeNo	string	`json:"out_trade_no"`
	Payer	struct{
		Openid 	string	`json:"openid"`
	}
	PromotionDetail	interface{}	`json:"promotion_detail"`
	SuccessTime	string	`json:"success_time"`
	TradeState	string	`json:"trade_state"`
	TradeStateDesc	string	`json:"trade_state_desc"`
	Tradetype	string	`json:"tradetype"`
	TransactionId	string	`json:"transaction_id"`
	Err
}

// 下单
func XiaDan(wxXiaDan v1.PayInfo) (*Pay , error) {

	url := "https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi"

	c := new(http.Client)
	req := request.NewRequest(c)
	req.Json = map[string]interface{}{
		"appid":setting.WxAppSetting.AppID,
		"mchid":wxXiaDan.Mchid,
		"description":wxXiaDan.Description,
		"out_trade_no":wxXiaDan.OutTradeNo,
		"time_expire":wxXiaDan.TimeExpire,
		"notify_url":wxXiaDan.NotifyUrl,
		"goods_tag":wxXiaDan.GoodsTag,
		"amount":map[string]interface{}{
			"total":wxXiaDan.Amount,
		},
		"payer":map[string]interface{}{
			"openid":wxXiaDan.Payer,
		},
		"scene":map[string]interface{}{
			"device_id":wxXiaDan.Payer,
			"store_info":map[string]interface{}{
				"id":wxXiaDan.StoreInfo,
			},
		},
	}

	resp,err := req.Post(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	wxResp := Pay{}
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

// 查询订单
func QueryOrder(data v1.PayInfo) (*OrderInfo , error) {

	var url string
	c := new(http.Client)
	req := request.NewRequest(c)

	if data.TransactionId != "" {
		url = "https://api.mch.weixin.qq.com/v3/pay/transactions/id/%s?mchid=%s"
	}else if data.OutTradeNo != "" {
		url = "https://api.mch.weixin.qq.com/v3/pay/transactions/out-trade-no/%s?mchid=%s"
	}else {
		return nil , errors.New("无订单号")
	}
	resp,err := req.Post(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	wxResp := OrderInfo{}
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

// 关闭订单
// 1、商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；
// 2、系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
func CloseOrder(data v1.PayInfo) (error) {
	url := "https://api.mch.weixin.qq.com/v3/pay/transactions/out-trade-no/%s/close"
	url = fmt.Sprintf(url,data.OutTradeNo)

	c := new(http.Client)
	req := request.NewRequest(c)
	req.Json = map[string]interface{}{
		"mchid":data.Mchid,
	}

	_,err := req.Post(url)
	if err != nil {
		return err
	}

	return nil

}

//

