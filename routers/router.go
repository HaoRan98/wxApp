package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gopkg.in/olahol/melody.v1"
	"log"
	"time"
	_ "wxApp/docs"
	"wxApp/middleware/cors"
	"wxApp/middleware/jwt"
	v1 "wxApp/routers/api/v1"
	v2 "wxApp/routers/api/v2"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	//r.MaxMultipartMemory = int64(setting.AppSetting.FileMaxSize) << 20

	var mr = melody.New()
	mr.Config.MaxMessageSize = 40960 * 2
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.CORSMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Static("/css", "runtime/static/css")
	r.Static("/js", "runtime/static/js")
	r.Static("/img", "runtime/static/img")

	r.POST("/login", v1.WxLogin)
	r.GET("/ws", v1.Websocket(mr))
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	// 微信端
	apiv1.Use(jwt.JWT())
	{



	}

	apiv2 := r.Group("/api/v2")
	// 服务端
	apiv2.Use(jwt.JWT())
	{
		// 人员
		// 创建人员信息
		apiv2.POST("/user/create",v2.CreateUser)
		// 修改人员信息
		apiv2.POST("/user/edit",v2.EditUser)
		// 删除人员
		apiv2.POST("/user/delete",v2.DeleteUser)
		// 查看人员列表
		//apiv2.POST("/user/list",v2.GetUserInfo)
		// 查看人员信息
		apiv2.POST("/user/detail",v2.GetUserInfo)

		// 地址
		// 创建地址
		apiv2.POST("/address/create",v2.CreateAddress)
		// 修改地址
		apiv2.POST("/address/edit",v2.EditAddress)
		// 删除地址
		apiv2.POST("/address/delete",v2.DeleteAddress)
		// 地址列表
		apiv2.POST("/address/list",v2.GetAddress)
		// 地址详情
		apiv2.POST("/address/detail",v2.GetAddressInfo)

		// 商品
		// 上传商品信息
		apiv2.POST("/product/create",v2.CreateProduct)
		// 修改商品信息
		apiv2.POST("/product/edit",v2.EditProduct)
		// 删除商品
		apiv2.POST("/product/delete",v2.DeleteProduct)
		// 商品列表
		apiv2.POST("/product/list",v2.GetProduct)
		// 商品详情
		apiv2.POST("/product/detail",v2.GetProductInfo)


	}


	// 保证文本顺序输出
	go func() {
		time.Sleep(100 * time.Millisecond)
		// In order to ensure that the text order output can be deleted
		log.Println(`默认自动化文档地址:http://127.0.0.1:80/swagger/index.html`)
	}()
	return r
}

func CheckError(err error) {
	log.Println(err)
}
