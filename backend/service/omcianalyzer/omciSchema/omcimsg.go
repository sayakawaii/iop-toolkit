package omciSchema

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"omciAnalyzer/utils"

	"github.com/mailru/easyjson"
)

type OMCI_DIRECTION byte

const (
	_ OMCI_DIRECTION = iota
	TX
	RX
)

type EngineOmciHeader struct {
	Tcid        uint16
	AR          byte
	AK          byte
	MsgType     byte
	MsgTypeName string
	DevId       byte
	MeClass     uint16
	MeClassName string
	MeInst      uint16
}

type MT byte

const (
	InvalidMT MT = iota
	_
	_
	_
	Create
	_
	Delete
	_
	Set
	Get
	_
	GetAllAlarms
	GetAllAlarmsNext
	MIBUpload
	MIBUploadNext
	MIBReset
	Alarm
	AVC
	Test
	StartSwDld
	DldSection
	EndSwDld
	ActivateSw
	CommitSw
	SyncTime
	Reboot
	GetNext
	TestResult
	GetCurrentData
	SetTable
	SetNoSync = SetTable
)

type MsgTypeNameMapType map[MT]string

type MeClassType uint16

const (
	_                                               MeClassType = iota
	MeClass_OnuData                                             = 2
	MeClass_CircuitPack                                         = 6
	MeClass_SoftwareImage                                       = 7
	MeClass_PhysicalPathTerminationPointEthernetUNI             = 11
	MeClass_OnuRemoteDebug                                      = 158
	MeClass_OnuG                                                = 256
	MeClass_Onu2G                                               = 257
	MeClass_TCont                                               = 262
	MeClass_AniG                                                = 263
	MeClass_PriortyQueue                                        = 277
	MeClass_TrafficScheduler                                    = 278
	MeClass_Last_Standard_OMCI_ME                               = 453

	MeClass_OntAggrGemPortPMHistoryData = 65281
	MeClass_OntGenericV2                = 65296
)

type RR byte

const (
	//Standard result reason
	CmdProcessSuccessful RR = iota
	// CmdProcessError
	// CmdNotSupported
	ParameterError
	UnknownME
	UnknownMEInst
	DeviceBusy
	InstExists
	_
	AttrFailedOrUnknown
	NumStandardRR
)

// Additional result reason
const (
	OmciTimeout RR = 0xf0 + iota
	OmciBiDirTruncatedResp
	OmciTcidMismatch
	OmciOnuDisconnected
	OmciAborted
	OmciUniDirTimeout
	OmciUniDirTruncatedResp
	OmciHeadMismatch
)

type CheckTCID byte

const (
	IgnoreTCID CheckTCID = iota
	MatchTCID
)

type ExpectedOmci byte

const (
	NoAdditionalOmci ExpectedOmci = iota
	UniDirOmciExpected
)

type OmciRecvCallback func(RR, OmciRespMsg, OmciMsgInterface) bool

type AttrMapType map[string]interface{}

//easyjson:json
type OmciMsgTypeInfo struct {
	MsgTypeInfo struct {
		Timeout    uint16        `json:"timeout"`
		Debouncing uint16        `json:"debouncing"`
		Retries    uint16        `json:"retries"`
		DBTimeout  uint16        `json:"device-busy-timeout"`
		DBRetries  uint16        `json:"device-busy-retries"`
		MsgTypes   []OmciMsgType `json:"message-types"`
	} `json:"message-type-info"`
}

//easyjson:json
type OmciMsgType struct {
	Name         string  `json:"name"`
	Id           byte    `json:"id"`
	IncrementMds bool    `json:"increment-mib-data-sync"`
	Timeout      *uint16 `json:"timeout,omitempty"`
	Debouncing   *uint16 `json:"debouncing,omitempty"`
	Retries      *uint16 `json:"retries,omitempty"`
	DBTimeout    *uint16 `json:"device-busy-timeout,omitempty"`
	DBRetries    *uint16 `json:"device-busy-retries,omitempty"`
}

type OmciMsgTypeListType map[byte]*OmciMsgType

type OmciContent struct {
	AttrMask uint16
	payload  []byte
	attr     []Attribute
}

type OmciTrailer struct {
	Cpcs uint32
}

type OmciMsg struct {
	header  EngineOmciHeader
	content OmciContent
	trailer OmciTrailer
}

type OmciRespMsg struct {
	header EngineOmciHeader
	//resp    OmciResp
	Result       RR
	content      OmciContent
	AttrExecMask uint16
	OptAttrMask  uint16
}

type OmciPkt struct {
	onuName string
	payload []byte
}

type OnuTcidRxChnlInfo struct {
	tcid     uint16
	omciChnl chan OmciPkt
	//yangChnl chan YangInfo
}

type OmciAlarm struct {
	header EngineOmciHeader
	bitmap []byte
	pad    [3]byte
	seqNum byte
}

//easyjson:json
type OmciTrace struct {
	payload  string                 //`json:"payload,omitempty"`
	Header   EngineOmciHeader       `json:"header"`
	Contents map[string]interface{} `json:"contents"`
}

//option1
// type OmciContents struct {
// 	Contents map[string]interface{} `json:"contents"`
// }
//option1

type OmciMsgInterface interface {
	Parse(def *SchemaDef, payload []byte) error
	//String() string
	OmciTrace(*SchemaDef) OmciTrace
	OmciHeader() EngineOmciHeader
	OmciResultReason() RR
	OmciPayload() []byte
	OmciContext(*SchemaDef) OmciContext
}

type OmciCreateTxMsg struct {
	OmciMsgInterface
	Payload []byte
	Header  EngineOmciHeader
	Attrs   []Attribute
}

type OmciCreateRxMsg struct {
	OmciMsgInterface
	Payload      []byte
	Header       EngineOmciHeader
	Result       RR
	AttrExecMask uint16
}

type OmciDeleteTxMsg struct {
	OmciMsgInterface
	Payload []byte
	Header  EngineOmciHeader
}

type OmciDeleteRxMsg struct {
	OmciMsgInterface
	Payload []byte
	Header  EngineOmciHeader
	Result  RR
}

type OmciSetTxMsg struct {
	OmciMsgInterface
	Payload  []byte
	Header   EngineOmciHeader
	AttrMask uint16
	Attrs    []Attribute
}

type OmciSetRxMsg struct {
	OmciMsgInterface
	Payload      []byte
	Header       EngineOmciHeader
	Result       RR
	OptAttrMask  uint16
	AttrExecMask uint16
}

type OmciGetTxMsg struct {
	OmciMsgInterface
	Payload  []byte
	Header   EngineOmciHeader
	AttrMask uint16
	Attrs    []Attribute
}

type OmciGetRxMsg struct {
	OmciMsgInterface
	Payload      []byte
	Header       EngineOmciHeader
	Result       RR
	AttrMask     uint16
	Attrs        []Attribute
	OptAttrMask  uint16
	AttrExecMask uint16
	attrMap      AttrMapType
}

