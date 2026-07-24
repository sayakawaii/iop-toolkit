/*
# ------------------------------------------------------------
# -- meASM.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import (
	"omciAnalyzer/service/omcianalyzer/omciDiagram/api"
)

const (
	ASM = 148
)

type meASM struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meASM) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meASM)
	me.relShip = api.MeRelationship{ShortName: "ASM", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(ASM, me)
	api.MeHandlerRegist(ASM, me)
}
