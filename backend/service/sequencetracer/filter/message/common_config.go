/*
# ------------------------------------------------------------
# -- config.go
# --
# -- Huang Minghe
# -- 2025-3-17
# ------------------------------------------------------------
*/
package message

type onuDeleteRequest struct {
	VaniObjectIndex    vaniObjectIndex    `json:"vaniObjectIndex"`
	ChannelObjectIndex channelObjectIndex `json:"channelObjectIndex"`
}

type onuDeleteResponse struct {
	VaniObjectIndex    vaniObjectIndex    `json:"vaniObjectIndex"`
	ChannelObjectIndex channelObjectIndex `json:"channelObjectIndex"`
	ReturnCode         string             `json:"returnCode"`
}

type onuEnableRequest struct {
	ChannelObjectIndex channelObjectIndex `json:"channelObjectIndex"`
	SerialNumber       string             `json:"serialNumber"`
	OnuEnableType      string             `json:"onuEnableType"`
}

type onuEnableResponse struct {
	ChannelObjectIndex channelObjectIndex `json:"channelObjectIndex"`
	SerialNumber       string             `json:"serialNumber"`
	ReturnCode         string             `json:"returnCode"`
}

// onuAuthResultInform represents the onuAuthResultInform field.
type onuAuthResultInform struct {
	ChannelObjectIndex channelObjectIndex `json:"channelObjectIndex"`
	SerialNumber       string             `json:"serialNumber"`
}

type dmsCommitInform struct {
	DmsCommitState string `json:"dmsCommitState"`
}

type syncOnuDot1xRequest struct {
	OnuName string `json:"onuName"`
}
