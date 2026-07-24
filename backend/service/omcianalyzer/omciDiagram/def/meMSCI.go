/*
# ------------------------------------------------------------
# -- meMSCI.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	MSCI = 310
)

type meMSCI struct {
	relShip api.MeRelationship
}

func (me *meMSCI) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var tpPointer uint64
	if v, err := attrs["Multicast operations profile pointer"]; err {
		tpPointer = (uint64)(v.(float64))
		if tpPointer == 0 {
			return *chart
		}
	} else {
		return *chart
	}
	edgeTP := new(api.MeEdges)
	relship := me.relShip
	tpMeName := api.MeRelationshipDef[309].GetRelShip().ShortName
	eType := api.MeRelationshipDef[309].GetRelShip().RelationshipType
	eNote := []string(nil)
	eDir := "down"
	edgeTP.MeEdgesSet(eType, eDir, tpMeName, tpPointer, eNote)
	chart.Edge["Multicast operations profile pointer"] = *edgeTP

	// edgeTP := new(api.MeEdges)
	// relship := me.relShip
	tpMeName = api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
	eType = relship.RelationshipType
	eNote = []string(nil)
	eDir = "right"
	tpPointer = node.MeInst
	edgeTP.MeEdgesSet(eType, eDir, tpMeName, tpPointer, eNote)
	chart.Edge["TP pointer"] = *edgeTP

	chart.Note = nil
	return *chart
}

func (me *meMSCI) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meMSCI)
	me.relShip = api.MeRelationship{ShortName: "MSCI", AssociationType: []uint64{47, 130}, RelationshipType: 0}
	api.MeRelationshipRegist(MSCI, me)
	api.MeHandlerRegist(MSCI, me)
}
