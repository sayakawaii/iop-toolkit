/*
# ------------------------------------------------------------
# -- meVCD.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	VCD = 138
)

type meVCD struct {
	relShip api.MeRelationship
}

func (me *meVCD) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)

	relship := me.relShip
	eType := relship.RelationshipType
	var eDir string
	eNote := []string(nil)

	var valuePointer uint64
	if v, err := attrs["VoIP configuration address pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["VoIP configuration address pointer"] = *edgeME
	}

	if v, err := attrs["Signalling protocol used"]; err {
		api.VoipProtocol = (uint64)(v.(float64))
	}

	chart.Note = nil
	return *chart
}

func (me *meVCD) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meVCD)
	me.relShip = api.MeRelationship{ShortName: "VCD", AssociationType: []uint64{137}, RelationshipType: 1}
	api.MeRelationshipRegist(VCD, me)
	api.MeHandlerRegist(VCD, me)
}
