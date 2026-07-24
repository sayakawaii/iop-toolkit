/*
# ------------------------------------------------------------
# -- OnuConnStatus.go
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

type ProtoOnuConnStatus struct {
	ObjType         string                        `json:"objType"`
	ConnStatusNotif vaniConnectStatusNotification `json:"vaniConnectStatusNotification"`
}

const (
	KeyProtoOnuConnStatus = `"vaniConnectStatusNotification": {`
)

func (msg *ProtoOnuConnStatus) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoOnuConnStatus
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	notes := fmt.Sprintf("[State:%s] [ONU: %d] [OnuName: %s] [Reason: %s]", root.ConnStatusNotif.ConnStatus, root.ConnStatusNotif.OnuId, root.ConnStatusNotif.OnuName, root.ConnStatusNotif.DetectReason)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoOnuConnStatus)
	msgHandlerRegist(KeyProtoOnuConnStatus, msg)
}
