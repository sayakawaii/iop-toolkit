package omciSchema

import (
	"testing"
)

var (
	logType uint8 = 5
)

func TestOmciRawDataParse(t *testing.T) {
	var logPath string = "./testdata/AONT.log"
	var outputPath string = "./testoutput/AONT.json"
	path := OmciRawDataParse(logPath, outputPath, logType)
	if path == "" {
		t.Error("Test OMCI Shema Failed")
	} else {
		t.Log("Test OMCI Shema Successed")
	}
}

func BenchmarkOmciRawDataParse(b *testing.B) {
	var logPath string = "./testdata/I240GR.log"
	var outputPath string = "./testoutput/I240GR.json"
	path := OmciRawDataParse(logPath, outputPath, logType)
	if path == "" {
		b.Error("Test OMCI Shema Failed")
	} else {
		b.Log("Test OMCI Shema Successed")
	}
}
