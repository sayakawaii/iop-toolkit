/*
# ------------------------------------------------------------
# -- meSACD.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	SACD = 150
)

type meSACD struct {
	relShip api.MeRelationship
}

func (me *meSACD) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)

	relship := me.relShip
	eType := relship.RelationshipType
	var eDir string
	eNote := []string(nil)

	var valuePointer uint64
	if v, err := attrs["TCP/UDP pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["TCP/UDP pointer"] = *edgeME
	}
	if v, err := attrs["Proxy server address pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[1]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["Proxy server address pointer"] = *edgeME
	}
	if v, err := attrs["Host part URI"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[1]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["Host part URI"] = *edgeME
	}
	if v, err := attrs["SIP registrar"]; err && v.(float64) != 65535 {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[2]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["SIP registrar"] = *edgeME
	}

	chart.Note = nil
	return *chart
}

func (me *meSACD) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meSACD)
	me.relShip = api.MeRelationship{ShortName: "SACD", AssociationType: []uint64{136, 157, 137}, RelationshipType: 1}
	api.MeRelationshipRegist(SACD, me)
	api.MeHandlerRegist(SACD, me)
}
