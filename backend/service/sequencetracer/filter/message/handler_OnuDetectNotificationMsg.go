/*
# ------------------------------------------------------------
# -- OnuDetectNotificationMsg.go
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

type ProtoOnuDetectNotificationMsg struct {
	OnuDetectNotificationMsg onuDetectNotificationMsg `json:"onuDetectNotificationMsg"`
}

const (
	KeyProtoOnuDetectNotificationMsg = `"onuDetectNotificationMsg": {`
)

func (msg *ProtoOnuDetectNotificationMsg) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoOnuDetectNotificationMsg
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	utils.Log("msg:", jsonData)
	utils.Log("root:", root)
	notes := fmt.Sprintf("[SN:%s] [OnuId:%d]", root.OnuDetectNotificationMsg.SerialNumber, root.OnuDetectNotificationMsg.OnuId)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoOnuDetectNotificationMsg)
	msgHandlerRegist(KeyProtoOnuDetectNotificationMsg, msg)
}
