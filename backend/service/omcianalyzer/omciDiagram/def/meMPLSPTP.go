/*
# ------------------------------------------------------------
# -- meMPLSPTP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	MPLSPTP = 333
)

type meMPLSPTP struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meMPLSPTP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meMPLSPTP)
	me.relShip = api.MeRelationship{ShortName: "MPLSPTP", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(MPLSPTP, me)
}
