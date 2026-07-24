/*
# ------------------------------------------------------------
# -- meNetAddr.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	NetAddr = 137
)

type meNetAddr struct {
	relShip api.MeRelationship
}

func (me *meNetAddr) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)

	relship := me.relShip
	eType := relship.RelationshipType
	var eDir string
	eNote := []string(nil)

	var valuePointer uint64
	if v, err := attrs["Security pointer"]; err && v.(float64) != 65535 {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["Security pointer"] = *edgeME
	}
	if v, err := attrs["Address pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[1]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["Address pointer"] = *edgeME
	}

	chart.Note = nil
	return *chart
}

func (me *meNetAddr) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meNetAddr)
	me.relShip = api.MeRelationship{ShortName: "NetAddr", AssociationType: []uint64{148, 157}, RelationshipType: 1}
	api.MeRelationshipRegist(NetAddr, me)
	api.MeHandlerRegist(NetAddr, me)
}
