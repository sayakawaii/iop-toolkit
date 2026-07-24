/*
# ------------------------------------------------------------
# -- mePPTPRS232UNI.go
# --
# -- Huang Minghe
# -- 2026-3-11
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	PPTPRS232UNI = 401
)

type mePPTPRS232UNI struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *mePPTPRS232UNI) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(mePPTPRS232UNI)
	me.relShip = api.MeRelationship{ShortName: "PPTPRS232UNI", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(PPTPRS232UNI, me)
	api.MeHandlerRegist(PPTPRS232UNI, me)
}
