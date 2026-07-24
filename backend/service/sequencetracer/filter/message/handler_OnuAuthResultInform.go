/*
# ------------------------------------------------------------
# -- OnuAuthResultInform.go
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

type ProtoOnuAuthResultInform struct {
	OnuAuthResultInform onuAuthResultInform `json:"onuAuthResultInform"`
}

const (
	KeyProtoOnuAuthResultInform = `"onuAuthResultInform": {`
)

func (msg *ProtoOnuAuthResultInform) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoOnuAuthResultInform
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	notes := fmt.Sprintf("[SN:%s] [ChannelObjIndex:%d]", root.OnuAuthResultInform.SerialNumber, root.OnuAuthResultInform.ChannelObjectIndex.Index)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoOnuAuthResultInform)
	msgHandlerRegist(KeyProtoOnuAuthResultInform, msg)
}
