/*
# ------------------------------------------------------------
# -- meLS.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	LS = 157
)

type meLS struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meLS) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meLS)
	me.relShip = api.MeRelationship{ShortName: "LS", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(LS, me)
	api.MeHandlerRegist(LS, me)
}
