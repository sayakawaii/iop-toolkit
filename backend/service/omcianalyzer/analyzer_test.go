/*
# ------------------------------------------------------------
# -- analyzer_test.go
# --
# -- Huang Minghe
# -- 2022-7-30
# ------------------------------------------------------------
*/

package omcianalyzer

import (
	"fmt"
	"testing"
)

func TestAnalyzer(t *testing.T) {
	var logPath string = "../static/uploads/20220729/313635393039383238383533363932343130308eaa32d97567aa25b28eebdc3a283f6f/omciIsam.log"
	// var logPath string = "./testdata/omciLightSpan.log"
	var outputPath string = "../static/uploads/20220729/313635393039383238383533363932343130308eaa32d97567aa25b28eebdc3a283f6f/"
	var c = make(chan LogContent)
	var p = make(chan RespStatusContent)
	go Analyzer(logPath, outputPath, c, p)
	for content := range c {
		fmt.Println(content)
	}
	t.Log("Test OMCI Shape successed")
}

func BenchmarkAnalyzer(b *testing.B) {
	var logPath string = "../static/uploads/20220729/313635393039383238383533363932343130308eaa32d97567aa25b28eebdc3a283f6f/omciLightSpan.log"
	var outputPath string = "../static/uploads/20220729/313635393039383238383533363932343130308eaa32d97567aa25b28eebdc3a283f6f"
	var c chan LogContent
	var p chan RespStatusContent
	Analyzer(logPath, outputPath, c, p)
	b.Log("Benchmark OMCI Shape successed")
}
