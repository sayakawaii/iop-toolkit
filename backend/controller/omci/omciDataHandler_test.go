package omci

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"omciAnalyzer/service/omcianalyzer/omciSchema"
)

func TestFormatDevId(t *testing.T) {
	tests := []struct {
		devId byte
		want  string
	}{
		{omciSchema.DevIdBaseline, "Baseline OMCI"},
		{omciSchema.DevIdExtended, "Extended OMCI"},
		{0x05, "Unknown (0x05)"},
	}
	for _, tc := range tests {
		if got := omciSchema.FormatDevId(tc.devId); got != tc.want {
			t.Fatalf("FormatDevId(0x%02X) = %q, want %q", tc.devId, got, tc.want)
		}
	}
}

func TestAssembleOmciDataIncludesMsgFormat(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "omci.json")
	sample := `[
  {
    "Timestamp": "2026-08-21T08:00:00Z",
    "Omci": "{\"header\":{\"Tcid\":28990,\"AR\":1,\"AK\":0,\"MsgType\":9,\"MsgTypeName\":\"Get\",\"DevId\":10,\"MeClass\":2,\"MeClassName\":\"OnuData\",\"MeInst\":0},\"contents\":{\"AttrMask\":32768,\"Attrs\":[{\"Name\":\"MIB data sync\",\"Value\":0}]}}"
  }
]`
	if err := os.WriteFile(path, []byte(sample), 0o644); err != nil {
		t.Fatal(err)
	}

	rows := AssembleOmciData(path)
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0].MsgFormat != "Baseline OMCI" {
		t.Fatalf("row msgFormat = %q, want Baseline OMCI", rows[0].MsgFormat)
	}
	if rows[0].Content.Header.MsgFormat != "Baseline OMCI" {
		t.Fatalf("header MsgFormat = %q, want Baseline OMCI", rows[0].Content.Header.MsgFormat)
	}
	payloadJSON, err := json.Marshal(rows[0].Content.Payload)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(payloadJSON), "Baseline OMCI") {
		t.Fatalf("payload JSON missing MsgFormat: %s", payloadJSON)
	}
}
