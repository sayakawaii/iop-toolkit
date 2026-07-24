/*
# ------------------------------------------------------------
# -- meMOP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import (
	"fmt"
	"omciAnalyzer/service/omcianalyzer/omciDiagram/api"
	"omciAnalyzer/service/omcianalyzer/omciSchema"
)

const (
	MOP = 309
)

type meMOP struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

type tableControl struct {
	RowKey    uint64 `json:"Row key"`
	RowPartID uint64 `json:"Row part ID"`
	SetCtrl   uint64 `json:"Set ctrl"`
	Test      uint64 `json:"Test"`
}

type rowPart0 struct {
	GemPortID             uint64 `json:"GEM port ID"`
	VlanIDAni             uint64 `json:"VLAN ID ANI"`
	SourceIP              string `json:"Source IP address"`
	DestIPStart           string `json:"Destination IP address start of range"`
	DestIPEnd             string `json:"Destination IP address end of range"`
	ImputedGroupBandwidth uint64 `json:"Imputed group bandwidth"`
	Reserved              uint64 `json:"Reserved"`
}

func (row *rowPart0) set(mapData interface{}) *rowPart0 {
	err := omciSchema.GetContentFromInterface(mapData, row)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return row
}

func (row *rowPart0) get() *rowPart0 {
	return row
}

func (row *rowPart0) getRowPartString() string {
	rowPartString := ""
	rowPartString += "GEM port ID: " + fmt.Sprintf("%d", row.GemPortID) + ", "
	rowPartString += "VLAN ID ANI: " + fmt.Sprintf("%d", row.VlanIDAni) + ", "
	rowPartString += "Source IP address: " + row.SourceIP + ", "
	rowPartString += "Destination IP address start of range: " + row.DestIPStart + ", "
	rowPartString += "Destination IP address end of range: " + row.DestIPEnd + ", "
	rowPartString += "Imputed group bandwidth: " + fmt.Sprintf("%d", row.ImputedGroupBandwidth) + ", "
	rowPartString += "Reserved: " + fmt.Sprintf("%d", row.Reserved) + "\n"
	return rowPartString
}

type rowPart1 struct {
	LeadingBytes       string `json:"Leading bytes of IPv6 source address"`
	PreviewLength      uint64 `json:"Preview length"`
	PreviewRepeatTime  uint64 `json:"Preview repeat time"`
	PreviewRepeatCount uint64 `json:"Preview repeat count"`
	PreviewResetTime   uint64 `json:"Preview reset time"`
	Reserved           uint64 `json:"Reserved"`
}

func (row *rowPart1) set(mapData interface{}) *rowPart1 {
	err := omciSchema.GetContentFromInterface(mapData, row)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return row
}

func (row *rowPart1) get() *rowPart1 {
	return row
}

func (row *rowPart1) getRowPartString() string {
	rowPartString := ""
	rowPartString += "Leading bytes of IPv6 source address: " + row.LeadingBytes + ", "
	rowPartString += "Preview length: " + fmt.Sprintf("%d", row.PreviewLength) + ", "
	rowPartString += "Preview repeat time: " + fmt.Sprintf("%d", row.PreviewRepeatTime) + ", "
	rowPartString += "Preview repeat count: " + fmt.Sprintf("%d", row.PreviewRepeatCount) + ", "
	rowPartString += "Preview reset time: " + fmt.Sprintf("%d", row.PreviewResetTime) + "\n"
	rowPartString += "Reserved: " + fmt.Sprintf("%d", row.Reserved) + "\n"
	return rowPartString
}

type rowPart2 struct {
	LeadingBytes string `json:"Leading bytes of IPv6 destination address"`
	Reserved     uint64 `json:"Reserved"`
}

func (row *rowPart2) set(mapData interface{}) *rowPart2 {
	err := omciSchema.GetContentFromInterface(mapData, row)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return row
}

func (row *rowPart2) get() *rowPart2 {
	return row
}

func (row *rowPart2) getRowPartString() string {
	rowPartString := ""
	rowPartString += "Leading bytes of IPv6 destination address: " + row.LeadingBytes + ", "
	rowPartString += "Reserved: " + fmt.Sprintf("%d", row.Reserved) + "\n"
	return rowPartString
}

type aclTableStruct struct {
	TableControl tableControl `json:"Table control"`
	TableContent interface{}  `json:"Table content"`
}

func (acl *aclTableStruct) set(mapData interface{}) *aclTableStruct {
	err := omciSchema.GetContentFromInterface(mapData, acl)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return acl
}

func (acl *aclTableStruct) get() *aclTableStruct {
	return acl
}

func (acl *aclTableStruct) writeRowToNote(note map[string]string) {
	newList := ""
	newList += "RowKey: " + fmt.Sprintf("%d", acl.TableControl.RowKey) + ", "
	newList += "RowPartID: " + fmt.Sprintf("%d", acl.TableControl.RowPartID) + ", "
	newList += "SetCtrl: " + fmt.Sprintf("%d", acl.TableControl.SetCtrl) + ", "
	newList += "Test: " + fmt.Sprintf("%d", acl.TableControl.Test) + ", "
	if acl.TableControl.RowPartID == 0 {
		c := acl.TableContent.(map[string]interface{})["Row part 0"]
		row := new(rowPart0)
		row.set(c)
		newList += row.getRowPartString()
	} else if acl.TableControl.RowPartID == 1 {
		c := acl.TableContent.(map[string]interface{})["Row part 1"]
		row := new(rowPart1)
		row.set(c)
		newList += row.getRowPartString()
	} else if acl.TableControl.RowPartID == 2 {
		c := acl.TableContent.(map[string]interface{})["Row part 2"]
		row := new(rowPart2)
		row.set(c)
		newList += row.getRowPartString()
	}
	if v, err := note["DACL"]; err {
		list := v + newList
		note["DACL"] = list
	} else {
		note["DACL"] = newList
	}
}

func (me *meMOP) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	// get DACL
	getDACL := func(key string) bool {
		if v, err := attrs[key]; err {
			r := v.([]interface{})[0].(map[string]interface{})
			dacl := new(aclTableStruct)
			dacl.set(r)
			dacl.writeRowToNote(chart.Note)
			return true
		}
		return false
	}
	if getDACL("Dynamic access control list table") || getDACL("Static access control list table") {
		chart.Edge = nil
		return *chart
	} else {
		return *chart
	}
}

func (me *meMOP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meMOP)
	me.relShip = api.MeRelationship{ShortName: "MOP", AssociationType: []uint64{310}, RelationshipType: 1}
	api.MeRelationshipRegist(MOP, me)
	api.MeHandlerRegist(MOP, me)
}
