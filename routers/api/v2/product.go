package v2

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"wxApp/models"
	"wxApp/pkg/app"
	"wxApp/pkg/e"
	"wxApp/pkg/util"
)

type Product struct {
	ID				string		`json:"id"`
	Goodstype		string		`json:"goodstype"`
	Name			string		`json:"name" `
	ShortName		string		`json:"short_name" `
	PicUrls			[]string	`json:"pic_urls" `
	Banner			[]string	`json:"banner"`
	Sold			int			`json:"sold"`
	Price			int			`json:"price"`
	MinGroupPrice	int			`json:"min_group_price"`
	LinePrice		int			`json:"line_price"`
	CreateTime		time.Time	`json:"create_time"`
	UpdateTime		time.Time	`json:"update_time"`
	Type 			string		`json:"type"`
	SimpleInfo 		string		`json:"simple_info"`
}

func CreateProduct(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form Product
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	PicUrlsstring , err := json.Marshal(form.PicUrls)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	binnerstr , err := json.Marshal(form.Banner)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	product := models.Product{
		ID:            util.RandomString(30),
		Goodstype:     form.Goodstype,
		Name:          form.Name,
		ShortName:     form.ShortName,
		PicUrls:       string(PicUrlsstring),
		Banner:        string(binnerstr),
		Sold:          form.Sold,
		Price:         form.Price,
		MinGroupPrice: form.MinGroupPrice,
		LinePrice:     form.LinePrice,
		CreateTime:    time.Time{},
		UpdateTime:    time.Time{},
		Type:          form.Type,
		SimpleInfo:    form.SimpleInfo,
	}

	err = models.CreateData(product)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.PRODUCT_CREATE_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func EditProduct(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form Product
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	picUrlsstring , err := json.Marshal(form.PicUrls)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	binnerstr , err := json.Marshal(form.Banner)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	product := &models.Product{
		ID:            form.ID,
		Goodstype:     form.Goodstype,
		Name:          form.Name,
		ShortName:     form.ShortName,
		PicUrls:       string(picUrlsstring),
		Banner:        string(binnerstr),
		Sold:          form.Sold,
		Price:         form.Price,
		MinGroupPrice: form.MinGroupPrice,
		LinePrice:     form.LinePrice,
		CreateTime:    time.Time{},
		UpdateTime:    time.Time{},
		Type:          form.Type,
		SimpleInfo:    form.SimpleInfo,
	}

	err := models.EditData(product)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.PRODUCT_EDIT_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func DeleteProduct(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form Product
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	product  := &models.Product{
		ID:               form.ID,
	}

	err := models.DeleteData(product)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.PRODUCT_DELETE_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

// 获取商品列表
func GetProduct(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form Address
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	product , err := models.GetProduct(form.UserID)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.PRODUCT_GET_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, product)

}

// 获取商品详情
func GetProductInfo(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form Address
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	product , err := models.GetProductInfo(form.ID)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.PRODUCT_GET_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, product)

}
