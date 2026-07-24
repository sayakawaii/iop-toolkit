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
	"bytes"
	"omciAnalyzer/utils"
	"os/exec"
)

func Draw(logPath string, outputPath string) {
	shape := `./resource/plantUml/plantuml.jar`
	cmd := exec.Command("java", "-jar", shape, logPath, "-tsvg")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		utils.Log(stderr.String())
		return
	}
}
