/*
# ------------------------------------------------------------
# -- meTS.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	TS = 278
)

type meTS struct {
	relShip api.MeRelationship
}

func (me *meTS) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var tcontPointer uint64
	relship := me.relShip
	tpMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
	if v, err := attrs["T-CONT pointer"]; err {
		//Traffic Scheduler ponint T-CONT
		tcontPointer = (uint64)(v.(float64))
	} else if v, err := attrs["Traffic scheduler pointer"]; err {
		//Traffic Scheduler ponint another Traffic Scheduler
		tcontPointer = (uint64)(v.(float64))
		tpMeName = api.MeRelationshipDef[relship.AssociationType[1]].GetRelShip().ShortName
	} else {
		//here is default, Traffic Scheduler ponint T-CONT
		tcontPointer = node.MeInst
	}
	edgeTP := new(api.MeEdges)
	eType := relship.RelationshipType
	eNote := []string(nil)
	eDir := "up"
	edgeTP.MeEdgesSet(eType, eDir, tpMeName, tcontPointer, eNote)
	chart.Edge["T-CONT pointer"] = *edgeTP
	chart.Note = nil
	return *chart
}

func (me *meTS) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meTS)
	me.relShip = api.MeRelationship{ShortName: "TS", AssociationType: []uint64{262, 278}, RelationshipType: 1}
	api.MeRelationshipRegist(TS, me)
	api.MeHandlerRegist(TS, me)
}
