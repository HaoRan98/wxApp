package v2

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"wxApp/models"
	"wxApp/pkg/app"
	"wxApp/pkg/e"
	"wxApp/pkg/util"
)

type Address struct {
	ID 					string		`json:"id"`
	UserID				string		`json:"user_id"`
	ReceiverName		string		`json:"receiver_name" gorm:" comment '收件人姓名';" `
	ReceiverPhone		string		`json:"receiver_phone" gorm:" comment '收件人固定电话';"`
	ReceiverMobile		string		`json:"receiver_mobile" gorm:" comment '收件人移动电话';"`
	ReceiverProvince	string 		`json:"receiver_province" gorm:" comment '省份';"`
	ReceiverCity		string		`json:"receiver_city" gorm:" comment '城市';"`
	ReceiverDistrict	string		`json:"receiver_district" gorm:" comment '县区';"`
	ReceicerAddress		string		`json:"receicer_address" gorm:" comment '详细地址';"`
	ReceiverZip			string		`json:"receiver_zip" gorm:" comment '邮编';"`
}


func CreateAddress(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form Address
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	address  := &models.Address{
		ID:               util.RandomString(30),
		UserID:           form.UserID,
		ReceiverName:     form.ReceiverName,
		ReceiverPhone:    form.ReceiverPhone,
		ReceiverMobile:   form.ReceiverMobile,
		ReceiverProvince: form.ReceiverProvince,
		ReceiverCity:     form.ReceiverCity,
		ReceiverDistrict: form.ReceiverDistrict,
		ReceicerAddress:  form.ReceicerAddress,
		ReceiverZip:      form.ReceiverZip,
		CreateTime:       time.Time{},
		UpdateTime:       time.Time{},
	}

	err := models.CreateData(address)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ADDRESS_CREATE_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func EditAddress(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form Address
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	address  := &models.Address{
		ID:               form.ID,
		ReceiverName:     form.ReceiverName,
		ReceiverPhone:    form.ReceiverPhone,
		ReceiverMobile:   form.ReceiverMobile,
		ReceiverProvince: form.ReceiverProvince,
		ReceiverCity:     form.ReceiverCity,
		ReceiverDistrict: form.ReceiverDistrict,
		ReceicerAddress:  form.ReceicerAddress,
		ReceiverZip:      form.ReceiverZip,
		UpdateTime:       time.Time{},
	}

	err := models.EditData(address)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ADDRESS_EDIT_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func DeleteAddress(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form Address
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	address  := &models.Address{
		ID:               form.ID,
	}

	err := models.DeleteData(address)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ADDRESS_DELETE_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

// 获取地址列表
func GetAddress(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form Address
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	address , err := models.GetAddress(form.UserID)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ADDRESS_GET_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, address)

}

// 获取地址详情
func GetAddressInfo(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form Address
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	address , err := models.GetAddressInfo(form.ID)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ADDRESS_GET_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, address)

}
