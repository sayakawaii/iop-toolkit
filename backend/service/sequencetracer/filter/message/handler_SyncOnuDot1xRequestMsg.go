/*
# ------------------------------------------------------------
# -- SyncOnuDot1xRequestMsg.go
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

type ProtoSyncOnuDot1xRequestMsg struct {
	SyncOnuDot1xRequest syncOnuDot1xRequest `json:"syncOnuDot1xRequest"`
}

const (
	KeySyncOnuDot1xRequestMsg = `"syncOnuDot1xRequest": {`
)

func (msg *ProtoSyncOnuDot1xRequestMsg) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoSyncOnuDot1xRequestMsg
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	notes := fmt.Sprintf("[OnuName:%s]", root.SyncOnuDot1xRequest.OnuName)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoSyncOnuDot1xRequestMsg)
	msgHandlerRegist(KeySyncOnuDot1xRequestMsg, msg)
}