type OmciGetAllAlarmsNextTxMsg struct {
	OmciMsgInterface
	Payload            []byte
	Header             EngineOmciHeader
	CommandSequenceNum uint16
}

type OmciGetAllAlarmsNextRxMsg struct {
	OmciMsgInterface
	Payload     []byte
	Header      EngineOmciHeader
	MeClass     uint16
	MeInst      uint16
	AlarmBitMap []byte
}

type OmciMIBUploadTxMsg struct {
	OmciMsgInterface
	Payload []byte
	Header  EngineOmciHeader
}

type OmciMIBUploadRxMsg struct {
	OmciMsgInterface
	Payload              []byte
	Header               EngineOmciHeader
	NumMIBUploadCommands uint16
}

type OmciMIBUploadNextTxMsg struct {
	OmciMsgInterface
	Payload            []byte
	Header             EngineOmciHeader
	CommandSequenceNum uint16
}

type OmciMIBUploadNextRxMsg struct {
	OmciMsgInterface
	Payload  []byte
	Header   EngineOmciHeader
	MeClass  uint16
	MeInst   uint16
	AttrMask uint16
	Attrs    []Attribute
	attrMap  AttrMapType
}

type OmciGetNextTxMsg struct {
	OmciMsgInterface
	Payload            []byte
	Header             EngineOmciHeader
	AttrMask           uint16
	CommandSequenceNum uint16
}

type OmciGetNextRxMsg struct {
	OmciMsgInterface
	Payload  []byte
	Header   EngineOmciHeader
	Result   RR
	AttrMask uint16
	Attr     Attribute
}

type OmciMIBResetTxMsg struct {
	OmciMsgInterface
	Payload []byte
	Header  EngineOmciHeader
}

type OmciMIBResetRxMsg struct {
	OmciMsgInterface
	Payload []byte
	Header  EngineOmciHeader
	Result  RR
}

type OmciGetCurrentDataTxMsg OmciGetTxMsg

type OmciGetCurrentDataRxMsg OmciGetRxMsg

type OmciStartSwDldTxMsg struct {
	OmciMsgInterface
	Header               EngineOmciHeader
	ImageSize            uint32
	WindowSize           uint8
	ParallelCircuitPacks uint8
	MSSoftwareImage      uint8
	LSSoftwareImage      uint8
	SwImgeIDs            []byte
	Payload              []byte
}

type OmciStartSwDldRxMsg struct {
	OmciMsgInterface
	Header              EngineOmciHeader
	Result              RR
	WindowSize          uint8
	InstancesResponding uint8
	SoftwareImageInst   uint16
	RRPayloadsSoftImage RR
	RRadditional        []byte
	TurboSectionSize    uint16 /*Stores the value of SectionSize if turbo mode is supported. If turbo is not supported, this variable should NOT be handled.*/
	Payload             []byte
}

type OmciDldSectionTxMsg struct {
	OmciMsgInterface
	Header                EngineOmciHeader
	Payload               []byte
	DownloadSectionNumber byte
	Data                  []byte
}

type OmciDldSectionRxMsg struct {
	OmciMsgInterface
	Header                EngineOmciHeader
	DownloadSectionNumber uint8
	Result                RR
	Payload               []byte
}

type OmciEndSwDldTxMsg struct {
	OmciMsgInterface
	Header          EngineOmciHeader
	CRC             uint32
	ImageSize       uint32
	ParallelDld     uint8
	MSSoftwareImage uint8
	LSSoftwareImage uint8
	SwImgeIDs       []byte
	Payload         []byte
}

type OmciEndSwDldRxMsg struct {
	OmciMsgInterface
	Header              EngineOmciHeader
	Result              RR
	InstancesResponding byte
	SoftwareImageInst   uint16
	RRPayloadsSoftImage byte
	RRadditional        []byte
	Payload             []byte
}

type OmciActivateImgTxMsg struct {
	OmciMsgInterface
	Header  EngineOmciHeader
	Flag    byte
	Payload []byte
}

type OmciActivateImgRxMsg struct {
	OmciMsgInterface
	Header  EngineOmciHeader
	Result  RR
	Payload []byte
}

type OmciCommitImageTxMsg struct {
	OmciMsgInterface
	Header  EngineOmciHeader
	Payload []byte
}

type OmciCommitImageRxMsg struct {
	OmciMsgInterface
	Header  EngineOmciHeader
	Result  RR
	Payload []byte
}

type OmciGetAllAlarmsTxMsg struct {
	OmciMsgInterface
	Payload            []byte
	Header             EngineOmciHeader
	AlarmRetrievalMode uint8
}

type OmciGetAllAlarmsRxMsg struct {
	OmciMsgInterface
	Payload                 []byte
	Header                  EngineOmciHeader
	NumGetAllAlarmsCommands uint16
}
type OmciTestOnuTxMsg struct {
	OmciMsgInterface
	Header     EngineOmciHeader
	Payload    []byte
	SelectTest uint8
	Attrs      []Attribute
	attrMap    AttrMapType
}

type OmciTestOnuRxMsg struct {
	OmciMsgInterface
	Header  EngineOmciHeader
	Result  RR
	Payload []byte
}

type OmciTestResultRxMsg struct {
	OmciMsgInterface
	Header   EngineOmciHeader
	Payload  []byte
	Contents []byte
	Attrs    []Attribute
	attrMap  AttrMapType
}

type OmciAlarmRxMsg struct {
	OmciMsgInterface
	Header           EngineOmciHeader
	AlarmBitMap      [28]byte
	AlarmSeqNum      byte
	AlarmName        string
	AlarmDescription string
	Payload          []byte
}

type OmciAVCRxMsg struct {
	OmciMsgInterface
	Payload  []byte
	Header   EngineOmciHeader
	AttrMask uint16
	Attrs    []Attribute
	attrMap  AttrMapType
}

type OmciRebootOnuTxMsg struct {
	OmciMsgInterface
	Payload []byte
	Header  EngineOmciHeader
}

type OmciRebootOnuRxMsg struct {
	OmciMsgInterface
	Header  EngineOmciHeader
	Result  RR
	Payload []byte
}

type OmciSyncTimeTxMsg struct {
	OmciMsgInterface
	Payload []byte
	Header  EngineOmciHeader
	Result  RR
	Year    uint16
	Month   uint8
	Day     uint8
	Hour    uint8
	Minute  uint8
	Second  uint8
}

type OmciSyncTimeRxMsg struct {
	OmciMsgInterface
	Header  EngineOmciHeader
	Result  RR
	Payload []byte
}

type OmciHandshakAVC struct {
	header   OmciHeader
	AttrMask uint16
	AvcValue uint16
}

type RetryMethod byte

const (
	AvoidRetryOnTimeout RetryMethod = iota
	AutoRetryOnTimeout
)

var omciMsgTypeInfo *OmciMsgTypeInfo

