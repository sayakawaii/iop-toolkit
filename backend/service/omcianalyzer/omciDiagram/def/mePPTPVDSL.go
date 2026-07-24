/*
# ------------------------------------------------------------
# -- mePPTPVDSL.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	PPTPVDSL = 117
)

type mePPTPVDSL struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *mePPTPVDSL) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(mePPTPVDSL)
	me.relShip = api.MeRelationship{ShortName: "PPTPVDSL", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(PPTPVDSL, me)
}
