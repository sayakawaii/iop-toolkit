/*
# ------------------------------------------------------------
# -- meIPHost.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	IPHost = 134
)

type meIPHost struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meIPHost) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var ipOptions uint64
	if v, err := attrs["IP options"]; err {
		ipOptions = (uint64)(v.(float64))
	} else {
		return *chart
	}
	chart.SetAttr("IPOptions", ipOptions)
	return *chart
}

func (me *meIPHost) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meIPHost)
	me.relShip = api.MeRelationship{ShortName: "IPHost", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(IPHost, me)
	api.MeHandlerRegist(IPHost, me)
}
