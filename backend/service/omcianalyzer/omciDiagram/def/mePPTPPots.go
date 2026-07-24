/*
# ------------------------------------------------------------
# -- mePPTPPots.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	PPTPPots = 53
)

type mePPTPPots struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *mePPTPPots) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(mePPTPPots)
	me.relShip = api.MeRelationship{ShortName: "PPTPPots", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(PPTPPots, me)
	api.MeHandlerRegist(PPTPPots, me)
}
