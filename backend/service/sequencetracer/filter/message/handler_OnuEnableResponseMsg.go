/*
# ------------------------------------------------------------
# -- OnuEnableResponseMsg.go
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

type ProtoOnuEnableResponseMsg struct {
	OnuEnableResponse onuEnableResponse `json:"onuEnableResponse"`
}

const (
	KeyProtoOnuEnableResponseMsg = `"onuEnableResponse": {`
)

func (msg *ProtoOnuEnableResponseMsg) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoOnuEnableResponseMsg
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	notes := fmt.Sprintf("[Channel:%d] [SN:%s]  [ReturnCode:%s] ", root.OnuEnableResponse.ChannelObjectIndex.Index, root.OnuEnableResponse.SerialNumber, root.OnuEnableResponse.ReturnCode)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoOnuEnableResponseMsg)
	msgHandlerRegist(KeyProtoOnuEnableResponseMsg, msg)
}
