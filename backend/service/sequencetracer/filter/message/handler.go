/*
# ------------------------------------------------------------
# -- OnuEnableResponseMsg.go
# --
# -- Huang Minghe
# -- 2025-3-17
# ------------------------------------------------------------
*/
package message

import "time"

// Schema represents the schema for an event
type Schema struct {
	From  string   `json:"from"`
	To    string   `json:"to"`
	Event []string `json:"event"`
	Notes string   `json:"notes"`
	Alias string   `json:"alias"`
}

type Meta struct {
	Time  time.Time
	From  string `json:"from"`
	To    string `json:"to"`
	Event string `json:"event"`
	Notes string `json:"notes"`
	Log   string `json:"alias"`
}

func (m *Meta) IsEmpty() bool {
	return m.Time.IsZero() && m.From == "" && m.To == "" && m.Event == "" && m.Notes == "" && m.Log == ""
}

type MsgHandler interface {
	Handle(jsonData string, timestamp time.Time, schema Schema, line string) Meta
}

var (
	MsgHandlerDef = make(map[string]MsgHandler)
)

func msgHandlerRegist(key string, handler MsgHandler) {
	MsgHandlerDef[key] = handler
}
func GetHandler(key string) (MsgHandler, bool) {
	h, ok := MsgHandlerDef[key]
	return h, ok
}
