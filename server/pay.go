package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mozillazg/request"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
	"net/http"
	"wxApp/pkg/setting"
)

type Pay struct {
	PrepayID	string	`json:"prepay_id"`
	Err
}
type PayInfo struct {
	Mchid			string	`json:"mchid"`			// 商户号
	Description		string	`json:"description"`	// 描述
	OutTradeNo		string	`json:"out_trade_no"`	// 订单号
	TransactionId	string	`json:"transaction_id"`	// 微信支付订单号
	TimeExpire		string	`json:"time_expire"`	// 交易结束时间
	Attach			string	`json:"attach"`			// 附加数据
	NotifyUrl		string	`json:"notify_url"`		// 通知地址
	GoodsTag		string	`json:"goods_tag"`		// 订单优惠标记 有值则优惠，无则不优惠
	Amount			int64	`json:"amount"`			// 订单金额
	Payer			string	`json:"payer"`			// openid
	Scene       	string	`json:"scene"`			// 用户终端设备号或者ip
	StoreInfo		string	`json:"store_info"`		// 商户门店id

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
func XiaDan(wxXiaDan PayInfo) (*Pay , error) {

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
func QueryOrder(data PayInfo) (*OrderInfo , error) {

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
func CloseOrder(data PayInfo) (error) {
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

// 小程序支付
func WxPay(data PayInfo) (*jsapi.PrepayWithRequestPaymentResponse , error){
	// 仅做测试用
	// 后期需要读取文件
	var (
		mchID                      string = "190000****"                                // 商户号
		mchCertificateSerialNumber string = "3775B6A45ACD588826D15E583A95F5DD********"  // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"          // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Fatal("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}

	// 发送请求，以下载微信支付平台证书为例
	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	svc := certificates.CertificatesApiService{Client: client}
	resp, result, err := svc.DownloadCertificates(ctx)
	log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)

	if err == nil {
		log.Println(resp)
	} else {
		log.Println(err)
		return nil ,err
	}

	svcc := jsapi.JsapiApiService{Client: client}

	resp1, result1, err := svcc.PrepayWithRequestPayment(ctx,
		jsapi.PrepayRequest{
			Appid:       &setting.WxAppSetting.AppID,
			Mchid:       &data.Mchid,
			Description: &data.Description,
			OutTradeNo:  &data.OutTradeNo,
			Attach:      &data.Attach,
			NotifyUrl:   &data.NotifyUrl,
			Amount: &jsapi.Amount{
				Total: &data.Amount,
			},
			Payer: &jsapi.Payer{
				Openid: core.String("oUpF8uMuAJO_M2pxb1Q9zNjWeS6o"),
			},
		},
	)

	if err == nil {
		log.Println(resp1)
		log.Println(result1)
	} else {
		log.Println(err)
		return nil , err
	}

	return resp1 , nil

}

