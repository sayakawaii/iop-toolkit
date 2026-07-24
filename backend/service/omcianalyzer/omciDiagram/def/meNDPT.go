/*
# ------------------------------------------------------------
# -- meNDPT.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	NDPT = 145
)

type meNDPT struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meNDPT) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meNDPT)
	me.relShip = api.MeRelationship{ShortName: "NDPT", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(NDPT, me)
	api.MeHandlerRegist(NDPT, me)
}
