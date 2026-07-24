/*
# ------------------------------------------------------------
# -- meUniSupplemental1V2.go
# --
# -- Huang Minghe
# -- 2025-8-21
# ------------------------------------------------------------
*/
package ALCL

import (
	"omciAnalyzer/service/omcianalyzer/omciDiagram/api"
	"strconv"
)

// meUniSupplemental1V2 represents the ME for UniSupplemental1V2
// It is used to manage the supplemental information for UNI in version 2.

const (
	UniSupplemental1V2 = 65297
)

type meUniSupplemental1V2 struct {
	relShip api.MeRelationship
}

func (me *meUniSupplemental1V2) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)

	if v, ok := attrs["Multicast VID Value"]; ok {
		VID := (uint64)(v.(float64))
		chart.Note["VID"] = "VID: " + strconv.FormatUint(VID, 10)
	}
	if v, ok := attrs["Multicast P-Bit Value"]; ok {
		PBit := (uint64)(v.(float64))
		chart.Note["P-Bit"] = "P-Bit: " + strconv.FormatUint(PBit, 10)
	}
	if v, ok := attrs["IGMP Channel Bridge Port Number"]; ok {
		PortNumber := (uint64)(v.(float64))
		edgeTP := new(api.MeEdges)
		relship := me.relShip
		tpMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
		eType := relship.RelationshipType
		eNote := []string(nil)
		eDir := "auto"
		edgeTP.MeEdgesSet(eType, eDir, tpMeName, PortNumber-128, eNote)
		chart.Edge["IGMP Channel Bridge Port Number"] = *edgeTP
	}

	return *chart
}

func (me *meUniSupplemental1V2) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meUniSupplemental1V2)
	me.relShip = api.MeRelationship{ShortName: "UniSupplemental1V2", AssociationType: []uint64{47}, RelationshipType: 1}
	api.MeRelationshipRegist(UniSupplemental1V2, me)
	api.MeHandlerRegist(UniSupplemental1V2, me)
}
