package omciShape

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMt2HtmlShapeAllHostsMerged(t *testing.T) {
	logPath := "./testdata/mt2Html.html"
	outputPath := t.TempDir() + "/"

	shaper := ShaperDef[Mt2Html]
	if shaper == nil {
		t.Fatalf("Mt2Html shaper not registered")
	}

	onus := shaper.Shape(logPath, outputPath)
	if len(onus) == 0 {
		t.Fatalf("no ONU data generated")
	}

	onuID := "HWTC-b6aea2ab"
	generatedPath, ok := onus[onuID]
	if !ok {
		t.Fatalf("expected ONU %s not found, generated ONUs: %v", onuID, onus)
	}

	generatedRaw, err := os.ReadFile(generatedPath)
	if err != nil {
		t.Fatalf("read generated file failed: %v", err)
	}

	// Keep output JSON format aligned with current omci shaper output style.
	if !bytes.HasPrefix(generatedRaw, []byte("[\n  {")) {
		t.Fatalf("generated JSON format mismatch, expected indented array output")
	}
	if !bytes.Contains(generatedRaw, []byte("\"Timestamp\"")) || !bytes.Contains(generatedRaw, []byte("\"RawData\"")) {
		t.Fatalf("generated JSON missing required fields")
	}

	generatedData := mustReadOmciMetaData(t, generatedPath)
	if len(generatedData) == 0 {
		t.Fatalf("generated file has no OMCI records")
	}
	assertRawDataLen80(t, generatedPath, generatedData)

	entries, err := os.ReadDir("./testoutput")
	if err != nil {
		t.Fatalf("read testoutput directory failed: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		existingPath := filepath.Join("./testoutput", entry.Name())
		existingData, err := tryReadOmciMetaData(existingPath)
		if err != nil {
			t.Logf("skip malformed sample %s: %v", existingPath, err)
			continue
		}
		if len(existingData) == 0 {
			continue
		}
		assertRawDataLen80(t, existingPath, existingData)
	}
}

func mustReadOmciMetaData(t *testing.T, filePath string) []OmciMetaData {
	t.Helper()

	items, err := tryReadOmciMetaData(filePath)
	if err != nil {
		t.Fatalf("read OMCI metadata failed %s: %v", filePath, err)
	}

	return items
}

func tryReadOmciMetaData(filePath string) ([]OmciMetaData, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(string(data))
	if trimmed == "" {
		return nil, nil
	}

	var items []OmciMetaData
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func assertRawDataLen80(t *testing.T, filePath string, items []OmciMetaData) {
	t.Helper()

	for idx, item := range items {
		if len(item.RawData) != 80 {
			t.Fatalf("RawData length mismatch in %s at index %d: got %d, want 80", filePath, idx, len(item.RawData))
		}
	}
}
