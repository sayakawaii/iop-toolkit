/*
# ------------------------------------------------------------
# -- meEVTOCD.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import (
	"fmt"
	"omciAnalyzer/service/omcianalyzer/omciDiagram/api"
	"strconv"
)

const (
	EVTOCD = 171
)

type meEVTOCD struct {
	relShip    api.MeRelationship
	dscpEnable bool
}

func meDSCPmappingParse(me *meEVTOCD, dscp []interface{}, note map[string]string) {
	hexStr := ""
	for _, decimal := range dscp {
		hex := fmt.Sprintf("%02X", (uint64)(decimal.(float64)))
		hexStr += hex
	}

	binaryStr := ""
	for _, hexChar := range hexStr {
		hex, _ := strconv.ParseInt(string(hexChar), 16, 64)
		binaryStr += fmt.Sprintf("%04b", hex)
	}

	decimalArr := make([]int, 0)
	for i := 0; i < len(binaryStr); i += 3 {
		binaryChunk := binaryStr[i : i+3]
		decimal, _ := strconv.ParseInt(binaryChunk, 2, 64)
		decimalArr = append(decimalArr, int(decimal))
	}

	dscpStr := ""
	for i, pbit := range decimalArr {
		if i > 0 && i%8 == 0 {
			dscpStr += "\n"
		}
		s := fmt.Sprintf("(%d:%d),", i, pbit)
		dscpStr += s
	}
	dscpStr = "[DSCP to p-bit mapping:\n" + dscpStr + "]"
	note["DSCP"] = dscpStr
	me.dscpEnable = false
}

func meEVTOCDruleParse(me *meEVTOCD, ruleTable map[string]interface{}, note map[string]string) {

	filter_outer_priority := (uint64)(ruleTable["FilterOuterPriority"].(float64))
	filter_outer_VID := (uint64)(ruleTable["FilterOuterVID"].(float64))
	filter_outer_TPID := (uint64)(ruleTable["FilterOuterTPID"].(float64))
	// padding1 := ((input[2] & 0x0f) << 8 ) | input[3];

	filter_inner_priority := (uint64)(ruleTable["FilterInnerPriority"].(float64))
	filter_inner_VID := (uint64)(ruleTable["FilterInnerVID"].(float64))
	filter_inner_TPID := (uint64)(ruleTable["FilterInnerTPID"].(float64))
	// padding2 := ((input[6] & 0x0f) << 4) | (input[7] >> 4);
	filter_Ethertype := (uint64)(ruleTable["FilterEthertype"].(float64))

	treatment_tags_to_remove := (uint64)(ruleTable["TreatmentTagsToRemove"].(float64))
	// padding3 := ((input[8] & 0x3f) << 4) | (input[9] >> 4);
	treatment_outer_priority := (uint64)(ruleTable["TreatmentOuterPriority"].(float64))
	treatment_outer_VID := (uint64)(ruleTable["TreatmentOuterVID"].(float64))
	treatment_outer_TPID := (uint64)(ruleTable["TreatmentOuterTPID"].(float64))

	// padding4 := (input[12]<<4) | (input[13] >> 4);
	treatment_inner_priority := (uint64)(ruleTable["TreatmentInnerPriority"].(float64))
	treatment_inner_VID := (uint64)(ruleTable["TreatmentInnerVID"].(float64))
	treatment_inner_TPID := (uint64)(ruleTable["TreatmentInnerTPID"].(float64))

	ruleFilter := "(" + strconv.FormatUint(filter_outer_priority, 10) + ", " + strconv.FormatUint(filter_outer_VID, 10) + ", " + strconv.FormatUint(filter_outer_TPID, 10) + "), (" + strconv.FormatUint(filter_inner_priority, 10) + ", " + strconv.FormatUint(filter_inner_VID, 10) + ", " + strconv.FormatUint(filter_inner_TPID, 10) + "), " + strconv.FormatUint(filter_Ethertype, 10)
	ruleTreatment := strconv.FormatUint(treatment_tags_to_remove, 10) + ", (" + strconv.FormatUint(treatment_outer_priority, 10) + ", " + strconv.FormatUint(treatment_outer_VID, 10) + ", " + strconv.FormatUint(treatment_outer_TPID, 10) + "), (" + strconv.FormatUint(treatment_inner_priority, 10) + ", " + strconv.FormatUint(treatment_inner_VID, 10) + ", " + strconv.FormatUint(treatment_inner_TPID, 10) + ")"

	if treatment_outer_VID == 8191 && treatment_inner_VID == 8191 {
		//delete
		delete(note, ruleFilter)
	} else {
		note[ruleFilter] = ruleFilter + " || " + ruleTreatment
		/*
			dscp enable while treatment inner priority set as 10
			refrence G.988 9.3.13 Extended VLAN tagging operation configuration data
			Received frame VLAN tagging operation table -> Treatment inner priority
		*/
		if treatment_inner_priority == 10 {
			me.dscpEnable = true
		}
	}
}

func (me *meEVTOCD) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	if v, err := attrs["Received frame VLAN tagging operation data"]; err {
		var rule = make(map[string]interface{})
		rule = v.([]interface{})[0].(map[string]interface{})
		meEVTOCDruleParse(me, rule, chart.Note)
	}
	if v, err := attrs["DSCP to p-bit mapping"]; err && me.dscpEnable {
		var dscp = make([]interface{}, 0)
		dscp = v.([]interface{})
		meDSCPmappingParse(me, dscp, chart.Note)
	}
	var tpType, tpPointer uint64
	if v, err := attrs["Association type"]; err {
		tpType = (uint64)(v.(float64))
	}
	if v, err := attrs["Associated me pointer"]; err {
		tpPointer = (uint64)(v.(float64))
	}
	if tpPointer == 0 {
		return *chart
	}
	edgeTP := new(api.MeEdges)
	relship := me.relShip
	tpMeName := api.MeRelationshipDef[relship.AssociationType[tpType]].GetRelShip().ShortName
	eType := relship.RelationshipType
	eNote := []string(nil)
	eDir := "left"
	edgeTP.MeEdgesSet(eType, eDir, tpMeName, tpPointer, eNote)
	chart.Edge["Associated me pointer"] = *edgeTP
	return *chart
}

func (me *meEVTOCD) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meEVTOCD)
	me.relShip = api.MeRelationship{ShortName: "EVTOCD", AssociationType: []uint64{47, 138, 11, 134, 98, 266, 281, 162, 0, 286, 329, 333}, RelationshipType: 1}

	me.dscpEnable = false
	api.MeRelationshipRegist(EVTOCD, me)
	api.MeHandlerRegist(EVTOCD, me)
}
