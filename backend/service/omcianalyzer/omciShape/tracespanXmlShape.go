/*
# ------------------------------------------------------------
# -- tracespanXmlShape.go
# --
# -- Huang Minghe
# -- 2022-8-16
# ------------------------------------------------------------
*/

package omciShape

import (
	"fmt"
	"log"
	"omciAnalyzer/models"
	"omciAnalyzer/utils"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
)

type tracespanXmlShaper struct {
	outputPath      string
	onus            map[string]string
	omciMetaDataMap map[string][]OmciMetaData
}

func (shaper *tracespanXmlShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("TraceSpan XML OMCI shape processing started")
	shaper.outputPath = outputPath
	shaper.onus = map[string]string{}
	shaper.omciMetaDataMap = map[string][]OmciMetaData{}

	// set default ONU name
	ontid := "TraceSpanOnu"

	data, err := os.ReadFile(logPath)
	if err != nil {
		log.Fatal(err)
	}

	if !isLikelyValidXML(data) {
		utils.Log("not valid xml format")
		return shaper.onus
	}
	data = xmlDataClean(data)
	doc, err := xmlquery.Parse(strings.NewReader(string(data)))
	if err != nil {
		utils.Log("[WARN] invalid XML: ", err)
		return shaper.onus
	}
	if serial := parseTraceSpanSerial(doc); serial != "" {
		ontid = serial
	}

	for _, entry := range xmlquery.Find(doc, "//entry") {
		rawNode := xmlquery.FindOne(entry, "./Raw_Data")
		if rawNode == nil {
			continue
		}
		rawText := strings.TrimSpace(rawNode.InnerText())
		rawText = strings.TrimPrefix(rawText, "0x")
		if len(rawText) < 80 {
			continue
		}
		raw := models.StringStrip(rawText[:80])

		timestampNode := xmlquery.FindOne(entry, "./paramaters/Timestamp_Start_of_capture")
		timestamp := parseTraceSpanTimestamp(timestampNode)
		if timestamp.IsZero() {
			timestamp = time.Now()
		}

		if _, exists := shaper.onus[ontid]; !exists {
			fileName := shaper.outputPath + ontid
			shaper.onus[ontid] = fileName
			shaper.omciMetaDataMap[ontid] = []OmciMetaData{}
		}

		shaper.omciMetaDataMap[ontid] = append(shaper.omciMetaDataMap[ontid], OmciMetaData{
			Timestamp: timestamp,
			RawData:   raw,
		})
	}

	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)
	utils.Log("TraceSpan XML OMCI shape processing completed, processed " + fmt.Sprintf("%d", len(shaper.onus)) + " ONUs")

	return shaper.onus
}

func parseTraceSpanSerial(doc *xmlquery.Node) string {
	vendorNode := xmlquery.FindOne(doc, `//paramater[@Name="Serial number.Vendor Id"]`)
	snNode := xmlquery.FindOne(doc, `//paramater[@Name="Serial number.SN"]`)

	vendor := ""
	serial := ""
	if vendorNode != nil {
		vendor = strings.TrimSpace(vendorNode.InnerText())
	}
	if snNode != nil {
		serial = parseHexString(strings.TrimSpace(snNode.InnerText()))
	}

	if vendor != "" && serial != "" {
		return vendor + serial
	}
	if serial != "" {
		return serial
	}

	return ""
}

func parseTraceSpanTimestamp(node *xmlquery.Node) time.Time {
	if node == nil {
		return time.Time{}
	}
	timestampStr := strings.TrimSpace(node.InnerText())
	if timestampStr == "" {
		return time.Time{}
	}

	parts := strings.SplitN(timestampStr, ".", 2)
	timeParts := strings.Split(parts[0], ":")
	if len(timeParts) != 3 {
		return time.Time{}
	}

	hours, err := strconv.ParseInt(timeParts[0], 10, 64)
	if err != nil {
		return time.Time{}
	}
	minutes, err := strconv.ParseInt(timeParts[1], 10, 64)
	if err != nil {
		return time.Time{}
	}
	seconds, err := strconv.ParseInt(timeParts[2], 10, 64)
	if err != nil {
		return time.Time{}
	}

	var milliseconds int64
	if len(parts) > 1 {
		msText := parts[1]
		if len(msText) > 3 {
			msText = msText[:3]
		}
		for len(msText) < 3 {
			msText += "0"
		}
		if ms, err := strconv.ParseInt(msText, 10, 64); err == nil {
			milliseconds = ms
		}
	}

	duration := time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second +
		time.Duration(milliseconds)*time.Millisecond

	if duration == 0 {
		return time.Time{}
	}

	return time.Unix(0, 0).Add(duration)
}

func init() {
	shaper := new(tracespanXmlShaper)
	shaperRegist(TraceSpanXml, shaper)
}
