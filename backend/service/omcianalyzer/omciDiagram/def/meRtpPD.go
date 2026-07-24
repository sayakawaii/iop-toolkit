/*
# ------------------------------------------------------------
# -- meRtpPD.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	RtpPD = 143
)

type meRtpPD struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meRtpPD) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meRtpPD)
	me.relShip = api.MeRelationship{ShortName: "RtpPD", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(RtpPD, me)
	api.MeHandlerRegist(RtpPD, me)
}
