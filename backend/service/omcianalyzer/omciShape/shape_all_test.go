package omciShape

import (
	"bytes"
	"os"
	"testing"
)

// shaperTestCase defines a single test scenario for an OMCI log shaper.
type shaperTestCase struct {
	logType         uint8
	logFile         string
	expectedOnus    []string // expected ONU IDs in the output
	allowVariableLen bool     // true for extended OMCI / PCAP where RawData may not be 80
}

// runShaperTest is the common test body shared by all shaper tests.
// It verifies:
//  1. shaper is registered
//  2. log type is correctly identified
//  3. output contains expected ONU IDs
//  4. output is valid indented JSON with Timestamp + RawData fields
//  5. every RawData entry has length 80
func runShaperTest(t *testing.T, tc shaperTestCase) {
	t.Helper()

	if _, err := os.Stat(tc.logFile); err != nil {
		t.Skipf("testdata file not available: %s", tc.logFile)
	}

	// Verify log type identification.
	identified := LogTypeAutoIdentify(tc.logFile)
	if identified != tc.logType {
		t.Fatalf("LogTypeAutoIdentify(%s) = %d (%s), want %d (%s)",
			tc.logFile, identified, LogTypeNameGet(identified), tc.logType, LogTypeNameGet(tc.logType))
	}

	shaper := ShaperDef[tc.logType]
	if shaper == nil {
		t.Fatalf("shaper not registered for type %d (%s)", tc.logType, LogTypeNameGet(tc.logType))
	}

	outputPath := t.TempDir() + "/"
	onus := shaper.Shape(tc.logFile, outputPath)
	if len(onus) == 0 {
		t.Fatalf("shaper produced no ONU output")
	}

	for _, expectedOnu := range tc.expectedOnus {
		generatedPath, ok := onus[expectedOnu]
		if !ok {
			t.Fatalf("expected ONU %q not found, generated ONUs: %v", expectedOnu, onuKeys(onus))
		}
		validateOutputFile(t, generatedPath, tc.allowVariableLen)
	}
}

// validateOutputFile checks a single output file for correct JSON format and content.
func validateOutputFile(t *testing.T, filePath string, allowVariableLen bool) {
	t.Helper()

	raw, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("read generated file failed: %v", err)
	}

	if !bytes.HasPrefix(raw, []byte("[\n  {")) {
		t.Fatalf("generated JSON format mismatch in %s, expected indented array output", filePath)
	}
	if !bytes.Contains(raw, []byte("\"Timestamp\"")) || !bytes.Contains(raw, []byte("\"RawData\"")) {
		t.Fatalf("generated JSON missing required fields in %s", filePath)
	}

	data := mustReadOmciMetaData(t, filePath)
	if len(data) == 0 {
		t.Fatalf("generated file %s has no OMCI records", filePath)
	}
	if !allowVariableLen {
		assertRawDataLen80(t, filePath, data)
	}
}

func onuKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// ---- Isam (4 log files) ----

func TestIsamShapeOmciIsam(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      Isam,
		logFile:      "./testdata/isam_omciIsam.log",
		expectedOnus: []string{"222", "219"},
	})
}

func TestIsamShapeOmciIsam2(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      Isam,
		logFile:      "./testdata/isam_omciIsam2.txt",
		expectedOnus: []string{"49"},
	})
}

func TestIsamShape5402zT063(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      Isam,
		logFile:      "./testdata/isam_5402z_T063.txt",
		expectedOnus: []string{"1920"},
	})
}

func TestIsamShape5402zUNI1(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      Isam,
		logFile:      "./testdata/isam_5402z_UNI1.txt",
		expectedOnus: []string{"256"},
	})
}

// ---- LightSpan ----

func TestLightSpanShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      LightSpan,
		logFile:      "./testdata/lightSpan.log",
		expectedOnus: []string{"V-ANI-1-1-1"},
	})
}

// ---- NokiaOnu (2 log files) ----

