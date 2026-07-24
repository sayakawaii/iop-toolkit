/*
# ------------------------------------------------------------
# -- alarm.go
# --
# -- Huang Minghe
# -- 2025-3-17
# ------------------------------------------------------------
*/
package message

type alarmData struct {
	Action    string      `json:"action"`
	AlarmInfo []alarmInfo `json:"alarmInfo"`
	Label     string      `json:"label"`
}

type alarmInfo struct {
	AlarmType    int        `json:"alarmType"`
	Bitmap       int        `json:"bitmap"`
	ObjectID     int        `json:"objectId"`
	ObjectXpath  string     `json:"objectXpath"`
	EntityName   string     `json:"entityName"`
	EntityType   string     `json:"entityType"`
	AddInfo      *string    `json:"addInfo,omitempty"`
	Severity     []severity `json:"severity"`
	Timestamp    *int64     `json:"timestamp,omitempty"`
	ObjectType   *string    `json:"objectType,omitempty"`
	HierMetaData *string    `json:"hierMetaData,omitempty"`
}

type onuBulkAlarmMsg struct {
	State                  string                   `json:"state"`
	MsgOnuPloamDefectAlarm []msgOnuPloamDefectAlarm `json:"msgOnuPloamDefectAlarm"`
}

type msgOnuPloamDefectAlarm struct {
	VaniObjectIndex      vaniObjectIndex       `json:"vaniObjectIndex"`
	CtObjectIndex        ctObjectIndex         `json:"ctObjectIndex"`
	DefectStatus         int                   `json:"defectStatus"`
	DefectChanged        int                   `json:"defectChanged"`
	NotifIsOn            bool                  `json:"notifIsOn"`
	OperState            bool                  `json:"operState"`
	SerialNumber         string                `json:"serialNumber"`
	OnuDetectedAfterSwo  *bool                 `json:"onuDetectedAfterSwo,omitempty"`
	IsOnuOnline          bool                  `json:"isOnuOnline"`
	IsOnuOnlinePrev      bool                  `json:"isOnuOnlinePrev"`
	FiberDistance        float64               `json:"fiberDistance"`
	PrimaryCtObjectIndex *primaryCtObjectIndex `json:"primaryCtObjectIndex,omitempty"`
	DefectTimestamp      *int64                `json:"defectTimestamp,omitempty"`
	OnuUpdateMode        *string               `json:"onuUpdateMode,omitempty"`
}

// newOnuAlarmRaisedInform represents the newOnuAlarmRaisedInform field.
type newOnuAlarmRaisedInform struct {
	ChannelObjectIndex channelObjectIndex `json:"channelObjectIndex"`
	SerialNumber       string             `json:"serialNumber"`
}
