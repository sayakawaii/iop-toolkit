/*
# ------------------------------------------------------------
# -- meGemITP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	GemITP = 266
)

type meGemITP struct {
	relShip api.MeRelationship
}

func (me *meGemITP) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var tpPointer uint64
	//GemCTP
	if v, err := attrs["GEM port network CTP connectivity pointer"]; err {
		tpPointer = (uint64)(v.(float64))
	} else {
		return *chart
	}
	edgeTP := new(api.MeEdges)
	relship := me.relShip
	tpMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
	eType := relship.RelationshipType
	eNote := []string(nil)
	eDir := "up"
	edgeTP.MeEdgesSet(eType, eDir, tpMeName, tpPointer, eNote)
	chart.Edge["GEM port network CTP connectivity pointer"] = *edgeTP
	//GAL
	// if v, err := attrs["GAL profile pointer"]; err {
	// 	tpPointer = (uint64)(v.(float64))
	// } else {
	// 	return *chart
	// }
	// tpMeName = api.MeRelationshipDef[relship.AssociationType[1]].GetRelShip().ShortName
	// eDir = "auto"
	// edgeTP.MeEdgesSet(eType, eDir, tpMeName, tpPointer, eNote)
	// chart.Edge["GAL profile pointer"] = *edgeTP
	chart.Note = nil
	return *chart
}

func (me *meGemITP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meGemITP)
	me.relShip = api.MeRelationship{ShortName: "GemITP", AssociationType: []uint64{268, 272}, RelationshipType: 1}
	api.MeRelationshipRegist(GemITP, me)
	api.MeHandlerRegist(GemITP, me)
}
