/*
# ------------------------------------------------------------
# -- meGemCTP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import (
	"omciAnalyzer/service/omcianalyzer/omciDiagram/api"
)

const (
	GemCTP = 268
)

type meGemCTP struct {
	relShip api.MeRelationship
}

func (me *meGemCTP) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var trafficPointer uint64
	// if v, err := attrs["T-CONT pointer"]; err {
	// 	tcontPointer = v
	// }else {
	// 	return *chart
	// }
	if v, err := attrs["Traffic management pointer for upstream"]; err {
		trafficPointer = (uint64)(v.(float64))
		if trafficPointer == 0 {
			return *chart
		}
	} else {
		return *chart
	}
	edgeTP := new(api.MeEdges)
	relship := me.relShip
	tpMeName := api.MeRelationshipDef[relship.AssociationType[1]].GetRelShip().ShortName
	eType := relship.RelationshipType
	eNote := []string(nil)
	eDir := "up"
	edgeTP.MeEdgesSet(eType, eDir, tpMeName, trafficPointer, eNote)
	chart.Edge["Traffic management pointer for upstream"] = *edgeTP
	var portID uint64
	if v, err := attrs["Port ID"]; err {
		portID = (uint64)(v.(float64))
		if portID == 0 {
			return *chart
		}
	} else {
		return *chart
	}
	chart.Note = nil
	chart.SetAttr("Port ID", portID)
	return *chart
}

func (me *meGemCTP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meGemCTP)
	me.relShip = api.MeRelationship{ShortName: "GemCTP", AssociationType: []uint64{262, 277}, RelationshipType: 1}
	api.MeRelationshipRegist(GemCTP, me)
	api.MeHandlerRegist(GemCTP, me)
}
