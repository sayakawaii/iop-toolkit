/*
# ------------------------------------------------------------
# -- meIVCCTP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	IVCCTP = 14
)

type meIVCCTP struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meIVCCTP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meIVCCTP)
	me.relShip = api.MeRelationship{ShortName: "IVCCTP", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(IVCCTP, me)
}
