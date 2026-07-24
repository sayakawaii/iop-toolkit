/*
# ------------------------------------------------------------
# -- meRS232PMHD.go
# --
# -- Huang Minghe
# -- 2026-3-11
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	RS232PMHD = 403
)

type meRS232PMHD struct {
	relShip api.MeRelationship
}

func (me *meRS232PMHD) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)

	relShip := me.relShip
	edgeME := new(api.MeEdges)
	pointerMeName := api.MeRelationshipDef[relShip.AssociationType[0]].GetRelShip().ShortName
	edgeME.MeEdgesSet(relShip.RelationshipType, "up", pointerMeName, node.MeInst, []string(nil))
	chart.Edge["PPTP RS232/RS485 UNI"] = *edgeME
	chart.Note = nil

	return *chart
}

func (me *meRS232PMHD) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meRS232PMHD)
	me.relShip = api.MeRelationship{ShortName: "RS232PMHD", AssociationType: []uint64{401}, RelationshipType: 0}
	api.MeRelationshipRegist(RS232PMHD, me)
	api.MeHandlerRegist(RS232PMHD, me)
}
