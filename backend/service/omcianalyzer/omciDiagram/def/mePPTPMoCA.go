/*
# ------------------------------------------------------------
# -- mePPTPMoCA.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	PPTPMoCA = 162
)

type mePPTPMoCA struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *mePPTPMoCA) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(mePPTPMoCA)
	me.relShip = api.MeRelationship{ShortName: "PPTPMoCA", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(PPTPMoCA, me)
}
