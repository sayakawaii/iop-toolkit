/*
# ------------------------------------------------------------
# -- meVTOCD.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import (
	"omciAnalyzer/service/omcianalyzer/omciDiagram/api"
	"strconv"
)

const (
	VTOCD = 78
)

type meVTOCD struct {
	relShip api.MeRelationship
}

// parseUStci extracts PCP, DEI and VID from the "Upstream VLAN tag TCI value"
// attribute, whose format is an Array Structure with a single element.
func parseUStci(v interface{}) (pcp, dei, vid uint64, ok bool) {
	arr, ok := v.([]interface{})
	if !ok || len(arr) == 0 {
		return 0, 0, 0, false
	}
	tci, ok := arr[0].(map[string]interface{})
	if !ok {
		return 0, 0, 0, false
	}
	pcpF, ok := tci["Priority Code Point(PCP)"].(float64)
	if !ok {
		return 0, 0, 0, false
	}
	deiF, ok := tci["Drop Eligible Indicator(DEI)"].(float64)
	if !ok {
		return 0, 0, 0, false
	}
	vidF, ok := tci["VLAN Identifier(VID)"].(float64)
	if !ok {
		return 0, 0, 0, false
	}
	return uint64(pcpF), uint64(deiF), uint64(vidF), true
}

func (me *meVTOCD) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var USmode, tpType, tpPointer uint64
	var UStciPCP, UStciDEI, UStciVID uint64
	if v, err := attrs["Association type"]; err {
		tpType = (uint64)(v.(float64))
	} else {
		return *chart
	}
	if v, err := attrs["Associated ME pointer"]; err {
		tpPointer = (uint64)(v.(float64))
	} else {
		return *chart
	}
	if v, err := attrs["Upstream VLAN tagging operation mode"]; err {
		USmode = (uint64)(v.(float64))
	} else {
		return *chart
	}
	if v, err := attrs["Upstream VLAN tag TCI value"]; err {
		pcp, dei, vid, ok := parseUStci(v)
		if !ok {
			return *chart
		}
		UStciPCP, UStciDEI, UStciVID = pcp, dei, vid
	} else {
		return *chart
	}
	edgeTP := new(api.MeEdges)
	relship := me.relShip
	tpMeName := api.MeRelationshipDef[relship.AssociationType[tpType]].GetRelShip().ShortName
	eType := relship.RelationshipType
	eNote := []string(nil)
	eDir := "left"
	edgeTP.MeEdgesSet(eType, eDir, tpMeName, tpPointer, eNote)
	chart.Edge["Associated ME pointer"] = *edgeTP
	chart.Note["USmode"] = "USmode: " + strconv.FormatUint(USmode, 10)
	chart.Note["UStci"] = "UStci: PCP=" + strconv.FormatUint(UStciPCP, 10) +
		", DEI=" + strconv.FormatUint(UStciDEI, 10) +
		", VID=" + strconv.FormatUint(UStciVID, 10)
	return *chart
}

func (me *meVTOCD) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meVTOCD)
	me.relShip = api.MeRelationship{ShortName: "VTOCD", AssociationType: []uint64{11, 134, 138, 47, 98, 266, 281, 162, 0, 286, 11, 329, 333}, RelationshipType: 1}
	api.MeRelationshipRegist(VTOCD, me)
	api.MeHandlerRegist(VTOCD, me)
}
