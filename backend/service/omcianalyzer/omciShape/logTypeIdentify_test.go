package omciShape

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLogTypeAutoIdentify_IsamRequiresAngleBrackets(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "diag_only.log")
	content := `[00:01:19.283]:OMCI <---> T<0xc72dd080>:OMCI MSG: get a mib reset
[00:01:19.283]:OMCI <---> T<0xc72dd080>:MIB RESET: set_to_default done
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := LogTypeAutoIdentify(path); got != Invalid {
		t.Fatalf("LogTypeAutoIdentify() = %d (%s), want Invalid", got, LogTypeNameGet(got))
	}
}

func TestLogTypeAutoIdentify_NokiaOnuAfterIsamDiagPrefix(t *testing.T) {
	logFile := "./testdata/nokia_onu_isam_diag_prefix.txt"
	if _, err := os.Stat(logFile); err != nil {
		t.Skipf("testdata file not available: %s", logFile)
	}
	if got := LogTypeAutoIdentify(logFile); got != NokiaOnu {
		t.Fatalf("LogTypeAutoIdentify() = %d (%s), want %d (%s)", got, LogTypeNameGet(got), NokiaOnu, LogTypeNameGet(NokiaOnu))
	}
}

func TestLogTypeAutoIdentify_IsamOmciIsam2StillMatches(t *testing.T) {
	logFile := "./testdata/isam_omciIsam2.txt"
	if _, err := os.Stat(logFile); err != nil {
		t.Skipf("testdata file not available: %s", logFile)
	}
	if got := LogTypeAutoIdentify(logFile); got != Isam {
		t.Fatalf("LogTypeAutoIdentify() = %d (%s), want %d (%s)", got, LogTypeNameGet(got), Isam, LogTypeNameGet(Isam))
	}
}
