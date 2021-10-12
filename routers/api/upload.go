package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wxApp/pkg/app"
	"wxApp/pkg/e"
	"wxApp/pkg/logging"
	"wxApp/pkg/upload"
)

func UploadFile(c *gin.Context) {
	appG := app.Gin{C: c}
	fHeader, err := c.FormFile("file")
	if err != nil {
		logging.Error(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	if fHeader == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}


	fileName := upload.GetFileName(fHeader.Filename)
	fullPath := upload.GetFileFullPath()
	src := fullPath + fileName

	if !upload.CheckFileExt(fileName) {
		log.Println("校验文件错误，文件格式不正确")
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_FILE_FORMAT, nil)
		return
	}

	err = upload.CheckFile(fullPath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_FILE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(fHeader, src); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_FILE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"file_name": fHeader.Filename,
		"file_url":  upload.GetFileFullUrl(fileName),
	})
}



