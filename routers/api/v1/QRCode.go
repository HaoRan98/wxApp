package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wxApp/pkg/app"
	"wxApp/pkg/e"
	"wxApp/server"
)

func GetQRCode(c *gin.Context) {

	var (
		appG = app.Gin{C: c}

	)

	data ,err := server.GetQRCode()
	if err != nil {
		log.Default()
		appG.Response(http.StatusInternalServerError, e.PRODUCT_GET_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)

}
