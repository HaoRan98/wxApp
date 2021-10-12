package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wxApp/pkg/app"
	"wxApp/pkg/e"
	"wxApp/server"
)

type PayInfo struct {
	Mchid			string	`json:"mchid"`			// 商户号
	Description		string	`json:"description"`	// 描述
	OutTradeNo		string	`json:"out_trade_no"`	// 订单号
	TransactionId	string	`json:"transaction_id"`	// 微信支付订单号
	TimeExpire		string	`json:"time_expire"`	// 交易结束时间
	Attach			string	`json:"attach"`			// 附加数据
	NotifyUrl		string	`json:"notify_url"`		// 通知地址
	GoodsTag		string	`json:"goods_tag"`		// 订单优惠标记 有值则优惠，无则不优惠
	Amount			float32	`json:"amount"`			// 订单金额
	Payer			string	`json:"payer"`			// openid
	Scene       	string	`json:"scene"`			// 用户终端设备号或者ip
	StoreInfo		string	`json:"store_info"`		// 商户门店id

}

// 下单
func WxAppXiaDan(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form PayInfo
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	data , err := server.XiaDan(form)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ORDER_CREATE_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)

}

// 查询订单
func WxAppSelectOrder(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form PayInfo
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	data , err := server.QueryOrder(form)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ORDER_GET_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)

}

// 关闭订单
func WxAppCloseOrder(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form PayInfo
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	err := server.CloseOrder(form)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ORDER_EDIT_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

// 小程序支付
func WxPay() {
	wx.requestPayment
}

