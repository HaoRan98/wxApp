package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wxApp/pkg/app"
	"wxApp/pkg/e"
	"wxApp/server"
)



// 下单
func WxAppXiaDan(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form server.PayInfo
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
		form server.PayInfo
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
		form server.PayInfo
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

func WxPay(c *gin.Context)  {
	var (
		appG = app.Gin{C: c}
		form server.PayInfo
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	resp , err := server.WxPay(form)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.WXPAY_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, resp)

}



