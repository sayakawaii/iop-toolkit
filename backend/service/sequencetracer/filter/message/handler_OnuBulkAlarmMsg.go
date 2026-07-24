/*
# ------------------------------------------------------------
# -- OnuBulkAlarmMsg.go
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

type ProtoOnuBulkAlarmMsg struct {
	OnuBulkAlarmMsg onuBulkAlarmMsg `json:"onuBulkAlarmMsg"`
}

const (
	KeyProtoOnuBulkAlarmMsg = `"onuBulkAlarmMsg": {`
)

func (msg *ProtoOnuBulkAlarmMsg) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoOnuBulkAlarmMsg
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	notes := ""
	for _, alarm := range root.OnuBulkAlarmMsg.MsgOnuPloamDefectAlarm {
		notes += fmt.Sprintf("[State:%s] [Vani:%d] [Ct:%d] [SN:%s]\n", root.OnuBulkAlarmMsg.State,
			alarm.VaniObjectIndex.Index,
			alarm.CtObjectIndex.Index,
			alarm.SerialNumber)
	}
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoOnuBulkAlarmMsg)
	msgHandlerRegist(KeyProtoOnuBulkAlarmMsg, msg)
}
