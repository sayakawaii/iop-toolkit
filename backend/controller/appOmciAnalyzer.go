package controller

import (
	"net/http"
	"omciAnalyzer/api"
	"omciAnalyzer/controller/omci"
	"omciAnalyzer/dao"
	"omciAnalyzer/models"
	"omciAnalyzer/utils"
	"slices"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (controller *AppController) OmciAnalyzerRequest(context *gin.Context) {
	var ret []dao.OmciAnalyzerRequestRecord
	//multiple files in one POST
	form, err := context.MultipartForm()
	if err != nil || form == nil {
		utils.Log("OmciAnalyzerRequest: multipart parse failed:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid or incomplete multipart upload",
		})
		return
	}
	files := form.File["omcianalyzerFile"]
	for _, file := range files {
		thisLog := new(dao.OmciAnalyzerRequestRecord)
		thisLog.SetLogStorageInfo(api.SaveUploadedFileWithSessionID(false, file, context))
		thisLog.RequestKey = models.GenerateNanoID()
		thisLog.Progress = 10
		thisLog.Status = dao.OmciAnalyzerRequestRecordStatusProcessing

		dao.MysqlRecordDataInsert(thisLog)
		//start analyzer routine
		go api.AnalyzerDataProc(thisLog)
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

func (controller *AppController) OmciAnalyzerMinio(context *gin.Context) {
	var ret []dao.OmciAnalyzerRequestRecord

	// Parse JSON body from POST request
	type MinioRequest struct {
		Minio string `json:"minio"`
	}
	var req MinioRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		utils.Log("Error binding JSON:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	minioKey := req.Minio
	utils.Log("Received minioKey:", minioKey)

	if minioKey == "" {
		utils.Log("minioKey is empty")
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "minioKey is required",
		})
		return
	}

	// Create log record for the single downloaded file
	thisLog := new(dao.OmciAnalyzerRequestRecord)
	thisLog.SetLogStorageInfo(api.SaveMinioFileWithSessionID(minioKey, context))
	thisLog.RequestKey = models.GenerateNanoID()
	thisLog.Progress = 10
	thisLog.Status = dao.OmciAnalyzerRequestRecordStatusProcessing

	dao.MysqlRecordDataInsert(thisLog)
	// Start analyzer routine
	go api.AnalyzerDataProc(thisLog)
	ret = append(ret, *thisLog)

	if ret != nil {
		context.JSON(http.StatusOK, ret)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	}
}

func (controller *AppController) GetOmciAnalyzerProgressData(context *gin.Context) {
	var retLog []dao.OmciAnalyzerRequestRecord
	//multiple logs in one POST
	form, err := context.MultipartForm()
	if err != nil || form == nil {
		utils.Log("GetOmciAnalyzerProgressData: multipart parse failed:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid or incomplete multipart upload",
		})
		return
	}
	logs := form.Value["requestKey"]
	for _, key := range logs {
		r := new(dao.OmciAnalyzerRequestRecord)
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

func (controller *AppController) GetOmciAnalyzerHistoryData(context *gin.Context) {
	session := sessions.Default(context)
	var user string
	key := session.Get("userid")
	if key == nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	} else {
		user = key.(string)
	}

	var retLog []dao.OmciAnalyzerRequestRecord
	//read all history data
	err := dao.MysqlRecordDataRead(&retLog, "log_user", user)
	if err != nil {
		retLog = nil
	}

	slices.Reverse(retLog)
	if retLog != nil {
		context.JSON(http.StatusOK, retLog)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	}
}

func pathClean(path string) string {
	path = strings.TrimPrefix(path, "./")
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimPrefix(path, "static/")
	if !strings.HasPrefix(path, "uploads/") {
		path = "uploads/" + path
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return path
}

func (controller *AppController) GetOmciAnalyzerDiagramData(context *gin.Context) {
	type respOnusData struct {
		Path     string `json:"path"`
		Category string `json:"category"`
	}
	requestKey := context.Query("requestKey")
	onuName := context.Query("onuName")

	record := new(dao.OmciAnalyzerRequestRecord)
	err := dao.MysqlRecordDataRead(record, "request_key", requestKey)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	} else {
		var onusResp []respOnusData
		for k, v := range record.Content.OnuContent {
			if onuName == k {
				for _, diagram := range v.Diagram {
					path := pathClean(diagram.Path)
					onusResp = append(onusResp, respOnusData{
						Path:     path,
						Category: diagram.Category,
					})
				}
			}
		}
		context.JSON(http.StatusOK, onusResp)
	}
}

func (controller *AppController) GetOmciAnalyzerPlantumlData(context *gin.Context) {
	type respOnusData struct {
		Path     string `json:"path"`
		Category string `json:"category"`
	}
	requestKey := context.Query("requestKey")
	onuName := context.Query("onuName")

	record := new(dao.OmciAnalyzerRequestRecord)
	err := dao.MysqlRecordDataRead(record, "request_key", requestKey)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	} else {
		var onusResp []respOnusData
		for k, v := range record.Content.OnuContent {
			if onuName == k {
				for _, plantuml := range v.Plantuml {
					path := pathClean(plantuml.Path)
					onusResp = append(onusResp, respOnusData{
						Path:     path,
						Category: plantuml.Category,
					})
				}
			}
		}
		context.JSON(http.StatusOK, onusResp)
	}
}

func (controller *AppController) GetOmciAnalyzerOnusData(context *gin.Context) {
	type respOnusData struct {
		OnuName   string `json:"onuName"`
		SwVersion string `json:"swVersion"`
		HwVersion string `json:"hwVersion"`
		Status    string `json:"status"`
	}
	requestKey := context.Query("requestKey")

	record := new(dao.OmciAnalyzerRequestRecord)
	err := dao.MysqlRecordDataRead(record, "request_key", requestKey)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	} else {
		var onusResp []respOnusData
		for onuName, onu := range record.Content.OnuContent {
			tmp := respOnusData{
				OnuName:   onuName,
				SwVersion: onu.SwVersion,
				HwVersion: onu.HwVersion,
				Status:    "Active",
			}
			onusResp = append(onusResp, tmp)
		}
		context.JSON(http.StatusOK, onusResp)
	}
}

func (controller *AppController) GetOmciAnalyzerOmciData(context *gin.Context) {
	requestKey := context.Query("requestKey")
	onuName := context.Query("onuName")

	record := new(dao.OmciAnalyzerRequestRecord)
	err := dao.MysqlRecordDataRead(record, "request_key", requestKey)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	} else {
		var onusResp []omci.RespOmciData
		for k, v := range record.Content.OnuContent {
			if onuName == k {
				onusResp = omci.AssembleOmciData(v.JsonPath)
			}
		}
		context.JSON(http.StatusOK, onusResp)
	}
}

func (controller *AppController) GetOmciAnalyzerCounters(context *gin.Context) {
	type respCounters struct {
		Total uint64 `json:"total"`
		Today uint64 `json:"today"`
	}
	total, err := dao.MysqlRecordDataCount(&dao.OmciAnalyzerRequestRecord{})
	if err != nil {
		utils.Log("Error:", err)
	}
	today, err := dao.MysqlRecordDataCountToday(&dao.OmciAnalyzerRequestRecord{})
	if err != nil {
		utils.Log("error:", err)
	}
	context.JSON(http.StatusOK, respCounters{
		Total: total,
		Today: today,
	})
}
