package dao

import (
	"encoding/json"
	service "omciAnalyzer/service/omcianalyzer"

	"github.com/jinzhu/gorm"
)

type LogStorageInfo struct {
	LogID      string `json:"logID" gorm:"type:varchar(32)"`
	LogUser    string `json:"logUser" gorm:"type:varchar(256)"`
	LogName    string `json:"logName" gorm:"type:varchar(256)"`
	UploadTime string `json:"uploadTime" gorm:"type:varchar(32)"`
	LogDir     string `json:"logDir" gorm:"type:varchar(256)"`
	LogPath    string `json:"logPath" gorm:"type:varchar(256)"`
}

// common file storage struct define
type LogStorage interface {
	GetLogStorageInfo() *LogStorageInfo
	SetLogStorageInfo(*LogStorageInfo)
}

const (
	OmciAnalyzerRequestRecordStatusProcessing = "processing"
	OmciAnalyzerRequestRecordStatusSuccess    = "success"
	OmciAnalyzerRequestRecordStatusError      = "error"
)

// inherit LogStorageInfo，implement GetLogStorageInfo function
type OmciAnalyzerRequestRecord struct {
	gorm.Model
	LogStorageInfo
	// other fields
	RequestKey  string             `json:"requestKey" gorm:"type:varchar(128)"`
	Progress    uint8              `json:"progress" gorm:"type:int"`
	LogType     string             `json:"logType" gorm:"type:varchar(32)"`
	Content     service.LogContent `json:"-" gorm:"-"`
	ContentJSON string             `json:"-" gorm:"type:longtext"`
	Status      string             `json:"status" gorm:"type:varchar(16)"`
	ErrorMsg    string             `json:"errorMsg" gorm:"type:varchar(256)"`
}

// set OmciAnalyzerRequestRecord table name as "omci_analyzer_request_records"
func (r *OmciAnalyzerRequestRecord) TableName() string {
	return "omci_analyzer_request_records"
}

func (r *OmciAnalyzerRequestRecord) GetLogStorageInfo() *LogStorageInfo {
	return &r.LogStorageInfo
}

func (r *OmciAnalyzerRequestRecord) SetLogStorageInfo(lsi *LogStorageInfo) {
	r.LogStorageInfo = *lsi
}

func (r *OmciAnalyzerRequestRecord) BeforeSave(tx *gorm.DB) error {
	jsonBytes, err := json.Marshal(r.Content)
	if err != nil {
		return err
	}
	r.ContentJSON = string(jsonBytes)
	return nil
}

func (r *OmciAnalyzerRequestRecord) AfterFind(tx *gorm.DB) error {
	if r == nil {
		return nil
	}
	if err := json.Unmarshal([]byte(r.ContentJSON), &r.Content); err != nil {
		return err
	}
	return nil
}

type Customer struct {
	gorm.Model
	Name string
}

// set Customer table name as "customers"
func (r *Customer) TableName() string {
	return "customers"
}

type Vendor struct {
	gorm.Model
	Name string
}

// set Vendor table name as "vendors"
func (r *Vendor) TableName() string {
	return "vendors"
}

type IOPLibraryRecordForm struct {
	Customer  string `form:"chosen-target-customer"`
	RcrNum    string `form:"text-rcr"`
	Vendor    string `form:"select-vendor"`
	OnuType   string `form:"text-onu-type"`
	PonType   string `form:"select-pon-type"`
	RGMode    string `form:"select-rg-mode"`
	EquipID   string `form:"text-equip-id"`
	HwVersion string `form:"text-hw-ver"`
	SwVersion string `form:"text-sw-ver"`
	LSRelease string `form:"text-ls-release"`
	Comments  string `form:"text-comments"`
}

type IOPLibraryRecord struct {
	gorm.Model
	IOPLibraryRecordForm
	InputLog string
}

// set IOPLibraryRecord table name as "iop_library_records"
func (r *IOPLibraryRecord) TableName() string {
	return "iop_library_records"
}

func (r *IOPLibraryRecord) SetFormData(f *IOPLibraryRecordForm) {
	r.IOPLibraryRecordForm = *f
}