var commandReturnStringMap = map[RR]string{
	//CmdProcessError: "CmdProcessError",
	//CmdNotSupported: "CmdNotSupported",
	ParameterError: "ParameterError",
	UnknownME:      "UnknownME",
	DeviceBusy:     "DeviceBusy",
	UnknownMEInst:  "UnknownMEInst"}

var testResponseStringMap = map[RR]string{
	//CmdProcessError: "command processing error",
	//CmdNotSupported: "command not supported",
	ParameterError: "parameter error",
	UnknownME:      "parameter error",
	UnknownMEInst:  "parameter error",
	DeviceBusy:     "device busy"}

var payloadLenghtThreshold int = 40

func (oh *EngineOmciHeader) populateInfo(def *SchemaDef) {
	switch oh.MsgType {
	case 4:
		oh.MsgTypeName = "Create"
	case 6:
		oh.MsgTypeName = "Delete"
	case 9:
		oh.MsgTypeName = "Get"
	case 8:
		oh.MsgTypeName = "Set"
	case 11:
		oh.MsgTypeName = "GetAllAlarms"
	case 12:
		oh.MsgTypeName = "GetAllAlarmsNext"
	case 13:
		oh.MsgTypeName = "MIBUpload"
	case 14:
		oh.MsgTypeName = "MIBUploadNext"
	case 15:
		oh.MsgTypeName = "MIBReset"
	case 16:
		oh.MsgTypeName = "Alarm"
	case 17:
		oh.MsgTypeName = "AVC"
	case 18:
		oh.MsgTypeName = "Test"
	case 19:
		oh.MsgTypeName = "StartSwDld"
	case 20:
		oh.MsgTypeName = "DldSection"
	case 21:
		oh.MsgTypeName = "EndSwDld"
	case 22:
		oh.MsgTypeName = "ActivateSw"
	case 23:
		oh.MsgTypeName = "CommitSw"
	case 24:
		oh.MsgTypeName = "SyncTime"
	case 25:
		oh.MsgTypeName = "Reboot"
	case 26:
		oh.MsgTypeName = "GetNext"
	case 27:
		oh.MsgTypeName = "TestResult"
	case 28:
		oh.MsgTypeName = "GetCurrentData"
	case 29:
		oh.MsgTypeName = "SetTable"
	default:
		oh.MsgTypeName = "UNKNOWN"
	}
	if schema, e := def.MeSchemaByClassId(ClassIdType(oh.MeClass)); e == nil {
		oh.MeClassName = schema.Name()
	} else {
		oh.MeClassName = "UNKNOWN"
	}
	return
}

func LoadOmciMsgTypeInfo(path string) (OmciMsgTypeListType, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		errString := fmt.Sprintf("Load Omci Message Type Info failed to read file [%v] error [%v]\n", path, err.Error())
		utils.Log(errString)
		return nil, errors.New(errString)
	}
	omciMsgTypeInfo = new(OmciMsgTypeInfo)
	err = easyjson.Unmarshal(raw, omciMsgTypeInfo)
	if err != nil {
		errString := fmt.Sprintf("Load Omci Message Type Info failed to parse file [%v] error [%v]\n", path, err.Error())
		utils.Log(errString)
		return nil, errors.New(errString)
	}

	f := func(v uint16) *uint16 {
		return &v
	}
	omciMsgTypeList := make(OmciMsgTypeListType)

	for count, msgType := range omciMsgTypeInfo.MsgTypeInfo.MsgTypes {
		if msgType.Timeout == nil {
			omciMsgTypeInfo.MsgTypeInfo.MsgTypes[count].Timeout = f(omciMsgTypeInfo.MsgTypeInfo.Timeout)
		}
		if msgType.Retries == nil {
			omciMsgTypeInfo.MsgTypeInfo.MsgTypes[count].Retries = f(omciMsgTypeInfo.MsgTypeInfo.Retries)
		}
		if msgType.DBTimeout == nil {
			omciMsgTypeInfo.MsgTypeInfo.MsgTypes[count].DBTimeout = f(omciMsgTypeInfo.MsgTypeInfo.DBTimeout)
		}
		if msgType.DBRetries == nil {
			omciMsgTypeInfo.MsgTypeInfo.MsgTypes[count].DBRetries = f(omciMsgTypeInfo.MsgTypeInfo.DBRetries)
		}
		if msgType.Debouncing == nil {
			omciMsgTypeInfo.MsgTypeInfo.MsgTypes[count].Debouncing = f(omciMsgTypeInfo.MsgTypeInfo.Debouncing)
		}
		omciMsgTypeList[msgType.Id] = &omciMsgTypeInfo.MsgTypeInfo.MsgTypes[count]
	}
	return omciMsgTypeList, nil
}

func (oh *EngineOmciHeader) decode(bytes []byte) uint16 {
	oh.Tcid = bytesUint16(bytes[:2])
	b := bytes[2]
	oh.AR = (b & 0x40) >> 6
	oh.AK = (b & 0x20) >> 5
	oh.MsgType = b & 0x1f
	oh.DevId = bytes[3]
	oh.MeClass = bytesUint16(bytes[4:6])
	oh.MeInst = bytesUint16(bytes[6:8])
	return 8
}

func (oh *EngineOmciHeader) isAlarmType() bool {
	return MT(oh.MsgType) == Alarm
}

func (oh *EngineOmciHeader) isAVCType() bool {
	return MT(oh.MsgType) == AVC
}

func (oh *EngineOmciHeader) isTestResultype() bool {
	return MT(oh.MsgType) == TestResult
}

func (oh *EngineOmciHeader) infoString() string {
	//logger.Debugf("\x1b[35;1m%+v\x1b[0m\n", *oh)
	return fmt.Sprintf("%+v ", *oh)
}

func (oh *EngineOmciHeader) isAK() bool {
	if oh.AK == 0 {
		return false
	}
	return true
}

func (oc *OmciContent) pad(width uint16) {
	oc.payload = pad(oc.payload, width)
}

func (oc *OmciContent) bytes() []byte {
	return oc.payload
}

func (oc *OmciContent) infoString() string {
	//logger.Debugf("\x1b[36;1mAttrMask: %v attrs: %+v\x1b[0m", oc.AttrMask, oc.attr)
	return fmt.Sprintf("AttrMask: %v attrs: %+v\n", oc.AttrMask, oc.attr)
}

func (oc *OmciContent) decode(mt MT, bytes []byte) (hl, tl uint16, e error) {
	e = nil
	hl = 0
	tl = 0
	switch mt {
	case Create:
	case Delete:
	case Set:
		oc.AttrMask = bytesUint16(bytes[0:2])
		hl = 2
	case Get:
		oc.AttrMask = bytesUint16(bytes[0:2])
		hl = 2
	default:
		e = errors.New("Invalid msg type " + string(mt))
	}
	oc.payload = append(oc.payload, bytes[hl:len(bytes)-int(tl)]...)
	return
}

