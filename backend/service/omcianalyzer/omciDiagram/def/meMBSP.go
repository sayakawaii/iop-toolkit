/*
# ------------------------------------------------------------
# -- meMBSP.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package def

import "omciAnalyzer/service/omcianalyzer/omciDiagram/api"

const (
	MBSP = 45
)

type meMBSP struct {
	relShip api.MeRelationship
	api.MeChartHandlerBase
}

func (me *meMBSP) GetMeChart(old *api.MeChart, node api.MeNode, attrs map[string]interface{}) api.MeChart {
	chart := new(api.MeChart).Init(node, old)
	var macLearningInd bool
	if v, err := attrs["Learning ind"]; err {
		macLearningInd = (bool)(v.(float64) != 0)
	} else {
		return *chart
	}
	chart.SetAttr("LearningInd", macLearningInd)
	return *chart
}

func (me *meMBSP) GetRelShip() api.MeRelationship {
	return me.relShip
}

func init() {
	me := new(meMBSP)
	me.relShip = api.MeRelationship{ShortName: "MBSP", AssociationType: []uint64{}, RelationshipType: 1}
	api.MeRelationshipRegist(MBSP, me)
	api.MeHandlerRegist(MBSP, me)
}
