/*
# ------------------------------------------------------------
# -- sequencetracer_test.go
# --
# -- Huang Minghe
# -- 2024-8-2
# ------------------------------------------------------------
*/
package sequencetracer

import (
	"fmt"
	"os"
	"testing"
)

// GenerateHTML generates an HTML file with Mermaid diagram
func generateHTML(filename, mermaid string) error {
	html := fmt.Sprintf(`
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Mermaid Diagram</title>
        <script type="module">
            import mermaid from 'https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs';
            mermaid.initialize({ startOnLoad: true });
        </script>
    </head>
    <body>
        <div class="mermaid">
            %s
        </div>
    </body>
    </html>`, mermaid)

	return os.WriteFile(filename, []byte(html), 0644)
}

func TestSequenceTracer(t *testing.T) {
	log1 := "./testdata/omci_slow_OnuMgntOlt.log"
	log2 := "./testdata/omci_slow_xponhwa.log"
	// log3 := "./testdata/trace.log"
	// log4 := "./testdata/find-3666_sn_case4_xponInframgnt"
	// var logPath string = "./testdata/omciLightSpan.log"
	var outputPath string = "./testoutput/"
	logs := []string{log1, log2}
	mermaid := SequenceTracer(logs)
	t.Log("mermaid:", mermaid)

	// Generate HTML file
	if err := generateHTML(outputPath+"diagram.html", mermaid); err != nil {
		t.Log("Error generating HTML file:", err)
		return
	}

	t.Log("HTML file generated successfully.")
	// for content := range c {
	// 	fmt.Println(content)
	// }
	t.Log("Test SequenceTracer successed")
}

func BenchmarkSequenceTracer(b *testing.B) {
	var logPath string = "./testdata/trace.log"
	// var outputPath string = "./testoutput/"
	logs := []string{logPath}
	SequenceTracer(logs)
	b.Log("Benchmark SequenceTracer successed")
}