func (msg *OmciMsg) decode(def *SchemaDef, bytes []byte) {
	nextByteLoc := msg.header.decode(bytes)
	msg.content.decode(MT(msg.header.MsgType), bytes[nextByteLoc:])
	schema, e := def.MeSchemaByClassId(ClassIdType(msg.header.MeClass))
	if e == nil {
		msg.content.attr = schema.ParseTxOmci(msg)
	}
	return
}

func isValidPayload(payload []byte) bool {
	if len(payload) < payloadLenghtThreshold {
		utils.Log("isValidPayload")
		return false
	}
	return true
}

func (msg *OmciMsg) infoString(payload []byte) string {
	ret := hex.Dump(payload)
	ret += msg.header.infoString()
	ret += msg.content.infoString()
	return ret
}

func (or *OmciRespMsg) decode(mt MT, bytes []byte) (hl, tl uint16, e error) {
	e = nil
	hl = 0
	tl = 0
	switch mt {
	case Create:
		or.Result = RR(bytes[0])
		or.AttrExecMask = bytesUint16(bytes[1:3])
		hl = 3
	case Delete:
		or.Result = RR(bytes[0])
		hl = 1
	case Set:
		or.Result = RR(bytes[0])
		or.OptAttrMask = bytesUint16(bytes[1:3])
		or.AttrExecMask = bytesUint16(bytes[3:5])
		hl = 5
	case Get:
		or.Result = RR(bytes[0])
		or.content.AttrMask = bytesUint16(bytes[1:3])
		or.OptAttrMask = bytesUint16(bytes[28:30])
		or.AttrExecMask = bytesUint16(bytes[30:32])
		hl = 3
		tl = 4
	case GetNext:
		or.Result = RR(bytes[0])
		or.content.AttrMask = bytesUint16(bytes[1:3])
		hl = 3
		tl = 4
	case GetCurrentData:
		or.Result = RR(bytes[0])
		or.content.AttrMask = bytesUint16(bytes[1:3])
		hl = 3
	case Test:
		or.Result = RR(bytes[0])
		hl = 1
	case TestResult:
		or.Result = RR(bytes[0])
		hl = 1
	case MIBReset:
		or.Result = RR(bytes[0])
		hl = 1
	default:
		e = errors.New("Invalid msg type " + string(mt))
	}
	return
}

func (omciAlarm *OmciAlarm) decode(payload []byte) {
	ohl := omciAlarm.header.decode(payload)
	omciAlarm.bitmap = payload[ohl : ohl+28]
	omciAlarm.seqNum = payload[39]
	return
}

func (omciTrace *OmciTrace) String() string {
	retBytes, _ := easyjson.Marshal(omciTrace)
	return omciTrace.payload + string(retBytes)
}

func (omciTrace *OmciTrace) EasyString() string {
	retBytes, _ := easyjson.Marshal(omciTrace)
	return string(retBytes)
}

// func (omciContext *OmciContext) EasyString() string {
// 	retBytes, _ := easyjson.Marshal(omciContext)
// 	return string(retBytes)
// }

func (msg *OmciCreateTxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	schema, e := def.MeSchemaByClassId(ClassIdType(msg.Header.MeClass))
	if e == nil {
		if msg.Header.DevId == 0x0B {
			msg.Attrs, e = schema.Attributes(schema.Me.Attributes.SbcMask, msg.Payload[hl+2:])
		} else {
			msg.Attrs, e = schema.Attributes(schema.Me.Attributes.SbcMask, msg.Payload[hl:])
		}
	}
	return e
}

func (msg *OmciCreateTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	omciTrace.Contents["Attrs"] = msg.Attrs
	return
}

func (msg *OmciCreateTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciCreateTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciCreateTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciCreateRxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.Result = RR(msg.Payload[hl])
	msg.AttrExecMask = bytesUint16(msg.Payload[hl+1 : hl+3])
	return nil
}

func (msg *OmciCreateRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["Result"] = msg.Result
	omciTrace.Contents["AttrExecMask"] = msg.AttrExecMask
	return
}

func (msg *OmciCreateRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciCreateRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciCreateRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciDeleteTxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	msg.Header.decode(msg.Payload)
	return nil
}

func (msg *OmciDeleteTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	return
}

func (msg *OmciRebootOnuTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	return
}

func (msg *OmciDeleteTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciDeleteTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciDeleteTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciDeleteRxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.Result = RR(msg.Payload[hl])
	return nil
}

func (msg *OmciDeleteRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["Result"] = msg.Result
	return
}

func (msg *OmciDeleteRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciDeleteRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciDeleteRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciSetTxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	if msg.Header.DevId == 0x0B {
		msg.AttrMask = bytesUint16(msg.Payload[hl+2 : hl+4])
	} else {
		msg.AttrMask = bytesUint16(msg.Payload[hl : hl+2])
	}
	schema, e := def.MeSchemaByClassId(ClassIdType(msg.Header.MeClass))
	if e == nil {
		if msg.Header.DevId == 0x0B {
			msg.Attrs, e = schema.Attributes(msg.AttrMask, msg.Payload[hl+4:])
		} else {
			msg.Attrs, e = schema.Attributes(msg.AttrMask, msg.Payload[hl+2:])
		}
	}
	return e
}

func (msg *OmciSetTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	omciTrace.Contents["AttrMask"] = msg.AttrMask
	omciTrace.Contents["Attrs"] = msg.Attrs
	return
}

//option2
// func (omciTrace *OmciTrace) ToOmciContext() (omciContext OmciContext, err error) {
// 	omciContext.Header.Tcid = uint64(omciTrace.Header.Tcid)
// 	omciContext.Header.AR = uint(omciTrace.Header.AR)
// 	omciContext.Header.AK = uint(omciTrace.Header.AK)
// 	omciContext.Header.MsgType = omciTrace.Header.MsgTypeName
// 	omciContext.Header.DevId = strconv.Itoa(int(omciTrace.Header.DevId))
// 	omciContext.Header.MeClass = uint64(omciTrace.Header.MeClass)
// 	omciContext.Header.MeClassName = omciTrace.Header.MeClassName
// 	omciContext.Header.MeInst = uint64(omciTrace.Header.MeInst)
// 	omciContext.Contents = omciTrace.Contents

// 	return
// }

func (msg *OmciSetTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciSetTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciSetTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciSetRxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.Result = RR(msg.Payload[hl])
	msg.OptAttrMask = bytesUint16(msg.Payload[hl+1 : hl+3])
	msg.AttrExecMask = bytesUint16(msg.Payload[hl+3 : hl+5])
	return nil
}

func (msg *OmciSetRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["Result"] = msg.Result
	omciTrace.Contents["OptAttrMask"] = msg.OptAttrMask
	omciTrace.Contents["AttrExecMask"] = msg.AttrExecMask
	return
}

