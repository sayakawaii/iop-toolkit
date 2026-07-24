/*
# ------------------------------------------------------------
# -- index.go
# --
# -- Huang Minghe
# -- 2025-3-17
# ------------------------------------------------------------
*/
package message

type severity struct {
	Position int    `json:"position"`
	Severity string `json:"severity"`
}

// CtObjectIndex represents the ctObjectIndex field.
type ctObjectIndex struct {
	Index int `json:"index"`
}

// VaniObjectIndex represents the vaniObjectIndex field.
type vaniObjectIndex struct {
	Index int `json:"index"`
}

// ChannelObjectIndex represents the channelObjectIndex field.
type channelObjectIndex struct {
	Index int `json:"index"`
}

type primaryCtObjectIndex struct {
	Index int `json:"index"`
}
