/*
# ------------------------------------------------------------
# -- meVASP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	VASP = 146
)

type meVASP struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meVASP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meVASP)
	me.relShip = api.MeRelationship{ShortName: "VASP", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(VASP, me)
	api.MeHandlerRegist(VASP, me)
}
