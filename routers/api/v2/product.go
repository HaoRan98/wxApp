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

type Product struct {
	ID			string		`json:"id"`
	CategoryID	string		`json:"category_id" gorm:" comment '分类ID';"`
	Name		string		`json:"name" gorm:" comment '名称';"`
	Subtitle	string		`json:"subtitle" gorm:" comment '副标题';"`
	MainImage	string		`json:"main_image" gorm:" comment '主图';"`
	SubImages	string		`json:"sub_images" gorm:" comment '支付流水号';"`
	Detail		string		`json:"detail"`
	Price		float32		`json:"price"`
	Stock		int			`json:"stock"`
	Status		string		`json:"status"`
	CreateTime	time.Time	`json:"create_time"`
	UpdateTime	time.Time	`json:"update_time"`
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

	product := &models.Product{
		ID:         util.RandomString(30),
		Name:       form.Name,
		Subtitle:   form.Subtitle,
		MainImage:  form.MainImage,
		SubImages:  form.SubImages,
		Detail:     form.Detail,
		Price:      form.Price,
		Stock:      form.Stock,
		Status:     "0",
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	}

	err := models.CreateData(product)
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

	product := &models.Product{
		ID:         util.RandomString(30),
		Name:       form.Name,
		Subtitle:   form.Subtitle,
		MainImage:  form.MainImage,
		SubImages:  form.SubImages,
		Detail:     form.Detail,
		Price:      form.Price,
		Stock:      form.Stock,
		Status:     form.Status,
		UpdateTime: time.Time{},
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
