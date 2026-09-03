/*
# ------------------------------------------------------------
# -- mt2XmlShape.go
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
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
)

type mt2XmlShaper struct {
	outputPath      string
	onus            map[string]string
	omciMetaDataMap map[string][]OmciMetaData
}

func parseSerialFromMessage(msg *xmlquery.Node) string {
	// Find the SerialNumber field node
	serialNode := xmlquery.FindOne(msg, `.//field[name/@value="SerialNumber"]`)
	if serialNode == nil {
		return ""
	}

	var vendorID string
	var serialNumber string

	// Try to get VendorId from child field
	vendorNode := xmlquery.FindOne(serialNode, `.//field[name/@value="VendorId"]`)
	if vendorNode != nil {
		vendorValue := vendorNode.SelectAttr("value")
		if vendorValue != "" {
			// Keep vendor ID in ASCII format (e.g., "ADTN")
			vendorID = vendorValue
		}
	}

	// Try to get SerialNumber from child field with the same name
	snNode := xmlquery.FindOne(serialNode, `.//field[name/@value="SerialNumber"]`)
	if snNode != nil {
		snValue := snNode.SelectAttr("value")
		if snValue != "" {
			// Parse hex bytes (e.g., "0x88,0x04,0x42,0xd8" -> "880442d8")
			serialNumber = parseHexString(snValue)
		}
	}

	// If child fields not found, try to parse from parent value attribute
	if vendorID == "" || serialNumber == "" {
		value := serialNode.SelectAttr("value")
		if value != "" {
			// Parse format: {VendorId: ADTN, SerialNumber: 0x88,0x04,0x42,0xd8}
			reVendor := regexp.MustCompile(`VendorId:\s*([A-Za-z0-9]+)`)
			reSerial := regexp.MustCompile(`SerialNumber:\s*(0x[0-9a-fA-F,\s]+)`)

			if matches := reVendor.FindStringSubmatch(value); len(matches) > 1 && vendorID == "" {
				// Keep vendor ID in ASCII format
				vendorID = matches[1]
			}

			if matches := reSerial.FindStringSubmatch(value); len(matches) > 1 && serialNumber == "" {
				serialNumber = parseHexString(matches[1])
			}
		}
	}

	// Return combined vendor ID and serial number
	if vendorID != "" && serialNumber != "" {
		return vendorID + serialNumber
	} else if serialNumber != "" {
		return serialNumber
	}

	return ""
}

func parseHexString(s string) string {
	re := regexp.MustCompile(`0x[0-9a-fA-F]+`)
	matches := re.FindAllString(s, -1)
	if len(matches) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, match := range matches {
		hex := strings.TrimPrefix(match, "0x")
		val, err := parseHexByte(hex)
		if err != nil {
			continue
		}
		sb.WriteString(fmt.Sprintf("%02x", val))
	}

	return sb.String()
}

func parseHexByte(s string) (byte, error) {
	var val byte
	_, err := fmt.Sscanf(s, "%x", &val)
	return val, err
}

func parseTimestampFromMessage(msg *xmlquery.Node, baseTime time.Time) time.Time {
	node := xmlquery.FindOne(msg, "./timestamp[@value]")
	if node == nil {
		return time.Time{}
	}

	timestampStr := node.SelectAttr("value")
	if timestampStr == "" {
		return time.Time{}
	}

	// MT2 XML format: timestamp in milliseconds
	timestampMs, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return time.Time{}
	}

	// Convert to time.Time (milliseconds to nanoseconds)
	timestamp := time.Unix(0, timestampMs*int64(time.Millisecond))

	// If base time is available, use it as reference
	if !baseTime.IsZero() {
		return baseTime.Add(time.Duration(timestampMs) * time.Millisecond)
	}

	return timestamp
}

func parseBaseTimeFromXML(doc *xmlquery.Node) time.Time {
	// Try to parse the startTime from description
	node := xmlquery.FindOne(doc, "//description/startTime")
	if node != nil && node.InnerText() != "" {
		// startTime format: 0000000001657081728492
		// This appears to be microseconds since some epoch
		startTimeStr := node.InnerText()
		if len(startTimeStr) > 16 {
			// Take the significant digits
			startTimeStr = startTimeStr[10:]
		}
		if timestamp, err := strconv.ParseInt(startTimeStr, 10, 64); err == nil {
			return time.Unix(timestamp/1000000, (timestamp%1000000)*1000)
		}
	}

	// Try to parse the date field
	dateNode := xmlquery.FindOne(doc, "//description/date")
	if dateNode != nil && dateNode.InnerText() != "" {
		if t, err := time.Parse(time.RFC3339, dateNode.InnerText()); err == nil {
			return t
		}
	}

	return time.Time{}
}

func (shaper *mt2XmlShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("MT2 XML OMCI shape processing started")

	shaper.outputPath = outputPath
	shaper.onus = map[string]string{}
	shaper.omciMetaDataMap = map[string][]OmciMetaData{}

	// Set default onu name
	var ontid string = "MT2-Onu"

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

	// Parse base time from XML description
	baseTime := parseBaseTimeFromXML(doc)

	messages := xmlquery.Find(doc, "//message")

	// First pass: find serial number
	for _, msg := range messages {
		sn := parseSerialFromMessage(msg)
		if sn != "" {
			ontid = sn
			break
		}
	}

	// Second pass: process all OMCI messages
	for _, msg := range messages {
		nameNode := xmlquery.FindOne(msg, "./name[contains(@value,'OMCI')]")
		if nameNode == nil {
			continue
		}

		rawNode := xmlquery.FindOne(msg, "./raw[@value]")
		if rawNode == nil {
			continue
		}

		rawVal := rawNode.SelectAttr("value")
		if len(rawVal) < 80 {
			continue
		}

		raw := models.StringStrip(rawVal[:80])

		// Extract timestamp
		timestamp := parseTimestampFromMessage(msg, baseTime)
		if timestamp.IsZero() {
			timestamp = time.Now()
		}

		// Initialize ONU data structure if needed
		if _, exists := shaper.onus[ontid]; !exists {
			fileName := shaper.outputPath + ontid
			shaper.onus[ontid] = fileName
			shaper.omciMetaDataMap[ontid] = []OmciMetaData{}
		}

		// Append metadata
		shaper.omciMetaDataMap[ontid] = append(shaper.omciMetaDataMap[ontid], OmciMetaData{
			Timestamp: timestamp,
			RawData:   raw,
		})
	}

	// Flush all data to files
	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)

	utils.Log("MT2 XML OMCI shape processing completed, processed " + fmt.Sprintf("%d", len(shaper.onus)) + " ONUs")

	return shaper.onus
}

func init() {
	shaper := new(mt2XmlShaper)
	shaperRegist(Mt2Xml, shaper)
}
