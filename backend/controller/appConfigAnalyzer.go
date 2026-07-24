package controller

import (
	"net/http"
	"omciAnalyzer/api"
	"omciAnalyzer/dao"
	"omciAnalyzer/models"
	"path"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (controller *AppController) ConfigAnalyzerRequest(context *gin.Context) {
	var ret []dao.ConfigAnalyzerRequestRecord
	var ponLogID uint
	var ponLogPath string
	//multiple files in one POST
	form, _ := context.MultipartForm()
	pon := form.File["file-input-log-pon"]
	for _, file := range pon {
		extName := path.Ext(file.Filename)
		if extName != ".xml" {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": true,
			})
			return
		}
		thisLog := new(dao.ConfigAnalyzerPonLogRecord)
		thisLog.SetLogStorageInfo(api.SaveUploadedFileWithSessionID(false, file, context))
		dao.MysqlRecordDataInsert(thisLog)
		ponLogID = thisLog.ID
		ponLogPath = thisLog.LogPath
	}
	onus := form.File["file-input-log-onu"]
	for _, file := range onus {
		thisLog := new(dao.ConfigAnalyzerRequestRecord)
		thisLog.SetLogStorageInfo(api.SaveUploadedFileWithSessionID(false, file, context))
		thisLog.RequestKey = models.GenerateNanoID()
		thisLog.PonLogID = ponLogID
		thisLog.Progress = 20

		extName := path.Ext(file.Filename)
		if extName == ".json" {
			filename := path.Base(thisLog.LogPath)
			thisLog.Onus = api.OnuJsonBound(strings.TrimSuffix(filename, extName), strings.TrimSuffix(thisLog.LogName, path.Ext(thisLog.LogName)), thisLog.LogDir)
		} else {
			thisLog.Onus = api.OnuLogShape(thisLog.LogPath, thisLog.LogDir)
		}

		if len(thisLog.Onus) == 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": true,
			})
			return
		}

		dao.MysqlRecordDataInsert(thisLog)
		api.SendDataToYangHelper(thisLog, ponLogPath)
		ret = append(ret, *thisLog)
	}

	if ret != nil {
		context.JSON(http.StatusOK, ret)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	}
}

func (controller *AppController) GetConfigAnalyzerProgressData(context *gin.Context) {
	var retLog []dao.ConfigAnalyzerRequestRecord
	//multiple logs in one POST
	form, _ := context.MultipartForm()
	logs := form.Value["requestKey"]
	for _, key := range logs {
		r := new(dao.ConfigAnalyzerRequestRecord)
		err := dao.MysqlRecordDataRead(r, "request_key", key)
		if err == nil {
			retLog = append(retLog, *r)
		}
	}

	if retLog != nil {
		context.JSON(http.StatusOK, retLog)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	}
}

func (controller *AppController) GetConfigAnalyzerHistoryData(context *gin.Context) {
	session := sessions.Default(context)
	var user string
	key := session.Get("userid")
	if key == nil {
		//session not set
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	} else {
		//session exist
		user = key.(string)
	}

	var retLog []dao.ConfigAnalyzerRequestRecord
	load, _ := context.GetPostForm("load")
	switch load {
	case "latest":
		//read latest data
		err := dao.MysqlRecordDataReadLast(&retLog, "log_user", user)
		if err != nil {
			retLog = nil
		}
	case "history":
		//read all history data
		err := dao.MysqlRecordDataRead(&retLog, "log_user", user)
		if err != nil {
			retLog = nil
		}
	}

	if retLog != nil {
		context.JSON(http.StatusOK, retLog)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	}
}