func (r *IOPLibraryRecord) GetFormData(f *IOPLibraryRecordForm) IOPLibraryRecordForm {
	return r.IOPLibraryRecordForm
}

type ConfigAnalyzerPonLogRecord struct {
	gorm.Model
	LogStorageInfo
}

// set ConfigAnalyzerPonLogRecord table name as "config_analyzer_pon_log_records"
func (r *ConfigAnalyzerPonLogRecord) TableName() string {
	return "config_analyzer_pon_log_records"
}

func (r *ConfigAnalyzerPonLogRecord) GetLogStorageInfo() *LogStorageInfo {
	return &r.LogStorageInfo
}

func (r *ConfigAnalyzerPonLogRecord) SetLogStorageInfo(lsi *LogStorageInfo) {
	r.LogStorageInfo = *lsi
}

type ConfigAnalyzerOnuCopyInfo struct {
	ActionKey     string `json:"actionKey"`
	OnuNativeName string `json:"onuNativeName"`
	Identifier    string `json:"identifier"`
	Operation     string `json:"operation"`
	OnuCopyPath   string `json:"onuCopyPath"`
	OnuCopyDir    string `json:"onuCopyDir"`
	OnuCopyName   string `json:"onuCopyName"`
}

func (c *ConfigAnalyzerOnuCopyInfo) SetActionKey(key string) string {
	c.ActionKey = key
	return c.ActionKey
}

func (c *ConfigAnalyzerOnuCopyInfo) GetActionKey() string {
	return c.ActionKey
}

func (c *ConfigAnalyzerOnuCopyInfo) NewActionKey() {
	c.ActionKey = c.OnuCopyName + "_" + c.Identifier + "_" + c.Operation
}

type YangHelperResponseInfo struct {
	PlantUmlPath string `json:"plantUmlPath"`
	DiagramPath  string `json:"diagramPath"`
}

type ConfigAnalyzerContent struct {
	Copy     ConfigAnalyzerOnuCopyInfo
	Response YangHelperResponseInfo
	Finished bool `json:"finished"`
}

type ConfigAnalyzerRequestRecord struct {
	gorm.Model
	LogStorageInfo
	RequestKey string                           `json:"requestKey"`
	Progress   uint8                            `json:"progress"`
	PonLogID   uint                             `json:"ponLogID"`
	Onus       map[string]ConfigAnalyzerContent `json:"-" gorm:"-"`
	OnusJSON   string                           `json:"-" gorm:"type:json"`
}


// set ConfigAnalyzerRequestRecord table name as "config_analyzer_request_records"
func (r *ConfigAnalyzerRequestRecord) TableName() string {
	return "config_analyzer_request_records"
}

func (r *ConfigAnalyzerRequestRecord) GetLogStorageInfo() *LogStorageInfo {
	return &r.LogStorageInfo
}

func (r *ConfigAnalyzerRequestRecord) SetLogStorageInfo(lsi *LogStorageInfo) {
	r.LogStorageInfo = *lsi
}

func (r *ConfigAnalyzerRequestRecord) BeforeSave(tx *gorm.DB) error {
	jsonBytes, err := json.Marshal(r.Onus)
	if err != nil {
		return err
	}
	r.OnusJSON = string(jsonBytes)
	return nil
}

func (r *ConfigAnalyzerRequestRecord) AfterFind(tx *gorm.DB) error {
	if r == nil {
		return nil
	}
	if err := json.Unmarshal([]byte(r.OnusJSON), &r.Onus); err != nil {
		return err
	}
	return nil
}

type ConfigAnalyzerResponseRecord struct {
	gorm.Model
	Payload string `json:"payload"`
	Status  string `json:"status"`
}

// set ConfigAnalyzerResponseRecord table name as "config_analyzer_response_records"
func (r *ConfigAnalyzerResponseRecord) TableName() string {
	return "config_analyzer_response_records"
}

type CollectorConnectionInfo struct {
	Platform string `form:"chosen-target-platform"`
	Board    string `form:"select-board"`
	IP       string `form:"text-ipaddr"`
	Port     uint   `form:"text-port"`
	User     string `form:"text-user"`
	Password string `form:"text-password"`
}
