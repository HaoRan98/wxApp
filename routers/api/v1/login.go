package v1

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wxApp/pkg/app"
	"wxApp/pkg/e"
	"wxApp/pkg/util"
	"wxApp/server"
)

type Login struct {
	Code 	string	`json:"code"`
}

func WxLogin(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form Login
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	log.Println(form)

	wxLoginResp,err := server.WXLogin(form.Code)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.WX_LOGIN_ERR, err)
		return
	}

	token ,err := util.GenerateToken(wxLoginResp.OpenId,wxLoginResp.UnionId,wxLoginResp.SessionKey)
	if err != nil {
		log.Println(err)
	}

	// 这里用openid和sessionkey的串接 进行MD5之后作为该用户的自定义登录态
	//mySession := GetMD5Encode(wxLoginResp.OpenId + wxLoginResp.SessionKey)

	// 更新库库
	data := map[string]string{
		"openid":wxLoginResp.OpenId,
		"unionid":wxLoginResp.UnionId,
		"sessionKey":wxLoginResp.SessionKey,
		"toke":token,
	}
	//
	//err = changgui.CreateData(data)
	//if err != nil {
	//	log.Println(err)
	//	appG.Response(http.StatusInternalServerError, e.WRITE_IN_FAIL, nil)
	//	return
	//}

	// 接下来可以将openid 和 sessionkey, mySession 存储到数据库中,
	// 但这里要保证mySession 唯一, 以便于用mySession去索引openid 和sessionkey
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
// 将一个字符串进行MD5加密后返回加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

