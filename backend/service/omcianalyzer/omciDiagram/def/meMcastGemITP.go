/*
# ------------------------------------------------------------
# -- meMcastGemITP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	McastGemITP = 281
)

type meMcastGemITP struct {
	relShip api.MeRelationship
}

func (me *meMcastGemITP) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var tpPointer uint64
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
	chart.Note = nil
	return *chart
}

func (me *meMcastGemITP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meMcastGemITP)
	me.relShip = api.MeRelationship{ShortName: "McastGemITP", AssociationType: []uint64{268}, RelationshipType: 1}
	api.MeRelationshipRegist(McastGemITP, me)
	api.MeHandlerRegist(McastGemITP, me)
}
