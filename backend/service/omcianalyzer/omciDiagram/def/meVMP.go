/*
# ------------------------------------------------------------
# -- meVMP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	VMP = 142
)

type meVMP struct {
	relShip api.MeRelationship
}

func (me *meVMP) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)

	relship := me.relShip
	eType := relship.RelationshipType
	var eDir string
	eNote := []string(nil)

	var valuePointer uint64
	if v, err := attrs["Voice service profile pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["Voice service profile pointer"] = *edgeME
	}
	if v, err := attrs["RTP profile pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[1]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["RTP profile pointer"] = *edgeME
	}

	chart.Note = nil
	return *chart
}

func (me *meVMP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meVMP)
	me.relShip = api.MeRelationship{ShortName: "VMP", AssociationType: []uint64{58, 143}, RelationshipType: 1}
	api.MeRelationshipRegist(VMP, me)
	api.MeHandlerRegist(VMP, me)
}
