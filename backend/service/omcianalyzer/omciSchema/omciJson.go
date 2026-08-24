/*
# ------------------------------------------------------------
# -- omciJson.go
# --
# -- Huang Minghe
# -- 2022-6-30
# ------------------------------------------------------------
*/

package omciSchema

import (
	"encoding/json"
)

type OmciHeader struct {
	Tcid        uint64 `json:"Tcid"`
	AR          byte   `json:"AR"`
	AK          byte   `json:"AK"`
	MsgType     byte   `json:"MsgType"`
	MsgTypeName string `json:"MsgTypeName"`
	DevId       byte   `json:"DevId"`
	MsgFormat   string `json:"MsgFormat"`
	MeClass     uint64 `json:"MeClass"`
	MeClassName string `json:"MeClassName"`
	MeInst      uint64 `json:"MeInst"`
}

type Attr struct {
	Name  string `json:"Name"`
	Value any    `json:"Value"`
}

type OmciArContents struct {
	AttrMask uint64 `json:"AttrMask"`
	Attrs    []Attr `json:"Attrs"`
}

type OmciAlarmContents struct {
	AlarmSeqNum      uint64 `json:"AlarmSeqNum"`
	AlarmName        string `json:"AlarmName"`
	AlarmDescription string `json:"AlarmDescription"`
}

type OmciAkContents struct {
	Result   byte   `json:"Result"`
	AttrMask uint64 `json:"AttrMask"`
	Attrs    []Attr `json:"Attrs"`
}

type OmciCreateAkContents struct {
	Result       byte   `json:"Result"`
	AttrExecMask uint64 `json:"AttrExecMask"` // Attribute Execution Mask for Create Response Message used with 0011 encoding
}

type OmciDeleteAkContents struct {
	Result byte `json:"Result"`
}

type OmciSetAkContents struct {
	Result       byte   `json:"Result"`
	OptAttrMask  uint64 `json:"OptAttrMask"`  // Optional Attribute Mask for Set Response Message used with 1001 encoding
	AttrExecMask uint64 `json:"AttrExecMask"` // Attribute Mask for Set Response Message used with 1001 encoding
}

type OmciGetAkContents struct {
	Result       byte   `json:"Result"`
	AttrMask     uint64 `json:"AttrMask"`
	Attrs        []Attr `json:"Attrs"`
	OptAttrMask  uint64 `json:"OptAttrMask"`  // Optional Attribute Mask for Set Response Message used with 1001 encoding
	AttrExecMask uint64 `json:"AttrExecMask"` // Attribute Mask for Set Response Message used with 1001 encoding
}

type OmciMibuploadAkContents struct {
	NumMIBUploadCommands uint64 `json:"NumMIBUploadCommands"`
}

type OmciMibuploadNextArContents struct {
	CommandSequenceNum uint64 `json:"CommandSequenceNum"`
}

type OmciMibuploadNextAkContents struct {
	UploadMeClass    uint64 `json:"MeClass"`
	UploadMeName     string `json:"MeClassName"`
	UploadMeInstance uint64 `json:"MeInst"`
	AttrMask         uint64 `json:"AttrMask"`
	// Attrs            []map[string]string `json:"Attrs"`
	Attrs []Attr `json:"Attrs"`
}

type OmciTestArContents struct {
	Payload          string `json:"payload"`
	TestRequestAttrs []Attr `json:"TestRequestAttrs"`
}

type OmciTestAkContents struct {
	Result byte `json:"Result"`
}

type OmciTestResultContents struct {
	UploadMeClass    uint64 `json:"MeClass"`
	UploadMeInstance uint64 `json:"MeInst"`
	TestResults      []Attr `json:"TestResults"`
	Contents         string `json:"Contents"`
}

type OmciContext struct {
	Header   OmciHeader `json:"header"`
	Contents any        `json:"contents"`
}

func GetContentFromInterface(src any, dest any) (e error) {
	arr, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(arr, dest)
	if err != nil {
		return err
	}
	return nil
}

func SliceMapConvToMap(src []Attr) map[string]any {
	dest := make(map[string]any)
	for index := range src {
		attr := src[index]
		dest[attr.Name] = attr.Value
	}
	return dest
}
