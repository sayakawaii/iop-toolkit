/*
# ------------------------------------------------------------
# -- meMgcCD.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	MgcCD = 155
)

type meMgcCD struct {
	relShip api.MeRelationship
}

func (me *meMgcCD) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)

	relship := me.relShip
	eType := relship.RelationshipType
	var eDir string
	eNote := []string(nil)

	var valuePointer uint64
	if v, err := attrs["Primary MGC"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges) //edge point to a network address that contains the name of the primary MGC
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["Primary MGC"] = *edgeME
	}
	if v, err := attrs["Secondary MGC"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges) //edge point to a network address that contains the name of the secondary MGC
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["Secondary MGC"] = *edgeME
	}
	if v, err := attrs["TCP/UDP pointer"]; err && v.(float64) != 65535 {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges) //edge point to the TCP/UDP config data
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[1]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["TCP/UDP pointer"] = *edgeME
	}
	if v, err := attrs["Message ID pointer"]; err && v.(float64) != 65535 {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges) //edge point to a large string whose value specifies the message id
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[2]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["Message ID pointer"] = *edgeME
	}

	chart.Note = nil
	return *chart
}

func (me *meMgcCD) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meMgcCD)
	me.relShip = api.MeRelationship{ShortName: "MgcCD", AssociationType: []uint64{148, 137, 157}, RelationshipType: 1}
	api.MeRelationshipRegist(MgcCD, me)
	api.MeHandlerRegist(MgcCD, me)
}