func (msg *OmciSetRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciSetRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciSetRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciGetTxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	if msg.Header.DevId == 0x0B {
		msg.AttrMask = bytesUint16(msg.Payload[hl+2 : hl+4])
	} else {
		msg.AttrMask = bytesUint16(msg.Payload[hl : hl+2])
	}
	schema, e := def.MeSchemaByClassId(ClassIdType(msg.Header.MeClass))
	if e == nil {
		if msg.Header.DevId == 0x0B {
			msg.Attrs, e = schema.Attributes(msg.AttrMask, msg.Payload[hl+4:])
		} else {
			msg.Attrs, e = schema.Attributes(msg.AttrMask, msg.Payload[hl+2:])
		}
	}
	return e
}

func (msg *OmciGetTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	omciTrace.Contents["AttrMask"] = msg.AttrMask
	omciTrace.Contents["Attrs"] = msg.Attrs
	return
}

func (msg *OmciGetTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciGetTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciGetTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciGetRxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.Result = RR(msg.Payload[hl])
	if msg.Header.DevId == 0x0B {
		msg.AttrMask = bytesUint16(msg.Payload[hl+3 : hl+5])
	} else {
		msg.AttrMask = bytesUint16(msg.Payload[hl+1 : hl+3])
	}
	schema, e := def.MeSchemaByClassId(ClassIdType(msg.Header.MeClass))
	if e == nil {
		if msg.Header.DevId == 0x0B {
			msg.Attrs, e = schema.Attributes(msg.AttrMask, msg.Payload[hl+5:])
		} else {
			msg.Attrs, e = schema.Attributes(msg.AttrMask, msg.Payload[hl+3:])
		}
		for _, attr := range msg.Attrs {
			msg.attrMap[attr.Name] = attr.Value
		}
	}
	msg.OptAttrMask = bytesUint16(msg.Payload[hl+28 : hl+30])
	msg.AttrExecMask = bytesUint16(msg.Payload[hl+30 : hl+32])
	return nil
}

func (msg *OmciGetRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["Result"] = msg.Result
	omciTrace.Contents["AttrMask"] = msg.AttrMask
	omciTrace.Contents["Attrs"] = msg.Attrs
	omciTrace.Contents["OptAttrMask"] = msg.OptAttrMask
	omciTrace.Contents["AttrExecMask"] = msg.AttrExecMask
	return
}

func (msg *OmciGetRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciGetRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciGetAllAlarmsNextTxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.CommandSequenceNum = bytesUint16(msg.Payload[hl : hl+2])
	return
}

func (msg *OmciGetAllAlarmsNextTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	omciTrace.Contents["CommandSequenceNum"] = msg.CommandSequenceNum
	return
}

func (msg *OmciGetAllAlarmsNextTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciGetAllAlarmsNextTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciGetAllAlarmsNextRxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.MeClass = bytesUint16(msg.Payload[hl : hl+2])
	msg.MeInst = bytesUint16(msg.Payload[hl+2 : hl+4])
	msg.AlarmBitMap = msg.Payload[hl+4 : hl+32]
	return
}

func (msg *OmciGetAllAlarmsNextRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["MeClassForReportedAlarms"] = msg.MeClass
	omciTrace.Contents["MeInstForReportedAlarms"] = msg.MeInst
	omciTrace.Contents["BitMap"] = msg.AlarmBitMap
	return
}

func (msg *OmciGetAllAlarmsNextRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciGetAllAlarmsNextRxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciGetRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciMIBUploadTxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	msg.Header.decode(msg.Payload)
	return e
}

func (msg *OmciMIBUploadTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	return
}

func (msg *OmciMIBUploadTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciMIBUploadTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciMIBUploadTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciMIBUploadRxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.NumMIBUploadCommands = bytesUint16(msg.Payload[hl : hl+2])
	return nil
}

func (msg *OmciMIBUploadRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["NumMIBUploadCommands"] = msg.NumMIBUploadCommands
	return
}

func (msg *OmciMIBUploadRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciMIBUploadRxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciMIBUploadRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciMIBUploadNextTxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.CommandSequenceNum = bytesUint16(msg.Payload[hl : hl+2])
	return e
}

func (msg *OmciMIBUploadNextTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	omciTrace.Contents["CommandSequenceNum"] = msg.CommandSequenceNum
	return
}

func (msg *OmciMIBUploadNextTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciMIBUploadNextTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciMIBUploadNextTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciMIBUploadNextRxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	if msg.Header.DevId == 0x0B {
		msg.MeClass = bytesUint16(msg.Payload[hl+4 : hl+6])
		msg.MeInst = bytesUint16(msg.Payload[hl+6 : hl+8])
		msg.AttrMask = bytesUint16(msg.Payload[hl+8 : hl+10])
	} else {
		msg.MeClass = bytesUint16(msg.Payload[hl : hl+2])
		msg.MeInst = bytesUint16(msg.Payload[hl+2 : hl+4])
		msg.AttrMask = bytesUint16(msg.Payload[hl+4 : hl+6])
	}
	schema, e := def.MeSchemaByClassId(ClassIdType(msg.MeClass))
	if e == nil {
		if msg.Header.DevId == 0x0B {
			msg.Attrs, e = schema.Attributes(msg.AttrMask, msg.Payload[hl+10:])
		} else {
			msg.Attrs, e = schema.Attributes(msg.AttrMask, msg.Payload[hl+6:])
		}
		for _, attr := range msg.Attrs {
			msg.attrMap[attr.Name] = attr.Value
		}
	}
	return nil
}

func (msg *OmciMIBUploadNextRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["MeClass"] = msg.MeClass
	schema, e := def.MeSchemaByClassId(ClassIdType(msg.MeClass))
	if e == nil {
		omciTrace.Contents["MeClassName"] = schema.Name()
	} else {
		omciTrace.Contents["MeClassName"] = "UNKNOWN"
	}
	omciTrace.Contents["MeInst"] = msg.MeInst
	omciTrace.Contents["AttrMask"] = msg.AttrMask
	omciTrace.Contents["Attrs"] = msg.Attrs
	return
}

func (msg *OmciMIBUploadNextRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciMIBUploadNextRxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciGetNextTxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.AttrMask = bytesUint16(msg.Payload[hl : hl+2])
	msg.CommandSequenceNum = bytesUint16(msg.Payload[hl+2 : hl+4])
	return e
}

func (msg *OmciGetNextTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	omciTrace.Contents["AttrMask"] = msg.AttrMask
	omciTrace.Contents["CommandSequenceNum"] = msg.CommandSequenceNum
	return
}

func (msg *OmciGetNextTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciGetNextTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciGetNextRxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.Result = RR(msg.Payload[hl])
	msg.AttrMask = bytesUint16(msg.Payload[hl+1 : hl+3])
	return e
}

func (msg *OmciGetNextRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["Result"] = msg.Result
	omciTrace.Contents["AttrMask"] = msg.AttrMask
	return
}

func (msg *OmciGetNextRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciGetNextRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciMIBUploadNextRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciMIBResetTxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	msg.Header.decode(msg.Payload)
	return e
}

