/*
# ------------------------------------------------------------
# -- OnuDbruCapabilityInform.go
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

type ProtoOnuDbruCapabilityInform struct {
	OnuDbruCapabilityInform onuDbruCapabilityInform `json:"onuDbruCapabilityInform"`
}

const (
	KeyOnuDbruCapabilityInform = `"onuDbruCapabilityInform": {`
)

func (msg *ProtoOnuDbruCapabilityInform) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoOnuDbruCapabilityInform
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	notes := fmt.Sprintf("[Ct:%d][OnuId:%d][Dbru:%t]", root.OnuDbruCapabilityInform.CtObjectIndex.Index, root.OnuDbruCapabilityInform.OnuId, root.OnuDbruCapabilityInform.OnuDbruCapability)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoOnuDbruCapabilityInform)
	msgHandlerRegist(KeyOnuDbruCapabilityInform, msg)
}
