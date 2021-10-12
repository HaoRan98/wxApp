package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"wxApp/models"
	"wxApp/pkg/logging"
	"wxApp/pkg/setting"
	"wxApp/pkg/util"
	"wxApp/pkg/waitgroup"
	"wxApp/redis"
	"wxApp/routers"
	server2 "wxApp/server"
)

func init() {
	setting.Setup()
	logging.Setup()
	util.Setup()
	models.Setup()
	models.InitCasbin()
}

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @host http://127.0.0.1
// @BasePath /
func main() {

	server2.TimeAccessToken()

	err := redis.SetupRedis()
	log.Println(err)

	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)
	log.Printf("[info] start http server listening %s", endPoint)

	err = server.ListenAndServe()
	if err != nil {
		log.Printf("init listen server fail:%v", err)
	}

	waitgroup.WG.Wait()
}


