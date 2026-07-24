/*
# ------------------------------------------------------------
# -- DmsCommitInform.go
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

type ProtoDmsCommitInform struct {
	DmsCommitInform dmsCommitInform `json:"dmsCommitInform"`
}

const (
	KeyProtoDmsCommitInform = `"dmsCommitInform": {`
)

func (msg *ProtoDmsCommitInform) Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta {
	var root ProtoDmsCommitInform
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		utils.Log("JSON decode error:", err)
		return Meta{}
	}
	notes := fmt.Sprintf("[State:%s]", root.DmsCommitInform.DmsCommitState)
	return Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
}

func init() {
	msg := new(ProtoDmsCommitInform)
	msgHandlerRegist(KeyProtoDmsCommitInform, msg)
}
