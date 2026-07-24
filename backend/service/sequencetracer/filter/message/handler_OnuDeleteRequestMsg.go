/*
# ------------------------------------------------------------
# -- OnuDeleteRequestMsg.go
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

type ProtoOnuDeleteRequestMsg struct {
	OnuDeleteRequest onuDeleteRequest `json:"onuDeleteRequest"`
}

const (
	KeyProtoOnuDeleteRequestMsg = `"onuDeleteRequest": {`
)

func (msg *ProtoOnuDeleteRequestMsg) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoOnuDeleteRequestMsg
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	notes := fmt.Sprintf("[Vani:%d] [Channel:%d]", root.OnuDeleteRequest.VaniObjectIndex.Index, root.OnuDeleteRequest.ChannelObjectIndex.Index)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoOnuDeleteRequestMsg)
	msgHandlerRegist(KeyProtoOnuDeleteRequestMsg, msg)
}
