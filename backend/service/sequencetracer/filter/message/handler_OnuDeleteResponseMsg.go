/*
# ------------------------------------------------------------
# -- OnuDeleteResponseMsg.go
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

type ProtoOnuDeleteResponseMsg struct {
	OnuDeleteResponse onuDeleteResponse `json:"onuDeleteResponse"`
}

const (
	KeyProtoOnuDeleteResponseMsg = `"onuDeleteResponse": {`
)

func (msg *ProtoOnuDeleteResponseMsg) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoOnuDeleteResponseMsg
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	notes := fmt.Sprintf("[ReturnCode:%s] [Vani:%d] [Channel:%d]", root.OnuDeleteResponse.ReturnCode, root.OnuDeleteResponse.VaniObjectIndex.Index, root.OnuDeleteResponse.ChannelObjectIndex.Index)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoOnuDeleteResponseMsg)
	msgHandlerRegist(KeyProtoOnuDeleteResponseMsg, msg)
}