func (msg *OmciMIBResetTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	return
}

func (msg *OmciMIBResetTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciMIBResetTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciMIBResetTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciMIBResetRxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.Result = RR(msg.Payload[hl])
	return nil
}

func (msg *OmciMIBResetRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["Result"] = msg.Result
	return
}

func (msg *OmciMIBResetRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciMIBResetRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciMIBResetRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciGetCurrentDataTxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	v := &OmciGetTxMsg{}
	e = v.Parse(def, payload)
	*msg = OmciGetCurrentDataTxMsg(*v)
	return
}

func (msg *OmciGetCurrentDataTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	v := &OmciGetTxMsg{}
	*v = OmciGetTxMsg(*msg)
	return v.OmciTrace(def)
}

func (msg *OmciGetCurrentDataTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciGetCurrentDataTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciGetCurrentDataTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciGetCurrentDataRxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	v := &OmciGetRxMsg{attrMap: make(AttrMapType)}
	e = v.Parse(def, payload)
	*msg = OmciGetCurrentDataRxMsg(*v)
	return
}

func (msg *OmciGetCurrentDataRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	v := &OmciGetRxMsg{}
	*v = OmciGetRxMsg(*msg)
	return v.OmciTrace(def)
}

func (msg *OmciGetCurrentDataRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciGetCurrentDataRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciGetCurrentDataRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciStartSwDldTxMsg) Parse(def *SchemaDef, payload []byte) error {

	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.WindowSize = bytesUInteger(msg.Payload[hl:hl+1], 1).(uint8)
	msg.ImageSize = bytesUInteger(msg.Payload[hl+1:hl+5], 4).(uint32)
	msg.ParallelCircuitPacks = bytesUInteger(msg.Payload[hl+5:hl+6], 1).(uint8)
	msg.MSSoftwareImage = bytesUInteger(msg.Payload[hl+6:hl+7], 1).(uint8)
	msg.LSSoftwareImage = bytesUInteger(msg.Payload[hl+7:hl+8], 1).(uint8)

	return nil
}

func (msg *OmciStartSwDldTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {

	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})

	omciTrace.Contents["WindowSize"] = msg.WindowSize
	omciTrace.Contents["ImageSize"] = msg.ImageSize
	omciTrace.Contents["ParallelCircuitPacks"] = msg.ParallelCircuitPacks
	omciTrace.Contents["MSSoftwareImage"] = msg.MSSoftwareImage
	omciTrace.Contents["LSSoftwareImage"] = msg.LSSoftwareImage

	return
}

func (msg *OmciStartSwDldTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciStartSwDldTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciStartSwDldTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciStartSwDldRxMsg) Parse(def *SchemaDef, payload []byte) error {

	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.Result = RR(msg.Payload[hl])
	msg.WindowSize = bytesUInteger(msg.Payload[hl+1:hl+2], 1).(uint8)
	msg.InstancesResponding = bytesUInteger(msg.Payload[hl+2:hl+3], 1).(uint8)
	msg.SoftwareImageInst = bytesUInteger(msg.Payload[hl+3:hl+5], 2).(uint16)
	msg.RRPayloadsSoftImage = RR(msg.Payload[hl+6])
	msg.TurboSectionSize = bytesUInteger(msg.Payload[hl+2:hl+4], 2).(uint16) //Stores the SectionSize 4 proprietary Turbo mode
	return nil
}

func (msg *OmciStartSwDldRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})

	omciTrace.Contents["WindowSize"] = msg.WindowSize
	omciTrace.Contents["InstancesResponding"] = msg.InstancesResponding
	omciTrace.Contents["SoftwareImageInst"] = msg.SoftwareImageInst
	omciTrace.Contents["RRPayloadsSoftImage"] = msg.RRPayloadsSoftImage
	return
}

func (msg *OmciStartSwDldRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciStartSwDldRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciStartSwDldRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciDldSectionTxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)

	msg.DownloadSectionNumber = bytesUInteger(msg.Payload[hl:hl+1], 1).(uint8)
	msg.Data = msg.Payload[hl+1 : hl+32]

	return nil
}
func (msg *OmciDldSectionTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	omciTrace.Contents["DownloadSectionNumber"] = msg.DownloadSectionNumber
	omciTrace.Contents["Data"] = msg.Data
	return
}

func (msg *OmciDldSectionTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciDldSectionTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciDldSectionTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciDldSectionRxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)

	msg.Result = RR(msg.Payload[hl])
	msg.DownloadSectionNumber = bytesUInteger(msg.Payload[hl+1:hl+2], 1).(uint8)

	return nil
}

func (msg *OmciDldSectionRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})

	omciTrace.Contents["Result"] = msg.Result
	omciTrace.Contents["DownloadSectionNumber"] = msg.DownloadSectionNumber

	return
}

func (msg *OmciDldSectionRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciDldSectionRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciDldSectionRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciEndSwDldTxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)

	msg.CRC = bytesUInteger(msg.Payload[hl:hl+4], 4).(uint32)
	msg.ImageSize = bytesUInteger(msg.Payload[hl+4:hl+8], 4).(uint32)
	msg.ParallelDld = bytesUInteger(msg.Payload[hl+8:hl+9], 1).(uint8)

	return nil
}

func (msg *OmciEndSwDldTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	omciTrace.Contents["CRC"] = msg.CRC
	omciTrace.Contents["ImageSize"] = msg.ImageSize
	omciTrace.Contents["ParallelDld"] = msg.ParallelDld

	return
}

func (msg *OmciEndSwDldTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciEndSwDldTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciEndSwDldTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciEndSwDldRxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)

	msg.Result = RR(msg.Payload[hl])

	return nil
}

func (msg *OmciEndSwDldRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header

	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["Result"] = msg.Result

	return
}

func (msg *OmciEndSwDldRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciEndSwDldRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciEndSwDldRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciActivateImgTxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)

	msg.Flag = msg.Payload[hl]

	return nil
}

func (msg *OmciActivateImgTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	omciTrace.Contents["Flag"] = msg.Flag

	return
}

func (msg *OmciActivateImgTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciActivateImgTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciActivateImgTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciActivateImgRxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)

	msg.Result = RR(msg.Payload[hl])

	return nil
}

func (msg *OmciActivateImgRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})

	omciTrace.Contents["Result"] = msg.Result

	return
}

func (msg *OmciActivateImgRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciActivateImgRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciActivateImgRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciCommitImageTxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	_ = msg.Header.decode(msg.Payload)

	return nil
}

func (msg *OmciCommitImageTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	return
}

func (msg *OmciCommitImageTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciCommitImageTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciCommitImageTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciCommitImageRxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.Result = RR(msg.Payload[hl])
	return nil
}

func (msg *OmciCommitImageRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})

	omciTrace.Contents["Result"] = msg.Result
	return
}

