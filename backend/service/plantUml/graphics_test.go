/*
# ------------------------------------------------------------
# -- planuml_test.go
# --
# -- Huang Minghe
# -- 2022-6-21
# ------------------------------------------------------------
*/

package plantUml

import (
	"testing"
)

func TestOmciShape(t *testing.T) {
	var logPath string = "./testdata/I240GR.wsd"
	var outputPath string = "./testoutput/"
	// logType := trace.Invalide
	Draw(logPath, outputPath)
	// } else {
	// 	t.Error("invalid log file")
	// }
	t.Log("Test plantUML successed")
}
