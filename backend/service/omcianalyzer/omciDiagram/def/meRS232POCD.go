/*
# ------------------------------------------------------------
# -- meRS232POCD.go
# --
# -- Huang Minghe
# -- 2026-3-11
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	RS232POCD = 402
)

type meRS232POCD struct {
	relShip api.MeRelationship
}

func (me *meRS232POCD) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)

	relShip := me.relShip
	edgeType := relShip.RelationshipType
	edgeNote := []string(nil)

	if v, ok := attrs["TCP/UDP pointer"]; ok {
		pointer := uint64(v.(float64))
		if pointer != 65535 {
			edgeME := new(api.MeEdges)
			pointerMeName := api.MeRelationshipDef[relShip.AssociationType[0]].GetRelShip().ShortName
			edgeME.MeEdgesSet(edgeType, "up", pointerMeName, pointer, edgeNote)
			chart.Edge["TCP/UDP pointer"] = *edgeME
		}
	}

	if v, ok := attrs["PPTP pointer"]; ok {
		pointer := uint64(v.(float64))
		edgeME := new(api.MeEdges)
		pointerMeName := api.MeRelationshipDef[relShip.AssociationType[1]].GetRelShip().ShortName
		edgeME.MeEdgesSet(edgeType, "down", pointerMeName, pointer, edgeNote)
		chart.Edge["PPTP pointer"] = *edgeME
	}

	chart.Note = nil
	return *chart
}

func (me *meRS232POCD) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meRS232POCD)
	me.relShip = api.MeRelationship{ShortName: "RS232POCD", AssociationType: []uint64{136, 401}, RelationshipType: 1}
	api.MeRelationshipRegist(RS232POCD, me)
	api.MeHandlerRegist(RS232POCD, me)
}
