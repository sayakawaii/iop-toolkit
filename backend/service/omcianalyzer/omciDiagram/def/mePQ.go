/*
# ------------------------------------------------------------
# -- mePQ.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	PQ = 277
)

type mePQ struct {
	relShip api.MeRelationship
}

func (me *mePQ) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var trafficPointer uint64
	if v, err := attrs["Traffic scheduler pointer"]; err {
		trafficPointer = (uint64)(v.(float64))
		if trafficPointer == 0 {
			return *chart
		}
	} else {
		return *chart
	}
	edgeTP := new(api.MeEdges)
	relship := me.relShip
	tpMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
	eType := relship.RelationshipType
	eNote := []string(nil)
	eDir := "up"
	edgeTP.MeEdgesSet(eType, eDir, tpMeName, trafficPointer, eNote)
	chart.Edge["Traffic scheduler pointer"] = *edgeTP
	chart.Note = nil
	return *chart
}

func (me *mePQ) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(mePQ)
	me.relShip = api.MeRelationship{ShortName: "PQ", AssociationType: []uint64{278, 262}, RelationshipType: 1}
	api.MeRelationshipRegist(PQ, me)
	api.MeHandlerRegist(PQ, me)
}
