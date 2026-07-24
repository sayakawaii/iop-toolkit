/*
# ------------------------------------------------------------
# -- meVEIP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	VEIP = 329
)

type meVEIP struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meVEIP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meVEIP)
	me.relShip = api.MeRelationship{ShortName: "VEIP", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(VEIP, me)
	api.MeHandlerRegist(VEIP, me)
}
