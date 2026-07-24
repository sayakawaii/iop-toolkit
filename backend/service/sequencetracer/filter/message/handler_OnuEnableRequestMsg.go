/*
# ------------------------------------------------------------
# -- OnuEnableRequestMsg.go
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

type ProtoOnuEnableRequestMsg struct {
	OnuEnableRequest onuEnableRequest `json:"onuEnableRequest"`
}

const (
	KeyProtoOnuEnableRequestMsg = `"onuEnableRequest": {`
)

func (msg *ProtoOnuEnableRequestMsg) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoOnuEnableRequestMsg
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	notes := fmt.Sprintf("[Channel:%d] [SN:%s]  [EnableType:%s] ", root.OnuEnableRequest.ChannelObjectIndex.Index, root.OnuEnableRequest.SerialNumber, root.OnuEnableRequest.OnuEnableType)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoOnuEnableRequestMsg)
	msgHandlerRegist(KeyProtoOnuEnableRequestMsg, msg)
}
