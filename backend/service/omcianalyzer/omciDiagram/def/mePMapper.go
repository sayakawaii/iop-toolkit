/*
# ------------------------------------------------------------
# -- mePMapper.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	PMapper = 130
)

type mePMapper struct {
	relShip api.MeRelationship
}

func (me *mePMapper) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var InterworkTPpointer, tpType uint64
	if v, err := attrs["TP type"]; err {
		tpType = (uint64)(v.(float64)) //keep this code for code rule check
		tpType = 0                     //current always set tpType as 0, not support other tpType. TODO:add other tpType supprot
	} else {
		tpType = 0
	}
	// if v, err := attrs["TP pointer"]; err {
	// 	tpPointer = v
	// }else {
	// 	tpPointer = "0xFFFF"
	// }
	relship := me.relShip
	tpMeName := api.MeRelationshipDef[relship.AssociationType[tpType]].GetRelShip().ShortName
	eType := relship.RelationshipType
	eDir := "up"
	if v, err := attrs["Interwork TP pointer for P-bit priority 0"]; err {
		InterworkTPpointer = (uint64)(v.(float64))
		if InterworkTPpointer != 65535 {
			edgeTP := new(api.MeEdges)
			var eNote []string
			eNote = append(eNote, "0")
			edgeTP.MeEdgesSet(eType, eDir, tpMeName, InterworkTPpointer, eNote)
			chart.Edge["Interwork TP pointer for P-bit priority 0"] = *edgeTP
		}
	}
	if v, err := attrs["Interwork TP pointer for P-bit priority 1"]; err {
		InterworkTPpointer = (uint64)(v.(float64))
		if InterworkTPpointer != 65535 {
			edgeTP := new(api.MeEdges)
			var eNote []string
			eNote = append(eNote, "1")
			edgeTP.MeEdgesSet(eType, eDir, tpMeName, InterworkTPpointer, eNote)
			chart.Edge["Interwork TP pointer for P-bit priority 1"] = *edgeTP
		}
	}
	if v, err := attrs["Interwork TP pointer for P-bit priority 2"]; err {
		InterworkTPpointer = (uint64)(v.(float64))
		if InterworkTPpointer != 65535 {
			edgeTP := new(api.MeEdges)
			var eNote []string
			eNote = append(eNote, "2")
			edgeTP.MeEdgesSet(eType, eDir, tpMeName, InterworkTPpointer, eNote)
			chart.Edge["Interwork TP pointer for P-bit priority 2"] = *edgeTP
		}
	}
	if v, err := attrs["Interwork TP pointer for P-bit priority 3"]; err {
		InterworkTPpointer = (uint64)(v.(float64))
		if InterworkTPpointer != 65535 {
			edgeTP := new(api.MeEdges)
			var eNote []string
			eNote = append(eNote, "3")
			edgeTP.MeEdgesSet(eType, eDir, tpMeName, InterworkTPpointer, eNote)
			chart.Edge["Interwork TP pointer for P-bit priority 3"] = *edgeTP
		}
	}
	if v, err := attrs["Interwork TP pointer for P-bit priority 4"]; err {
		InterworkTPpointer = (uint64)(v.(float64))
		if InterworkTPpointer != 65535 {
			edgeTP := new(api.MeEdges)
			var eNote []string
			eNote = append(eNote, "4")
			edgeTP.MeEdgesSet(eType, eDir, tpMeName, InterworkTPpointer, eNote)
			chart.Edge["Interwork TP pointer for P-bit priority 4"] = *edgeTP
		}
	}
	if v, err := attrs["Interwork TP pointer for P-bit priority 5"]; err {
		InterworkTPpointer = (uint64)(v.(float64))
		if InterworkTPpointer != 65535 {
			edgeTP := new(api.MeEdges)
			var eNote []string
			eNote = append(eNote, "5")
			edgeTP.MeEdgesSet(eType, eDir, tpMeName, InterworkTPpointer, eNote)
			chart.Edge["Interwork TP pointer for P-bit priority 5"] = *edgeTP
		}
	}
	if v, err := attrs["Interwork TP pointer for P-bit priority 6"]; err {
		InterworkTPpointer = (uint64)(v.(float64))
		if InterworkTPpointer != 65535 {
			edgeTP := new(api.MeEdges)
			var eNote []string
			eNote = append(eNote, "6")
			edgeTP.MeEdgesSet(eType, eDir, tpMeName, InterworkTPpointer, eNote)
			chart.Edge["Interwork TP pointer for P-bit priority 6"] = *edgeTP
		}
	}
	if v, err := attrs["Interwork TP pointer for P-bit priority 7"]; err {
		InterworkTPpointer = (uint64)(v.(float64))
		if InterworkTPpointer != 65535 {
			edgeTP := new(api.MeEdges)
			var eNote []string
			eNote = append(eNote, "7")
			edgeTP.MeEdgesSet(eType, eDir, tpMeName, InterworkTPpointer, eNote)
			chart.Edge["Interwork TP pointer for P-bit priority 7"] = *edgeTP
		}
	}
	chart.Note = nil
	return *chart
}

func (me *mePMapper) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(mePMapper)
	me.relShip = api.MeRelationship{ShortName: "PMapper", AssociationType: []uint64{266}, RelationshipType: 1}
	api.MeRelationshipRegist(PMapper, me)
	api.MeHandlerRegist(PMapper, me)
}
