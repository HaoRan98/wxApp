package v2

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wxApp/models"
	"wxApp/pkg/app"
	"wxApp/pkg/e"
)

func GetSlideshow(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
	)

	data , err := models.GetSlideshow()
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)

}
