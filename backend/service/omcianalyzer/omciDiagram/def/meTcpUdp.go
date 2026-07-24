/*
# ------------------------------------------------------------
# -- meTcpUdp.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	TcpUdp = 136
)

type meTcpUdp struct {
	relShip api.MeRelationship
}

func (me *meTcpUdp) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)

	relship := me.relShip
	eType := relship.RelationshipType
	var eDir string
	eNote := []string(nil)

	var valuePointer uint64
	if v, err := attrs["IP host pointer"]; err {
		valuePointer = (uint64)(v.(float64))
		edgeME := new(api.MeEdges)
		eDir = "up"
		pointerMeName := api.MeRelationshipDef[relship.AssociationType[0]].GetRelShip().ShortName
		edgeME.MeEdgesSet(eType, eDir, pointerMeName, valuePointer, eNote)
		chart.Edge["IP host pointer"] = *edgeME
	}

	chart.Note = nil
	return *chart
}

func (me *meTcpUdp) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meTcpUdp)
	me.relShip = api.MeRelationship{ShortName: "TcpUdp", AssociationType: []uint64{134}, RelationshipType: 1}
	api.MeRelationshipRegist(TcpUdp, me)
	api.MeHandlerRegist(TcpUdp, me)
}
