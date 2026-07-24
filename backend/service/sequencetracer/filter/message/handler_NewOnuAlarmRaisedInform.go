/*
# ------------------------------------------------------------
# -- NewOnuAlarmRaisedInform.go
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

type ProtoNewOnuAlarmRaisedInform struct {
	NewOnuAlarmRaisedInform newOnuAlarmRaisedInform `json:"newOnuAlarmRaisedInform"`
}

const (
	KeyProtoNewOnuAlarmRaisedInform = `"newOnuAlarmRaisedInform": {`
)

func (msg *ProtoNewOnuAlarmRaisedInform) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoNewOnuAlarmRaisedInform
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	notes := fmt.Sprintf("[SN:%s] [ChannelObjIndex:%d]", root.NewOnuAlarmRaisedInform.SerialNumber, root.NewOnuAlarmRaisedInform.ChannelObjectIndex.Index)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoNewOnuAlarmRaisedInform)
	msgHandlerRegist(KeyProtoNewOnuAlarmRaisedInform, msg)
}
