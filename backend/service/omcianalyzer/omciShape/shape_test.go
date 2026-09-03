/*
# ------------------------------------------------------------
# -- shape_test.go
# --
# -- Huang Minghe
# -- 2022-6-21
# ------------------------------------------------------------
*/

package omciShape

import (
	"fmt"
	"testing"
)

func TestOmciShape(t *testing.T) {
	// var logPath string = "./testdata/mt2xml.xml"
	var logPath string = "./testdata/huaweiOnu.log" //nolint:unused
	var outputPath string = "./testoutput/"
	// logType := utils.Invalide
	// var logType utils.LogType
	logType := LogTypeAutoIdentify(logPath)
	var onus map[string]string
	// onus = map[string]string{}
	if logType != Invalid {
		shaper := ShaperDef[logType]
		onus = shaper.Shape(logPath, outputPath)
		for k, v := range onus {
			fmt.Println(k, v)
		}
	} else {
		t.Error("invalid log file")
	}
	t.Log("Test OMCI Shape successed")
}

func BenchmarkOmciShape(b *testing.B) {
	var logPath string = "./testdata/lightSpan.log"
	var outputPath string = "./testoutput/"
	// logType := utils.Invalide
	// var logType utils.LogType
	logType := LogTypeAutoIdentify(logPath)
	var onus map[string]string
	// onus = map[string]string{}
	if logType != Invalid {
		shaper := ShaperDef[logType]
		onus = shaper.Shape(logPath, outputPath)
		for k, v := range onus {
			fmt.Println(k, v)
		}
	} else {
		b.Error("invalid log file")
	}
	b.Log("Benchmark OMCI Shape successed")
}
