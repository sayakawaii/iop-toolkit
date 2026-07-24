/*
# ------------------------------------------------------------
# -- meTCONT.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import (
	"omciAnalyzer/service/omcianalyzer/omciDiagram/api"
)

const (
	TCONT = 262
)

type meTCONT struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meTCONT) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var AllocID uint64
	if v, err := attrs["Alloc ID"]; err {
		AllocID = (uint64)(v.(float64))
		if AllocID == 0 {
			return *chart
		}
	} else {
		return *chart
	}
	chart.Note = nil
	chart.SetAttr("Port ID", AllocID)
	return *chart
}

func (me *meTCONT) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meTCONT)
	me.relShip = api.MeRelationship{ShortName: "TCONT", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(TCONT, me)
	api.MeHandlerRegist(TCONT, me)
}
