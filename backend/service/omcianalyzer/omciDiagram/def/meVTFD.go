/*
# ------------------------------------------------------------
# -- meVTFD.go
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
	VTFD = 84
)

type meVTFD struct {
	relShip api.MeRelationship
}

func (me *meVTFD) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var vfList []uint64
	var FWDop, numOfEntr uint64
	if v, err := attrs["Forward operation"]; err {
		FWDop = (uint64)(v.(float64))
	} else {
		return *chart
	}
	if v, err := attrs["Number of entries"]; err {
		numOfEntr = (uint64)(v.(float64))
	} else {
		return *chart
	}
	if v, err := attrs["Vlan filter list"]; err {
		for k := range v.([]interface{}) {
			vfList = append(vfList, (uint64)(v.([]interface{})[k].(float64)))
		}
	} else {
		return *chart
	}
	edgeTP := new(api.MeEdges)
	relship := me.relShip
	tpMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
	eType := relship.RelationshipType
	eNote := []string(nil)
	eDir := "auto"
	tpPointer := node.MeInst
	edgeTP.MeEdgesSet(eType, eDir, tpMeName, tpPointer, eNote)
	chart.Edge["TP pointer"] = *edgeTP
	chart.Note["FWDop"] = "FWDop: " + strconv.FormatUint(FWDop, 10)
	chart.Note["numOfEntr"] = "numOfEntr: " + strconv.FormatUint(numOfEntr, 10)

	// var i uint64
	// for i = 0; i < numOfEntr; i++ {
	// 	vlanID := vfList[i]
	// 	chart.Note["vlanID"+strconv.FormatUint(i, 10)] = "vlanID: " + strconv.FormatUint(vlanID, 10)
	// }
	for i := uint64(0); i < numOfEntr; i++ {
		vlanID := uint16(vfList[i]) // 只取低16位
		priority := vlanID >> 13    // 提取前3位作为优先级
		vlan := vlanID & 0x0FFF     // 提取后12位作为VLAN ID
		chart.Note["vlanID"+strconv.FormatUint(i, 10)] = fmt.Sprintf("vlanID: %d, Priority: %d", vlan, priority)
	}
	return *chart
}

func (me *meVTFD) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meVTFD)
	me.relShip = api.MeRelationship{ShortName: "VTFD", AssociationType: []uint64{47}, RelationshipType: 0}
	api.MeRelationshipRegist(VTFD, me)
	api.MeHandlerRegist(VTFD, me)
}