func (msg *OmciCommitImageRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciCommitImageRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciCommitImageRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciTestOnuTxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.SelectTest = uint8(msg.Payload[hl])
	schema, e := def.MeSchemaByClassId(ClassIdType(msg.Header.MeClass))
	if e == nil && len(msg.Payload) > 0 {
		var omciRspMsg OmciRespMsg
		omciRspMsg.content.payload = msg.Payload[hl:]
		msg.Attrs = schema.ParseTestAttrs(schema.Me.TestRequestAttributes.TestRequestAttribute, &omciRspMsg)
		for _, attr := range msg.Attrs {
			msg.attrMap[attr.Name] = attr.Value
		}
	}

	return nil
}

func (msg *OmciRebootOnuTxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	msg.Header.decode(msg.Payload)
	return nil
}

func (msg *OmciTestOnuTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	omciTrace.Contents["payload"] = msg.Payload
	omciTrace.Contents["TestRequestAttrs"] = msg.Attrs
	return
}

func (msg *OmciTestOnuTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciTestOnuTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciTestOnuTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciTestOnuRxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.Result = RR(msg.Payload[hl])
	return nil
}

func (msg *OmciRebootOnuRxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.Result = RR(msg.Payload[hl])
	return nil
}

func (msg *OmciTestOnuRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["Result"] = msg.Result
	return
}

func (msg *OmciTestOnuRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciTestOnuRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciRebootOnuRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciTestOnuRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciTestResultRxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	if len(msg.Payload) < 32 {
		utils.Log("Invalid payload size for Test Result")
		e = fmt.Errorf("Invalid payload size for Test Result")
		return
	}
	hl := msg.Header.decode(msg.Payload)
	//FNMS-67192
	msg.Header.populateInfo(def)

	msg.Contents = append(msg.Contents, msg.Payload[hl:hl+32]...)

	schema, e := def.MeSchemaByClassId(ClassIdType(msg.Header.MeClass))
	if e == nil && len(msg.Payload) > 0 {
		var omciRspMsg OmciRespMsg
		omciRspMsg.content.payload = msg.Payload[hl:]
		msg.Attrs = schema.ParseTestAttrs(schema.Me.TestResultAttributes.TestResultAttribute, &omciRspMsg)
		for _, attr := range msg.Attrs {
			msg.attrMap[attr.Name] = attr.Value
		}
	}

	return nil
}

func (msg *OmciTestResultRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["MeClass"] = msg.Header.MeClass
	omciTrace.Contents["MeInst"] = msg.Header.MeInst
	omciTrace.Contents["TestResults"] = msg.Attrs
	omciTrace.Contents["Contents"] = msg.Contents
	return
}

func (msg *OmciRebootOnuRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["Result"] = msg.Result
	return
}

func (msg *OmciTestResultRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciTestResultRxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciTestResultRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func isAlarmActive(alarmBitmap [28]byte, alarmNumber uint8) bool {
	if alarmNumber >= 224 {
		return false // 超出范围
	}
	byteIndex := alarmNumber / 8       // 计算在哪个字节
	bitOffset := 7 - (alarmNumber % 8) // 计算 bit 偏移（大端对齐）
	return (alarmBitmap[byteIndex] & (1 << bitOffset)) != 0
}

func (msg *OmciAlarmRxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	copy(msg.AlarmBitMap[:], msg.Payload[hl:hl+28])
	msg.AlarmSeqNum = msg.Payload[39]
	schema, e := def.MeSchemaByClassId(ClassIdType(msg.Header.MeClass))
	if e == nil && len(msg.Payload) > 0 {
		alarms := schema.Me.Alarms.Alarm
		for _, alarm := range alarms {
			if isAlarmActive(msg.AlarmBitMap, alarm.Number) {
				msg.AlarmName = alarm.Name
				msg.AlarmDescription = alarm.Description
			}
		}
	}
	return nil
}

func (msg *OmciAlarmRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["AlarmSeqNum"] = msg.AlarmSeqNum
	omciTrace.Contents["AlarmName"] = msg.AlarmName
	omciTrace.Contents["AlarmDescription"] = msg.AlarmDescription
	return
}

func (msg *OmciAlarmRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciAlarmRxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciAlarmRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciAVCRxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	if msg.Header.DevId == 0x0B {
		msg.AttrMask = bytesUint16(msg.Payload[hl+2 : hl+4])
	} else {
		msg.AttrMask = bytesUint16(msg.Payload[hl : hl+2])
	}
	schema, e := def.MeSchemaByClassId(ClassIdType(msg.Header.MeClass))
	if e == nil {
		if msg.Header.DevId == 0x0B {
			msg.Attrs, e = schema.Attributes(msg.AttrMask, msg.Payload[hl+4:])
		} else {
			msg.Attrs, e = schema.Attributes(msg.AttrMask, msg.Payload[hl+2:])
		}
		for _, attr := range msg.Attrs {
			msg.attrMap[attr.Name] = attr.Value
		}
	}
	return nil
}

func (msg *OmciAVCRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["AttrMask"] = msg.AttrMask
	omciTrace.Contents["Attrs"] = msg.Attrs
	return
}

func (msg *OmciAVCRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciAVCRxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciAVCRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciSyncTimeRxMsg) Parse(def *SchemaDef, payload []byte) error {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.Result = RR(msg.Payload[hl])
	return nil
}

func (msg *OmciSyncTimeRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	omciTrace.Contents["Result"] = msg.Result
	return
}

func (msg *OmciSyncTimeRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciSyncTimeRxMsg) OmciResultReason() RR {
	return msg.Result
}

func (msg *OmciSyncTimeRxMsg) OmciPayload() []byte {
	return msg.Payload
}

func (msg *OmciSyncTimeTxMsg) Parse(def *SchemaDef, payload []byte) (e error) {
	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.Year = bytesUInteger(msg.Payload[hl:hl+2], 2).(uint16)
	msg.Month = bytesUInteger(msg.Payload[hl+2:hl+3], 1).(uint8)
	msg.Day = bytesUInteger(msg.Payload[hl+3:hl+4], 1).(uint8)
	msg.Hour = bytesUInteger(msg.Payload[hl+4:hl+5], 1).(uint8)
	msg.Minute = bytesUInteger(msg.Payload[hl+5:hl+6], 1).(uint8)
	msg.Second = bytesUInteger(msg.Payload[hl+6:hl+7], 1).(uint8)

	return nil
}

func (msg *OmciSyncTimeTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {

	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})

	omciTrace.Contents["Year"] = msg.Year
	omciTrace.Contents["Month"] = msg.Month
	omciTrace.Contents["Day"] = msg.Day
	omciTrace.Contents["Hour"] = msg.Hour
	omciTrace.Contents["Minute"] = msg.Minute
	omciTrace.Contents["Second"] = msg.Second

	return
}

