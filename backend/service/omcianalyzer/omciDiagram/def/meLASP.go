/*
# ------------------------------------------------------------
# -- meLASP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	LASP = 455
)

type meLASP struct {
	relShip api.MeRelationship
}

func (me *meLASP) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var tpType, bridgePointer uint64
	if v, err := attrs["MAC bridge service profile ID pointer"]; err {
		bridgePointer = (uint64)(v.(float64))
	} else {
		return *chart
	}
	if v, err := attrs["Type"]; err {
		tpType = (uint64)(v.(float64))
	} else {
		return *chart
	}
	edgeBP := new(api.MeEdges) //edge point to bridge
	relship := me.relShip
	bridgeMeName := api.MeRelationshipDef[45].GetRelShip().ShortName
	eType := relship.RelationshipType
	eNote := []string(nil)
	var eDir string
	//bridge side edge
	if tpType == 1 {
		//UNI side MBPCD
		eDir = "up"
	} else {
		//ANI side MBPCD
		eDir = "down"
	}
	edgeBP.MeEdgesSet(eType, eDir, bridgeMeName, bridgePointer, eNote)
	chart.Edge["MAC bridge service profile ID pointer"] = *edgeBP
	chart.Note = nil
	return *chart
}

func (me *meLASP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meLASP)
	me.relShip = api.MeRelationship{ShortName: "LASP", AssociationType: []uint64{45}, RelationshipType: 1}
	api.MeRelationshipRegist(LASP, me)
	api.MeHandlerRegist(LASP, me)
}
