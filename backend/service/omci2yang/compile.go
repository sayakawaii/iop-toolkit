/*
# ------------------------------------------------------------
# -- compile.go
# --
# -- Compile an analysed OMCI log into Nokia LightSpan NETCONF
# -- configuration, by driving the omci2yang compiler.
# ------------------------------------------------------------
*/

package omci2yang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"omciAnalyzer/utils"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const (
	// Laid out like resource/plantUml: the tool ships with the image rather
	// than being installed system-wide.
	toolRoot   = "./resource/omci2yang"
	compilerPy = toolRoot + "/src/compile/compile_xml.py"
	yangRoot   = toolRoot + "/yang"

	// A warm compile takes well under a second. The ceiling is for the cold
	// case where the schema cache has to be rebuilt from ~1000 YANG modules,
	// which takes about 17s on first use after a fresh image.
	compileTimeout = 90 * time.Second
)

// Request is everything the compiler needs that the OMCI log cannot carry.
//
// An OMCI trace shows what landed on the ONU; it says nothing about how the
// OLT refers to it, which is why OnuName is required rather than defaulted.
type Request struct {
	// JSONPath is the per-ONU analysed OMCI file the analyzer already wrote.
	JSONPath string
	// Board selects the YANG set. The mounted tree being compiled against is
	// served by the LT's confd, and the platform-specific modules differ
	// between boards, so validating against the wrong one proves nothing.
	Board string
	// OnuName is the OLT's name for this ONU.
	OnuName string
	// Vendor is the ONU's OUI, which selects vendor YANGMAP overrides.
	Vendor string
	// ChassisName reuses a chassis the OLT already holds; a second one is
	// rejected outright.
	ChassisName string
	// VeipComponent is the virtual-UNI component a VEIP rides. Only an HGU
	// reports one, and without it the ONU instantiates no VEIP.
	VeipComponent string
}

// Result mirrors the compiler's own JSON, including the gaps.
//
// A gap is not a failure: it says the document is schema-correct but
// incomplete, e.g. a statically addressed IP host whose address the trace
// never carried. Whoever receives the configuration is the one who has to see
// them, so they travel with the XML instead of being logged away from it.
type Result struct {
	OnuName            string   `json:"onuName"`
	Xml                string   `json:"xml"`
	Services           int      `json:"services"`
	Valid              bool     `json:"valid"`
	Gaps               []string `json:"gaps"`
	Notes              []string `json:"notes"`
	ValidationErrors   []string `json:"validationErrors"`
	ValidationWarnings []string `json:"validationWarnings"`
	Summary            string   `json:"summary"`
	Board              string   `json:"board"`
	// Error is set when the compiler could not run to completion at all,
	// which is a different outcome from a config that failed validation.
	Error string `json:"error,omitempty"`
}

// clip keeps a diagnostic short enough to belong in an HTTP response.
func clip(text string, limit int) string {
	if len(text) <= limit {
		return text
	}
	return text[:limit] + "..."
}

// Boards lists the YANG sets present in the image, newest layout first.
//
// Read off disk rather than hard-coded so the backend and the UI cannot drift
// apart from what was actually shipped.
func Boards() []string {
	entries, err := os.ReadDir(yangRoot)
	if err != nil {
		utils.Log("omci2yang: cannot read", yangRoot, err)
		return nil
	}
	var boards []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		boards = append(boards, strings.ToUpper(entry.Name()))
	}
	sort.Strings(boards)
	return boards
}

// yangDirFor resolves a board name to its YANG directory, rejecting anything
// that is not a directory actually present. Board names reach here from a
// request, so this doubles as the guard against path traversal.
func yangDirFor(board string) (string, error) {
	if board == "" {
		return "", fmt.Errorf("no board selected")
	}
	wanted := strings.ToLower(board)
	for _, candidate := range Boards() {
		if strings.ToLower(candidate) == wanted {
			return filepath.Join(yangRoot, wanted), nil
		}
	}
	return "", fmt.Errorf("unknown board %q; have %s",
		board, strings.Join(Boards(), ", "))
}

// Compile runs the compiler over one analysed ONU and returns its verdict.
//
// The compiler is a separate process for a plain reason: it is Python and this
// is Go. It is a pure function of its arguments -- no state, no network, sub
// second -- so a synchronous call is the honest shape for it, the same way
// plantUml.Draw shells out to render a diagram.
func Compile(request Request) Result {
	yangDir, err := yangDirFor(request.Board)
	if err != nil {
		return Result{Error: err.Error()}
	}
	if _, err := os.Stat(request.JSONPath); err != nil {
		return Result{Error: fmt.Sprintf(
			"analysed OMCI for this ONU is not on disk: %v", err)}
	}

	args := []string{
		compilerPy, request.JSONPath,
		"--from-omci",
		"--onu-name", request.OnuName,
		"--yang-dir", yangDir,
		"--json",
	}
	for _, optional := range []struct{ flag, value string }{
		{"--vendor", request.Vendor},
		{"--chassis-name", request.ChassisName},
		{"--veip-component", request.VeipComponent},
	} {
		if optional.value != "" {
			args = append(args, optional.flag, optional.value)
		}
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("python3", args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	done := make(chan error, 1)
	if err := cmd.Start(); err != nil {
		return Result{Error: fmt.Sprintf("cannot start the compiler: %v", err)}
	}
	go func() { done <- cmd.Wait() }()

	select {
	case <-time.After(compileTimeout):
		_ = cmd.Process.Kill()
		<-done
		return Result{Error: fmt.Sprintf("the compiler did not finish within %s",
			compileTimeout)}
	case runErr := <-done:
		// Exit status 1 means the config failed validation, which is a real
		// result and still carries XML and errors; only a broken run has no
		// JSON to report.
		var result Result
		if jsonErr := json.Unmarshal(stdout.Bytes(), &result); jsonErr != nil {
			detail := strings.TrimSpace(stderr.String())
			if detail == "" {
				detail = strings.TrimSpace(stdout.String())
			}
			utils.Log("omci2yang: unreadable compiler output:", jsonErr, detail)
			return Result{Error: fmt.Sprintf("the compiler failed: %v: %s",
				runErr, clip(detail, 500))}
		}
		result.Board = strings.ToUpper(request.Board)
		return result
	}
}
