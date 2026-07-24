/*
# ------------------------------------------------------------
# -- meSUD.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	SUD = 153
)

type meSUD struct {
	relShip api.MeRelationship
}

func (me *meSUD) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)

	relship := me.relShip
	eType := relship.RelationshipType
	var eDir string
	eNote := []string(nil)

	var valuePointer uint64
	if v, err := attrs["SIP agent pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges) //edge point to the SIP agent config data
		eDir = "up"
		//pointerMeName := api.MeRelationshipDef[150].ShortName
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["SIP agent pointer"] = *edgeME
	}
	if v, err := attrs["User part AOR"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges) //edge point to a large string that contains the user identification part
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[1]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["User part AOR"] = *edgeME
	}
	if v, err := attrs["Username and password"]; err && v.(float64) != 65535 {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges) //edge point to an authentication security method
		eDir = "right"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[2]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["Username and password"] = *edgeME
	}
	if v, err := attrs["Voicemail server SIP URI"]; err && v.(float64) != 65535 {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges) //edge point to a network address
		eDir = "right"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[3]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["Voicemail server SIP URI"] = *edgeME
	}
	if v, err := attrs["Network dial plan pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges) //edge point to a network dial plan table
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[4]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["Network dial plan pointer"] = *edgeME
	}
	if v, err := attrs["Application services profile pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges) //edge point to a VoIP application services profile
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[5]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["Application services profile pointer"] = *edgeME
	}
	if v, err := attrs["Feature code pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges) //edge point to the VoIP feature access codes
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[6]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["Feature code pointer"] = *edgeME
	}
	if v, err := attrs["PPTP pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges) //edge point to the PPTP POTS UNI
		eDir = "down"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[7]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["PPTP pointer"] = *edgeME
	}

	chart.Note = nil
	return *chart
}

func (me *meSUD) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meSUD)
	me.relShip = api.MeRelationship{ShortName: "SUD", AssociationType: []uint64{150, 157, 148, 137, 145, 146, 147, 53}, RelationshipType: 1}
	api.MeRelationshipRegist(SUD, me)
	api.MeHandlerRegist(SUD, me)
}