func (msg *OmciSyncTimeTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciSyncTimeTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciSyncTimeTxMsg) OmciPayload() []byte {
	return msg.Payload
}

func OmciParseMsgInterface(dir OMCI_DIRECTION, payload []byte, def *SchemaDef) (ret OmciMsgInterface, e error) {
	e = nil
	var oh EngineOmciHeader

	oh.decode(payload)
	switch MT(oh.MsgType) {
	case Create:
		if dir == TX {
			ret = &OmciCreateTxMsg{}
		} else {
			ret = &OmciCreateRxMsg{}
		}
	case Delete:
		if dir == TX {
			ret = &OmciDeleteTxMsg{}
		} else {
			ret = &OmciDeleteRxMsg{}
		}
	case Set:
		if dir == TX {
			ret = &OmciSetTxMsg{}
		} else {
			ret = &OmciSetRxMsg{}
		}
	case SetNoSync:
		if dir == TX {
			ret = &OmciSetTxMsg{}
		} else {
			ret = &OmciSetRxMsg{}
		}
	case Get:
		if dir == TX {
			ret = &OmciGetTxMsg{}
		} else {
			ret = &OmciGetRxMsg{attrMap: make(AttrMapType)}
		}
	case GetAllAlarmsNext:
		if dir == TX {
			ret = &OmciGetAllAlarmsNextTxMsg{}
		} else {
			ret = &OmciGetAllAlarmsNextRxMsg{}
		}
	case GetNext:
		if dir == TX {
			ret = &OmciGetNextTxMsg{}
		} else {
			ret = &OmciGetNextRxMsg{}
		}
	case MIBUpload:
		if dir == TX {
			ret = &OmciMIBUploadTxMsg{}
		} else {
			ret = &OmciMIBUploadRxMsg{}
		}
	case MIBUploadNext:
		if dir == TX {
			ret = &OmciMIBUploadNextTxMsg{}
		} else {
			ret = &OmciMIBUploadNextRxMsg{attrMap: make(AttrMapType)}
		}
	case MIBReset:
		if dir == TX {
			ret = &OmciMIBResetTxMsg{}
		} else {
			ret = &OmciMIBResetRxMsg{}
		}
	case GetCurrentData:
		if dir == TX {
			ret = &OmciGetCurrentDataTxMsg{}
		} else {
			ret = &OmciGetCurrentDataRxMsg{attrMap: make(AttrMapType)}
		}
	case StartSwDld:
		if dir == TX {
			ret = &OmciStartSwDldTxMsg{}
		} else {
			ret = &OmciStartSwDldRxMsg{}
		}
	case DldSection:
		if dir == TX {
			ret = &OmciDldSectionTxMsg{}
		} else {
			ret = &OmciDldSectionRxMsg{}
		}
	case EndSwDld:
		if dir == TX {
			ret = &OmciEndSwDldTxMsg{}
		} else {
			ret = &OmciEndSwDldRxMsg{}
		}
	case ActivateSw:
		if dir == TX {
			ret = &OmciActivateImgTxMsg{}
		} else {
			ret = &OmciActivateImgRxMsg{}
		}
	case CommitSw:
		if dir == TX {
			ret = &OmciCommitImageTxMsg{}
		} else {
			ret = &OmciCommitImageRxMsg{}
		}
	case GetAllAlarms:
		if dir == TX {
			ret = &OmciGetAllAlarmsTxMsg{}
		} else {
			ret = &OmciGetAllAlarmsRxMsg{}
		}
	case Test:
		if dir == TX {
			ret = &OmciTestOnuTxMsg{attrMap: make(AttrMapType)}
		} else {
			ret = &OmciTestOnuRxMsg{}
		}
	case TestResult:
		if dir == RX {
			ret = &OmciTestResultRxMsg{attrMap: make(AttrMapType)}
		}
	case Alarm:
		if dir == RX {
			ret = &OmciAlarmRxMsg{}
		}
	case Reboot:
		if dir == TX {
			ret = &OmciRebootOnuTxMsg{}
		} else {
			ret = &OmciRebootOnuRxMsg{}
		}
	case SyncTime:
		if dir == TX {
			ret = &OmciSyncTimeTxMsg{}
		} else {
			ret = &OmciSyncTimeRxMsg{}
		}
	case AVC:
		if dir == RX {
			ret = &OmciAVCRxMsg{attrMap: make(AttrMapType)}
		}
	default:
		e = fmt.Errorf("omci message parsing not available for type %v\n", oh.MsgType)
	}
	if e != nil {
		return
	}
	if ret == nil {
		utils.Log("unknown message type")
	} else {
		e = ret.Parse(def, payload)
	}
	return
}

func OmciParse(dir OMCI_DIRECTION, payload []byte) (ret OmciTrace, e error) {
	msg, e := OmciParseMsgInterface(dir, payload, DefaultSchemaDef())
	if e != nil {
		return
	}
	return OmciParseFromInterface(dir, msg)
}

func OmciParseFromInterface(dir OMCI_DIRECTION, msg OmciMsgInterface) (ret OmciTrace, e error) {
	e = nil
	ret = msg.OmciTrace(DefaultSchemaDef())
	return
}

func (msg *OmciGetAllAlarmsTxMsg) Parse(def *SchemaDef, payload []byte) error {

	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.AlarmRetrievalMode = bytesUInteger(msg.Payload[hl:hl+1], 1).(uint8)
	return nil
}

func (msg *OmciGetAllAlarmsTxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})
	// omciTrace.Contents["Result"] = nil
	omciTrace.Contents["AlarmRetrievalMode"] = msg.AlarmRetrievalMode
	return
}

func (msg *OmciGetAllAlarmsTxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciGetAllAlarmsTxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func (msg *OmciGetAllAlarmsRxMsg) Parse(def *SchemaDef, payload []byte) error {

	msg.Payload = append(msg.Payload, payload...)
	hl := msg.Header.decode(msg.Payload)
	msg.NumGetAllAlarmsCommands = bytesUInteger(msg.Payload[hl:hl+2], 2).(uint16)
	return nil
}

func (msg *OmciGetAllAlarmsRxMsg) OmciTrace(def *SchemaDef) (omciTrace OmciTrace) {
	msg.Header.populateInfo(def)
	omciTrace.payload = hex.Dump(msg.Payload)
	omciTrace.Header = msg.Header
	omciTrace.Contents = make(map[string]interface{})

	omciTrace.Contents["NumGetAllAlarmsCommands"] = msg.NumGetAllAlarmsCommands
	return
}

func (msg *OmciGetAllAlarmsRxMsg) OmciHeader() EngineOmciHeader {
	return msg.Header
}

func (msg *OmciGetAllAlarmsRxMsg) OmciResultReason() RR {
	return CmdProcessSuccessful
}

func ResultReasonString(resultReason RR) string {
	return commandReturnStringMap[resultReason]
}

type Attribute struct {
	Name            string
	Value           interface{}
	orderInMeSchema uint8
}
