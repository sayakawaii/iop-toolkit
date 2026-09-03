/*
# ------------------------------------------------------------
# -- MeChart_test.go
# --
# -- Huang Minghe
# -- 2022-6-25
# ------------------------------------------------------------
*/
package omciDiagram

import (
	"fmt"
	"testing"
)

func TestMeChart(t *testing.T) {
	fmt.Println("omciDiagram MeChart test")

	var logPath string = "./testdata/CINA.json"
	var outputPath string = "./testoutput/"
	diagramSet, _ := DiagramProc(logPath)

	fmt.Println("chart:")
	for _, v := range diagramSet {
		// fmt.Println(k, v)
		onuName := "V-ANI-123"
		fileName := outputPath + onuName + "-" + v.Category + "-" + v.Name + ".wsd"
		PlantUmlGenerate(fileName, v.Chart)
	}
	t.Log("Test MeChart end")
}
