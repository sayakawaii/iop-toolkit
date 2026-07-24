/*
# ------------------------------------------------------------
# -- gtcXmlShape.go
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

type gtcXmlShaper struct {
	outputPath      string
	onus            map[string]string
	omciMetaDataMap map[string][]OmciMetaData
}

func (shaper *gtcXmlShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("GTC XML OMCI shape processing started")
	shaper.outputPath = outputPath
	shaper.onus = map[string]string{}
	shaper.omciMetaDataMap = map[string][]OmciMetaData{}

	// set default ONU name
	ontid := "GTC-Onu"

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

	if node := xmlquery.FindOne(doc, "//PLOAMuSerialNumberOnu"); node != nil {
		if decoded := node.SelectAttr("serialNumberDecoded"); decoded != "" {
			ontid = decoded
		}
	}

	nodes := xmlquery.Find(doc, "//OMCI-DOWNSTREAM | //OMCI-UPSTREAM")
	for _, n := range nodes {
		bitstream := n.SelectAttr("BitStream")
		if len(bitstream) < 80 {
			continue
		}

		raw := models.StringStrip(bitstream[:80])

		timestamp := parseGtcTimestamp(n.SelectAttr("timeStamp"))
		if timestamp.IsZero() {
			timestamp = time.Now()
		}

		onuKey := ontid
		if onuKey == "" || onuKey == "GTC-Onu" {
			if onuID := n.SelectAttr("onuId"); onuID != "" {
				onuKey = "ONU-" + onuID
			}
		}

		if _, exists := shaper.onus[onuKey]; !exists {
			fileName := shaper.outputPath + onuKey
			shaper.onus[onuKey] = fileName
			shaper.omciMetaDataMap[onuKey] = []OmciMetaData{}
		}

		shaper.omciMetaDataMap[onuKey] = append(shaper.omciMetaDataMap[onuKey], OmciMetaData{
			Timestamp: timestamp,
			RawData:   raw,
		})
	}

	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)
	utils.Log("GTC XML OMCI shape processing completed, processed " + fmt.Sprintf("%d", len(shaper.onus)) + " ONUs")

	return shaper.onus
}

func parseGtcTimestamp(timestampStr string) time.Time {
	if timestampStr == "" {
		return time.Time{}
	}
	parts := strings.Split(timestampStr, ":")
	if len(parts) < 3 {
		return time.Time{}
	}

	values := make([]int64, 0, len(parts))
	for _, part := range parts {
		v, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return time.Time{}
		}
		values = append(values, v)
	}

	// Format observed: H:MM:SS:ms:us:ns (some parts may be missing)
	hours := values[0]
	minutes := values[1]
	seconds := values[2]
	var milliseconds int64
	var microseconds int64
	var nanoseconds int64

	if len(values) > 3 {
		milliseconds = values[3]
	}
	if len(values) > 4 {
		microseconds = values[4]
	}
	if len(values) > 5 {
		nanoseconds = values[5]
	}

	duration := time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second +
		time.Duration(milliseconds)*time.Millisecond +
		time.Duration(microseconds)*time.Microsecond +
		time.Duration(nanoseconds)*time.Nanosecond

	if duration == 0 {
		return time.Time{}
	}

	return time.Unix(0, 0).Add(duration)
}

func init() {
	shaper := new(gtcXmlShaper)
	shaperRegist(GtcXml, shaper)
}
