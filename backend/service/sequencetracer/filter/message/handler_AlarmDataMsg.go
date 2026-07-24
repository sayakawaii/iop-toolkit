/*
# ------------------------------------------------------------
# -- AlarmDataMsg.go
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

type ProtoAlarmDataMsg struct {
	AlarmData alarmData `json:"alarmData"`
}

const (
	KeyProtoAlarmDataMsg = `"alarmData": {`
)

func (msg *ProtoAlarmDataMsg) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoAlarmDataMsg
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	// utils.Log("root:", root)
	notes := fmt.Sprintf("[SN:%s] [Type:%d]", root.AlarmData.AlarmInfo[0].EntityName, root.AlarmData.AlarmInfo[0].AlarmType)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoAlarmDataMsg)
	msgHandlerRegist(KeyProtoAlarmDataMsg, msg)
}
