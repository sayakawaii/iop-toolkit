/*
# ------------------------------------------------------------
# -- notif.go
# --
# -- Huang Minghe
# -- 2025-3-17
# ------------------------------------------------------------
*/
package message

type vaniConnectStatusNotification struct {
	CtObjectIndex    ctObjectIndex   `json:"ctObjectIndex"`
	OnuId            int             `json:"onuId"`
	ConnStatus       string          `json:"status"`
	OnuName          string          `json:"onuName"`
	VaniObjectIndex  vaniObjectIndex `json:"vaniObjectIndex"`
	DetectReason     string          `json:"detectReason"`
	DetectOntPonType *string         `json:"detectOntPonType,omitempty"`
}

// OnuDetectNotificationMsg represents the onuDetectNotificationMsg field.
type onuDetectNotificationMsg struct {
	CtObjectIndex          ctObjectIndex   `json:"ctObjectIndex"`
	VaniObjectIndex        vaniObjectIndex `json:"vaniObjectIndex"`
	NotifIsOn              bool            `json:"notifIsOn"`
	SerialNumber           string          `json:"serialNumber"`
	OnuId                  int             `json:"onuId"`
	RangingState           int             `json:"rangingState"`
	DetectedRegistrationId string          `json:"detectedRegistrationId"`
	DetectedUpstreamRate   int             `json:"detectedUpstreamRate"`
	FiberDistance          float64         `json:"fiberDistance"`
	MultiVani              bool            `json:"multiVani"`
	OnuStatus              int             `json:"onuStatus"`
	OnuAid                 int             `json:"onuAid"`
}

type onuDetectCompletedNotification struct {
	VaniObjectIndex vaniObjectIndex `json:"vaniObjectIndex"`
}

type onuDbruCapabilityInform struct {
	CtObjectIndex     ctObjectIndex `json:"ctObjectIndex"`
	OnuId             int           `json:"onuId"`
	OnuDbruCapability bool          `json:"onuDbruCapability"`
}
