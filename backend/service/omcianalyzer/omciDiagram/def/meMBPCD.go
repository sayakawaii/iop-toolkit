/*
# ------------------------------------------------------------
# -- meMBPCD.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	MBPCD = 47
)

type meMBPCD struct {
	relShip api.MeRelationship
}

func (me *meMBPCD) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var tpType, bridgePointer, tpPointer, laspPointer uint64
	if v, err := attrs["Bridge id pointer"]; err {
		bridgePointer = (uint64)(v.(float64))
	} else {
		return *chart
	}
	if v, err := attrs["TP type"]; err {
		tpType = (uint64)(v.(float64))
	} else {
		return *chart
	}
	if v, err := attrs["TP pointer"]; err {
		tpPointer = (uint64)(v.(float64))
	} else {
		return *chart
	}
	laspPointer = 0
	if v, err := attrs["LASP ID pointer"]; err {
		laspPointer = (uint64)(v.(float64))
	}
	edgeBP := new(api.MeEdges) //edge point to bridge
	edgeTP := new(api.MeEdges) //edge point to UNI or ANI
	relship := me.relShip
	bridgeMeName := api.MeRelationshipDef[45].GetRelShip().ShortName
	tpMeName := api.MeRelationshipDef[relship.AssociationType[tpType]].GetRelShip().ShortName
	eType := relship.RelationshipType
	eNote := []string(nil)
	var eDir string
	//bridge side edge
	if tpType == 1 || tpType == 4 || tpType == 11 {
		//UNI side MBPCD
		eDir = "up"
	} else {
		//ANI side MBPCD
		eDir = "down"
	}
	edgeBP.MeEdgesSet(eType, eDir, bridgeMeName, bridgePointer, eNote)
	//handle LASP edge
	if laspPointer != 0 {
		edgeLASP := new(api.MeEdges) //edge point to LASP
		laspMeName := api.MeRelationshipDef[455].GetRelShip().ShortName
		edgeLASP.MeEdgesSet(eType, eDir, laspMeName, laspPointer, eNote)
		chart.Edge["LASP ID pointer"] = *edgeLASP
	}
	//UNI or ANI side edge
	if tpType == 1 || tpType == 4 || tpType == 11 {
		//UNI side MBPCD
		eDir = "down"
	} else {
		//ANI side MBPCD
		eDir = "up"
	}
	edgeTP.MeEdgesSet(eType, eDir, tpMeName, tpPointer, eNote)
	chart.Edge["Bridge id pointer"] = *edgeBP
	chart.Edge["TP pointer"] = *edgeTP
	chart.Note = nil
	return *chart
}

func (me *meMBPCD) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meMBPCD)
	me.relShip = api.MeRelationship{ShortName: "MBPCD", AssociationType: []uint64{0, 11, 14, 130, 134, 266, 281, 98, 117, 286, 0, 329, 162}, RelationshipType: 1}
	api.MeRelationshipRegist(MBPCD, me)
	api.MeHandlerRegist(MBPCD, me)
}
