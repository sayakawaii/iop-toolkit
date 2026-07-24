/*
# ------------------------------------------------------------
# -- OnuDetectCompletedNotification.go
# --
# -- Huang Minghe
# -- 2025-3-17
# ------------------------------------------------------------
*/
package message

import (
	"encoding/json"
	"fmt"
	"omciAnalyzer/utils"
	"time"
)

type ProtoOnuDetectCompletedNotification struct {
	OnuDetectCompletedNotification onuDetectCompletedNotification `json:"onuDetectCompletedNotification"`
}

const (
	KeyOnuDetectCompletedNotification = `"onuDetectCompletedNotification": {`
)

func (msg *ProtoOnuDetectCompletedNotification) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoOnuDetectCompletedNotification
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	notes := fmt.Sprintf("[Vani:%d]", root.OnuDetectCompletedNotification.VaniObjectIndex.Index)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoOnuDetectCompletedNotification)
	msgHandlerRegist(KeyOnuDetectCompletedNotification, msg)
}
