/*
# ------------------------------------------------------------
# -- meVCTP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	VCTP = 139
)

type meVCTP struct {
	relShip api.MeRelationship
}

func (me *meVCTP) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)

	relship := me.relShip
	eType := relship.RelationshipType
	var eDir string
	eNote := []string(nil)

	var valuePointer uint64
	if v, err := attrs["User protocol pointer"]; err && api.VoipProtocol == api.Protocol_SIP {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "right"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["User protocol pointer"] = *edgeME
	}
	if v, err := attrs["User protocol pointer"]; err && api.VoipProtocol == api.Protocol_H248 {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "right"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[1]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["User protocol pointer"] = *edgeME
	}
	if v, err := attrs["PPTP pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "down"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[2]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["PPTP pointer"] = *edgeME
	}
	if v, err := attrs["VoIP media profile pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[3]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["VoIP media profile pointer"] = *edgeME
	}

	chart.Note = nil
	return *chart
}

func (me *meVCTP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meVCTP)
	me.relShip = api.MeRelationship{ShortName: "VCTP", AssociationType: []uint64{153, 155, 53, 142}, RelationshipType: 1}
	api.MeRelationshipRegist(VCTP, me)
	api.MeHandlerRegist(VCTP, me)
}
