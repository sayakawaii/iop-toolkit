/*
# ------------------------------------------------------------
# -- meVSP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	VSP = 58
)

type meVSP struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meVSP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meVSP)
	me.relShip = api.MeRelationship{ShortName: "VSP", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(VSP, me)
	api.MeHandlerRegist(VSP, me)
}
