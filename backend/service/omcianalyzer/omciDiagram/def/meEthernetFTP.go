/*
# ------------------------------------------------------------
# -- meEthernetFTP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	EthernetFTP = 286
)

type meEthernetFTP struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meEthernetFTP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meEthernetFTP)
	me.relShip = api.MeRelationship{ShortName: "EthernetFTP", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(EthernetFTP, me)
}
