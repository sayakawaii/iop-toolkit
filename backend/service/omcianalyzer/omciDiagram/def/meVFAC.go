/*
# ------------------------------------------------------------
# -- meVFAC.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	VFAC = 147
)

type meVFAC struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meVFAC) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meVFAC)
	me.relShip = api.MeRelationship{ShortName: "VFAC", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(VFAC, me)
	api.MeHandlerRegist(VFAC, me)
}
