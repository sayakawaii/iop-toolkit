package controller

import (
	"net/http"
	"omciAnalyzer/api"
	"omciAnalyzer/dao"
	"omciAnalyzer/models"
	"omciAnalyzer/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func isKeyExist(record any, key string, value any) bool {
	err := dao.MysqlRecordDataReadFirst(record, key, value)
	if err == nil {
		return true
	}
	return false
}

func (controller *AppController) LibraryAddCustomer(context *gin.Context) {
	customerContext := context.PostForm("name")

	if customerContext == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "empty name"})
		return
	}

	if isKeyExist(&dao.Customer{}, "name", customerContext) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "name exist"})
		return
	}

	customer := &dao.Customer{Name: customerContext}
	if err := dao.MysqlRecordDataInsert(customer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "insert sql error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": true})
}

func (controller *AppController) LibraryAddVendor(context *gin.Context) {
	vendorContext, _ := context.GetPostForm("name")

	if vendorContext == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "empty name"})
		return
	}

	if isKeyExist(&dao.Vendor{}, "name", vendorContext) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "name exist"})
		return
	}

	vendor := &dao.Vendor{Name: vendorContext}
	if err := dao.MysqlRecordDataInsert(vendor); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "insert sql error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": true})
}

func (controller *AppController) LibrarySearch(context *gin.Context) {
	searchContext, _ := context.GetPostForm("search")
	recordList := dao.IOPLibraryRecordDataSearch(searchContext)

	if recordList != nil {
		context.JSON(http.StatusOK, recordList)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	}
}

func (controller *AppController) LibraryAddRecord(context *gin.Context) {
	inputLogID := ""
	form, _ := context.MultipartForm()
	files := form.File["file-input-log"]
	for _, file := range files {
		thisLog := new(dao.OmciAnalyzerRequestRecord)
		thisLog.SetLogStorageInfo(api.SaveUploadedFileWithSessionID(true, file, context))
		thisLog.RequestKey = models.GenerateNanoID()
		thisLog.Progress = 20

		dao.MysqlRecordDataInsert(thisLog)
		//start analyzer routine
		go api.AnalyzerDataProc(thisLog)

		inputLogID = thisLog.RequestKey
	}
	var f dao.IOPLibraryRecordForm
	models.MapFormToStruct(form.Value, &f)

	id, _ := context.GetPostForm("ID")
	var err error
	if id != "" {
		var record dao.IOPLibraryRecord
		dao.MysqlRecordDataReadFirst(&record, "id", id)
		old := record
		record.SetFormData(&f)
		if inputLogID != "" {
			record.InputLog = inputLogID
		}
		err = dao.MysqlRecordDataUpdateByRecord(old, &record)
		utils.Log("old record: ", old)
		utils.Log("updated record: ", record)
	} else {
		var record dao.IOPLibraryRecord
		record.SetFormData(&f)
		record.InputLog = inputLogID
		err = dao.MysqlRecordDataInsert(&record)
		utils.Log("new record: ", record)
	}
	if err == nil {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	}
}

func (controller *AppController) LibraryDeleteRecord(context *gin.Context) {
	id, _ := context.GetPostForm("ID")
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
		return
	}
	//delete
	var record dao.IOPLibraryRecord
	err := dao.MysqlRecordDataReadFirst(&record, "id", id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
		return
	}
	err = dao.MysqlRecordDataDelete(&record)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	}
}

func (controller *AppController) LibraryUpdateRecord(context *gin.Context) {
	id, _ := context.GetPostForm("id")
	var record dao.IOPLibraryRecord
	err := dao.MysqlRecordDataReadFirst(&record, "id", id)
	if err == nil {
		context.JSON(http.StatusOK, record)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	}
}

func (controller *AppController) LibraryGetCustomers(context *gin.Context) {
	var customer []string
	var customerList []dao.Customer
	dao.MysqlRecordDataReadAll(&customerList)
	for _, value := range customerList {
		customer = append(customer, value.Name)
	}

	if customer != nil {
		context.JSON(http.StatusOK, customer)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	}
}

func (controller *AppController) LibraryGetVendors(context *gin.Context) {
	var vendor []string
	var vendorList []dao.Vendor
	dao.MysqlRecordDataReadAll(&vendorList)
	for _, value := range vendorList {
		vendor = append(vendor, value.Name)
	}

	if vendor != nil {
		context.JSON(http.StatusOK, vendor)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
	}
}

func (controller *AppController) LibraryDownload(context *gin.Context) {
	requestKey := context.Query("requestKey")
	reqPath := ""
	r := new(dao.OmciAnalyzerRequestRecord)
	err := dao.MysqlRecordDataRead(r, "request_key", requestKey)
	if err == nil {
		reqPath = r.LogPath
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}
	absTarget, info, err := resolveTarget(reqPath)
	if err != nil {
		// map common errors to appropriate status codes
		switch {
		case err.Error() == "missing path":
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case strings.Contains(err.Error(), "access denied"):
			context.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case strings.Contains(err.Error(), "server configuration error"):
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		default:
			// os.Stat errors etc.
			if os.IsNotExist(err) {
				context.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
			} else {
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}
		return
	}

	if info.IsDir() {
		context.JSON(http.StatusBadRequest, gin.H{"error": "path is a directory"})
		return
	}

	// send binary file as attachment (sets content-disposition)
	context.FileAttachment(absTarget, filepath.Base(absTarget))
}