func TestNokiaOnuShapeStandard(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      NokiaOnu,
		logFile:      "./testdata/nokiaOnu_standard.log",
		expectedOnus: []string{"NOKIA-ONU"},
	})
}

func TestNokiaOnuShapeExtended(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:         NokiaOnu,
		logFile:         "./testdata/nokiaOnu_extended.log",
		expectedOnus:    []string{"NOKIA-ONU"},
		allowVariableLen: true,
	})
}

func TestNokiaOnuShapeDiagPrefix(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      NokiaOnu,
		logFile:      "./testdata/nokia_onu_isam_diag_prefix.txt",
		expectedOnus: []string{"NOKIA-ONU"},
	})
}

// ---- CsvRaw ----

func TestCsvRawShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      CsvRaw,
		logFile:      "./testdata/csvRaw_gponDoctor.csv",
		expectedOnus: []string{"AONT"},
	})
}

// ---- Mt2Xml ----

func TestMt2XmlShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      Mt2Xml,
		logFile:      "./testdata/mt2Xml.xml",
		expectedOnus: []string{"ADTN880442d8"},
	})
}

// ---- GtcXml ----

func TestGtcXmlShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      GtcXml,
		logFile:      "./testdata/gtcXml.xml",
		expectedOnus: []string{"DSNW-665726624"},
	})
}

// ---- SstRaw ----

func TestSstRawShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      sstRaw,
		logFile:      "./testdata/sstRaw.json",
		expectedOnus: []string{"AONT"},
	})
}

// ---- CopperOnu ----

func TestCopperOnuShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      CopperOnu,
		logFile:      "./testdata/copperOnu.txt",
		expectedOnus: []string{"CopperOnu"},
	})
}

// ---- ZteRaw ----

func TestZteRawShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      ZteRaw,
		logFile:      "./testdata/zteRaw.raw",
		expectedOnus: []string{"ZTE-ONU"},
	})
}

// ---- HuaweiOnu ----

func TestHuaweiOnuShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      HuaweiOnu,
		logFile:      "./testdata/huaweiOnu.log",
		expectedOnus: []string{"HWTCQ*1e15"},
	})
}

// ---- HisenseOnu ----

func TestHisenseOnuShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      HisenseOnu,
		logFile:      "./testdata/HisenseOnu.log",
		expectedOnus: []string{"HisenseOnu"},
	})
}

// ---- TraceSpanXml ----

func TestTraceSpanXmlShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      TraceSpanXml,
		logFile:      "./testdata/traceSpanXml.xml",
		expectedOnus: []string{"TraceSpanOnu"},
	})
}

// ---- XponHwa ----

func TestXponHwaShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      XponHwa,
		logFile:      "./testdata/xponHwa.txt",
		expectedOnus: []string{"1"},
	})
}

// ---- Pcap ----

func TestPcapShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:         Pcap,
		logFile:         "./testdata/pcap_onu.pcap",
		expectedOnus:    []string{"CINA"},
		allowVariableLen: true,
	})
}

func TestPcapShapeRawOmci(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:         Pcap,
		logFile:         "./testdata/onu_restore-factory_reset_gem.pcap",
		expectedOnus:    []string{"onu_restore-factory_reset_gem"},
		allowVariableLen: true,
	})
}

// ---- Mt2Html (already has dedicated test in mt2HtmlShape_test.go) ----

func TestMt2HtmlShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      Mt2Html,
		logFile:      "./testdata/mt2Html.html",
		expectedOnus: []string{"HWTC-b6aea2ab"},
	})
}

// ---- CalixOnu ----

func TestCalixOnuShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      CalixOnu,
		logFile:      "./testdata/calixOnu.log",
		expectedOnus: []string{"CalixOnu"},
	})
}

// ---- ZteOnu: no testdata file available, skip ----

func TestZteOnuShape(t *testing.T) {
	runShaperTest(t, shaperTestCase{
		logType:      ZteOnu,
		logFile:      "./testdata/zteOnu.log",
		expectedOnus: []string{"ZTE-ONU"},
	})
}
