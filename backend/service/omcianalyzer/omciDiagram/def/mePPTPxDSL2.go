/*
# ------------------------------------------------------------
# -- mePPTPxDSL2.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	PPTPxDSL2 = 99
)

type mePPTPxDSL2 struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *mePPTPxDSL2) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(mePPTPxDSL2)
	me.relShip = api.MeRelationship{ShortName: "PPTPxDSL2", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(PPTPxDSL2, me)
}
