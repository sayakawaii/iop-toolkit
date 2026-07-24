/*
# ------------------------------------------------------------
# -- meGAL.go
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
	GAL = 272
)

type meGAL struct {
	relShip api.MeRelationship
}

func (me *meGAL) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var mtu uint64
	if v, err := attrs["Maximum GEM payload size"]; err {
		mtu = (uint64)(v.(float64))
	} else {
		return *chart
	}
	chart.Edge = nil
	chart.Note["Maximum GEM payload size"] = "Maximum GEM payload size:" + strconv.FormatUint(mtu, 10)
	return *chart
}

func (me *meGAL) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meGAL)
	me.relShip = api.MeRelationship{ShortName: "GAL", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(GAL, me)
	api.MeHandlerRegist(GAL, me)
}
