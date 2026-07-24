/*
# ------------------------------------------------------------
# -- mePPTPxDSL1.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	PPTPxDSL1 = 98
)

type mePPTPxDSL1 struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *mePPTPxDSL1) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(mePPTPxDSL1)
	me.relShip = api.MeRelationship{ShortName: "PPTPxDSL1", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(PPTPxDSL1, me)
}
