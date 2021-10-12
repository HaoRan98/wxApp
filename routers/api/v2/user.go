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

type User struct {
	ID			string 		`json:"id"`
	Username  	string   	`json:"username" gorm:"unique_index;not null"`
	Nickname  	string    	`json:"nickname"`
	Password  	string    	`json:"password"`
	AvatarUrl 	string    	`json:"avatar_url" gorm:"default:'static/upload/avatar/default.png'"`
	AddressID	string 		`json:"address_id"`
	Status      int     	`json:"status" gorm:"default:1"`
	Dept      	int       	`json:"dept_id" `
	Phone      	string      `json:"phone" `
	Email      	string      `json:"email" `
	Remark  	string  	`json:"remark" `
}

type UserInfo struct {
	User 	*models.SysUser
	Address []*models.Address
	Order   []*models.Order
}

func CreateUser(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form User
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	user := &models.SysUser{
		ID:        util.RandomString(30),
		Username:  form.Username,
		Nickname:  form.Nickname,
		Password:  util.EncodeMD5(form.Password),
		AvatarUrl: form.AvatarUrl,
		Status:    1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Phone:     form.Phone,
		Email:     form.Email,
		Remark:    form.Remark,
		SysRole:   models.SysRole{},
	}

	err := models.CreateData(user)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ADDRESS_CREATE_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func EditUser(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form User
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	user  := &models.SysUser{
		ID:        form.ID,
		Username:  form.Username,
		Nickname:  form.Nickname,
		Password:  util.EncodeMD5(form.Password),
		AvatarUrl: form.AvatarUrl,
		Status:    0,
		UpdatedAt: time.Time{},
		Phone:     form.Phone,
		Email:     form.Email,
		Remark:    form.Remark,
		SysRole:   models.SysRole{},
	}

	err := models.EditData(user)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ADDRESS_CREATE_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func DeleteUser(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form User
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	user  := &models.SysUser{
		ID: form.ID,
	}

	err := models.DeleteData(user)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ADDRESS_CREATE_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func GetUserInfo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form User
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	user , err := models.GetUserInfo(form.ID)
	if err != nil {
		log.Println(err)
	}

	address ,err := models.GetAddress(form.ID)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ADDRESS_GET_ERROR, nil)
		return
	}

	order ,err := models.GetOrderInfo(form.ID)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ORDER_GET_ERROR, nil)
		return
	}

	data := UserInfo{
		User: user,
		Address: address,
		Order: order,
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)

}
