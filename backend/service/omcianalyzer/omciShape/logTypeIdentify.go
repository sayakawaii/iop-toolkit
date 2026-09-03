/*
# ------------------------------------------------------------
# -- logTypeIdentify.go
# --
# -- Huang Minghe
# -- 2022-4-30
# ------------------------------------------------------------
*/

package omciShape

import (
	"bufio"
	"fmt"
	"io"
	"omciAnalyzer/utils"
	"os"
	"path"
	"strings"
)

const (
	Isam         = 1
	LightSpan    = 2
	NokiaOnu     = 3
	CsvRaw       = 4
	Mt2Xml       = 5
	GtcXml       = 6
	sstRaw       = 7
	CopperOnu    = 8
	ZteRaw       = 9
	ZteOnu       = 10
	HuaweiOnu    = 11
	HisenseOnu   = 12
	TraceSpanXml = 13
	XponHwa      = 14
	Pcap         = 15
	Mt2Html      = 16
	CalixOnu     = 17
	Invalid      = 18
)

var logTypeDef map[uint8]string = map[uint8]string{
	1:  "Isam",
	2:  "LightSpan",
	3:  "NokiaOnu",
	4:  "CsvRaw",
	5:  "Mt2Xml",
	6:  "GtcXml",
	7:  "sstRaw",
	8:  "CopperOnu",
	9:  "ZteRaw",
	10: "ZteOnu",
	11: "HuaweiOnu",
	12: "HisenseOnu",
	13: "TraceSpanXml",
	14: "XponHwa",
	15: "Pcap",
	16: "Mt2Html",
	17: "CalixOnu",
	18: "Invalid",
}

func LogTypeNameGet(logType uint8) string {
	return logTypeDef[logType]
}

type logIdentifier struct {
	name    string
	check   func(filename string, line string) bool
	logType uint8
}

var extensionIdentifiers = []logIdentifier{
	{name: "PCAP", logType: Pcap, check: func(f, _ string) bool { return path.Ext(f) == ".pcap" || path.Ext(f) == ".pcapng" }},
	{name: "ZTE Raw", logType: ZteRaw, check: func(f, _ string) bool { return path.Ext(f) == ".raw" }},
}

var contentIdentifiers = []logIdentifier{
	{name: "csv Raw", logType: CsvRaw, check: func(f, l string) bool {
		return path.Ext(f) == ".csv" && strings.Contains(l, "Row, Timestamp, RawData")
	}},
	{name: "mt2 Xml", logType: Mt2Xml, check: func(f, l string) bool {
		ext := path.Ext(f)
		return (ext == ".xml" || ext == ".xgspon" || ext == ".gpon") && strings.Contains(l, "<mt2-report>")
	}},
	{name: "gtc Xml", logType: GtcXml, check: func(f, l string) bool {
		return path.Ext(f) == ".xml" && strings.Contains(l, "<GPONDOCTOR-TRACE>")
	}},
	{name: "tracespan Xml", logType: TraceSpanXml, check: func(f, l string) bool {
		return path.Ext(f) == ".xml" && strings.Contains(l, "<traceReport")
	}},
	{name: "mt2 Html", logType: Mt2Html, check: func(f, l string) bool {
		return path.Ext(f) == ".html" && strings.Contains(l, "<tr onclick=\"drawInfo(")
	}},
	// Match real ISAM hex-dump lines (<OMCI MSG> ... OntId), not diagnostic "OMCI MSG:" events.
	{name: "ISAM", logType: Isam, check: func(_, l string) bool { return strings.Contains(l, "<OMCI MSG>") }},
	{name: "LightSpan", logType: LightSpan, check: func(_, l string) bool { return strings.Contains(l, "Dir: Tx --> Onu:") }},
	{name: "XponHwa", logType: XponHwa, check: func(_, l string) bool {
		return strings.Contains(l, "[OMCI:OmciTxRx]") && strings.Contains(l, "omci_packet") &&
			(strings.Contains(l, "ONUid[") || strings.Contains(l, "OnuId["))
	}},
	{name: "NokiaOnu", logType: NokiaOnu, check: func(_, l string) bool {
		return strings.Contains(l, "[OMCI]OMCI_RX#") ||
			strings.Contains(l, "OmciMain::") ||
			strings.Contains(l, "OmciMsgRxTx::") ||
			strings.Contains(l, "OMCI_RX#") ||
			strings.Contains(l, "parser_RX#")
	}},
	{name: "json Raw", logType: sstRaw, check: func(f, l string) bool {
		return path.Ext(f) == ".json" &&
			(strings.Contains(strings.ToLower(l), "tx") || strings.Contains(strings.ToLower(l), "rx"))
	}},
	{name: "copper ONU", logType: CopperOnu, check: func(_, l string) bool { return strings.Contains(l, "RX_ONT_MSG --") }},
	{name: "ZTE ONU", logType: ZteOnu, check: func(_, l string) bool { return strings.Contains(l, "Recive packet:") }},
	{name: "Huawei ONU", logType: HuaweiOnu, check: func(_, l string) bool {
		return strings.Contains(l, "OLT--------->ONT:") || strings.Contains(l, "ONT--------->OLT:")
	}},
	{name: "CalixOnu", logType: CalixOnu, check: func(_, l string) bool {
		return strings.HasPrefix(l, "OMCCRX[") || strings.HasPrefix(l, "OMCCTX[")
	}},
	{name: "HisenseOnu", logType: HisenseOnu, check: func(_, l string) bool {
		return strings.Contains(l, "[DEBUG]: Transaction correlation identifier:")
	}},
}

func LogTypeAutoIdentify(logPath string) uint8 {
	utils.Log("log type auto identify")

	file, err := os.Open(logPath)
	if err != nil {
		utils.Log("open file failed: " + fmt.Sprintf("%s", err))
		return Invalid
	}
	defer file.Close()

	filename := file.Name()

	for _, id := range extensionIdentifiers {
		if id.check(filename, "") {
			utils.Log("this is a " + id.name + " log")
			return id.logType
		}
	}

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				utils.Log("end of read")
			} else {
				utils.Log("read file failed: " + fmt.Sprintf("%s", err))
			}
			break
		}

		if len(line) == 0 {
			continue
		}

		str := string(line)
		for _, id := range contentIdentifiers {
			if id.check(filename, str) {
				utils.Log("this is a " + id.name + " log")
				return id.logType
			}
		}
	}

	return Invalid
}
