/*
# ------------------------------------------------------------
# -- meIPV6Host.go
# --
# -- Huang Minghe
# -- 2026-3-11
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	IPV6Host = 347
)

type meIPV6Host struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meIPV6Host) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var ipOptions uint64
	if v, err := attrs["IPv6 options"]; err {
		ipOptions = (uint64)(v.(float64))
	} else {
		return *chart
	}
	chart.SetAttr("IPv6Options", ipOptions)
	return *chart
}

func (me *meIPV6Host) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meIPV6Host)
	me.relShip = api.MeRelationship{ShortName: "IPV6Host", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(IPV6Host, me)
	api.MeHandlerRegist(IPV6Host, me)
}
