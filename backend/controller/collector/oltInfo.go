package collector

type State struct {
	StateLastChanged string `json:"state_last_changed"`
	AdminState       string `json:"admin_state"`
	OperState        string `json:"oper_state"`
	StandbyState     string `json:"standby_state"`
}

type OnuPresentt struct {
	OnuFiberDistance       int    `json:"onu_fiber_distance"`
	DetectedSerialNumber   string `json:"detected_serial_number"`
	DetectedRegistrationId string `json:"detected_registration_id"`
}

type Vani struct {
	OnuId                  int         `json:"onu_id"`
	ManagementTcontAllocId int         `json:"management_tcont_alloc_id"`
	ManagementGemportId    int         `json:"management_gemport_id"`
	OnuPresenceState       string      `json:"onu_presence_state"`
	OnuPresentt            OnuPresentt `json:"onu_presentt"`
}

type ONUInfo struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	AdminState   string `json:"admin_state"`
	OperState    string `json:"oper_state"`
	LastChanged  string `json:"last_changed"`
	IfIndex      string `json:"if_index,omitempty"`
	LowerLayerIf string `json:"lower_layer_if,omitempty"`
	Vani         Vani   `json:"vani"`
}

type Component struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Parent string `json:"parent,omitempty"`
	Model  string `json:"model,omitempty"`
	State  State  `json:"state,omitempty"`
}

type PONInfo struct {
	Component
	ONUInfo []ONUInfo `json:"onuinfo,omitempty"`
}

type LTInfo struct {
	Component
	PONInfo []PONInfo `json:"poninfo,omitempty"`
}

type NTInfo struct {
	Component
	LTInfo []LTInfo `json:"ltinfo,omitempty"`
}

type RespOltTopology struct {
	RequestID string `json:"request_id"`
	NTInfo    NTInfo `json:"ntinfo"`
}

type OltInfo struct {
	OamIP  string `json:"oam_ip"`
	NTInfo NTInfo `json:"ntinfo"`
}

func (info *OltInfo) GetTopologyResponse(requestID string) RespOltTopology {
	return RespOltTopology{
		RequestID: requestID,
		NTInfo:    info.NTInfo,
	}
}
