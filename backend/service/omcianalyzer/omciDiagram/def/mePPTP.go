/*
# ------------------------------------------------------------
# -- mePPTP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	PPTP = 11
)

type mePPTP struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *mePPTP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(mePPTP)
	me.relShip = api.MeRelationship{ShortName: "PPTP", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(PPTP, me)
	api.MeHandlerRegist(PPTP, me)
}
