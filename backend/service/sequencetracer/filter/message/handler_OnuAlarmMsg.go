/*
# ------------------------------------------------------------
# -- OnuAlarmMsg.go
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

type ProtoOnuAlarmMsg struct {
	OnuAlarmMsg msgOnuPloamDefectAlarm `json:"alarmData"`
}

const (
	KeyProtoOnuAlarmMsg = `"onuAlarmMsg": {`
)

func (msg *ProtoOnuAlarmMsg) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoOnuAlarmMsg
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	notes := fmt.Sprintf("[SN:%s] [Vani:%d] [Channel:%d]", root.OnuAlarmMsg.SerialNumber, root.OnuAlarmMsg.VaniObjectIndex.Index, root.OnuAlarmMsg.CtObjectIndex.Index)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoOnuAlarmMsg)
	msgHandlerRegist(KeyProtoOnuAlarmMsg, msg)
}
