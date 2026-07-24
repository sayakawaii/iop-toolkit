/*
# ------------------------------------------------------------
# -- xponHwaShape.go
# --
# -- Extract OMCI messages from xponhwa logs
# -- Based on nokiaOnuShape.go
# -- 2024
# ------------------------------------------------------------
*/

package omciShape

import (
	"fmt"
	"io"
	"omciAnalyzer/models"
	"omciAnalyzer/utils"
	"os"
	"regexp"
	"strings"
	"time"
)

type xponHwaShaper struct {
	cachedRegex     *regexp.Regexp
	cachedLayout    string
	outputPath      string
	onus            map[string]string
	omciMetaDataMap map[string][]OmciMetaData
}

func (shaper *xponHwaShaper) getTimestamp(s string) time.Time {
	// example: [10/09/1973-18:58:31.660766],[OMCI:OmciTxRx],omci_packet TX: ...
	// Try to use cached regex and layout first
	if shaper.cachedRegex != nil {
		if match := shaper.cachedRegex.FindString(s); match != "" {
			timeStr := strings.Trim(match, "[]")
			if t, err := time.Parse(shaper.cachedLayout, timeStr); err == nil {
				return t
			}
			shaper.cachedRegex = nil
			shaper.cachedLayout = ""
		}
	}

	// Pattern for xponhwa timestamp: [DD/MM/YYYY-HH:MM:SS.mmmmmm]
	patterns := []struct {
		regex  string
		layout string
	}{
		{`\[\d{2}/\d{2}/\d{4}-\d{2}:\d{2}:\d{2}\.\d{6}\]`, "02/01/2006-15:04:05.000000"},
		{`\d{2}/\d{2}/\d{4}-\d{2}:\d{2}:\d{2}\.\d{6}`, "02/01/2006-15:04:05.000000"},
		{`\d{2}/\d{2}/\d{2}\s+\d{2}:\d{2}:\d{2}\.\d{6}`, "02/01/06 15:04:05.000000"},
		{`\d{4}-\d{2}-\d{2}\s+\d{2}:\d{2}:\d{2}`, "2006-01-02 15:04:05"},
	}

	for _, p := range patterns {
		re := regexp.MustCompile(p.regex)
		if match := re.FindString(s); match != "" {
			timeStr := strings.Trim(match, "[]")
			if t, err := time.Parse(p.layout, timeStr); err == nil {
				shaper.cachedRegex = re
				shaper.cachedLayout = p.layout
				return t
			}
		}
	}

	return time.Time{}
}

func (shaper *xponHwaShaper) getOnuId(s string) string {
	// Extract ONUid from line like: ONUid[1] or OnuId[2]
	onuIdPattern := regexp.MustCompile(`(?i)ONUid\[(\d+)\]`)
	onuIdMatches := onuIdPattern.FindStringSubmatch(s)
	if len(onuIdMatches) < 2 {
		return ""
	}

	ontid := onuIdMatches[1]
	fileName := shaper.outputPath + ontid

	if _, exists := shaper.onus[ontid]; !exists {
		shaper.onus[ontid] = fileName
		shaper.omciMetaDataMap[ontid] = []OmciMetaData{}
	}

	return ontid
}

func extractOMCIFromXponHwaLine(str string) (hexOut string, onuId string, ok bool) {
	// Pattern: omci_packet TX/RX: ... ONUid[1] or OnuId[2] followed by hex data
	// Extract ONUid and hex data

	// Find ONUid (case-insensitive to support both ONUid and OnuId)
	onuIdPattern := regexp.MustCompile(`(?i)ONUid\[(\d+)\]`)
	onuIdMatches := onuIdPattern.FindStringSubmatch(str)
	if len(onuIdMatches) < 2 {
		return "", "", false
	}

	onuId = onuIdMatches[1]

	// Find hex data after ONUid[X]
	onuIdIndex := strings.Index(str, onuIdMatches[0])
	if onuIdIndex == -1 {
		return "", "", false
	}

	// Start searching for hex data after ONUid[X]
	startPos := onuIdIndex + len(onuIdMatches[0])
	hexStartStr := str[startPos:]

	// Extract all hex values (space-separated)
	// Pattern: 34 c7 49 0a 00 02 ...
	hexPattern := regexp.MustCompile(`([0-9A-Fa-f]{2})`)
	hexMatches := hexPattern.FindAllString(hexStartStr, -1)

	if len(hexMatches) < 40 {
		return "", "", false
	}

	// Take only first 40 bytes (80 hex chars)
	hexString := strings.Join(hexMatches[:40], "")

	return hexString, onuId, true
}

func (shaper *xponHwaShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("xponhwa omci shape")
	shaper.outputPath = outputPath
	shaper.onus = map[string]string{}
	shaper.omciMetaDataMap = map[string][]OmciMetaData{}

	file, err := os.Open(logPath)
	if err != nil {
		utils.Log("open file failed: " + fmt.Sprintf("%s", err))
		return nil
	}
	defer file.Close()

	reader := models.GetReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			utils.Log("end of read")
			break
		}
		if len(line) == 0 {
			continue
		}

		str := string(line)

		// Check if this line contains OMCI data
		// Pattern: omci_packet TX/RX: ... ONUid[X] or OnuId[X] followed by hex data
		if strings.Contains(str, "omci_packet") && (strings.Contains(str, "ONUid[") || strings.Contains(str, "OnuId[")) {
			hexData, onuId, ok := extractOMCIFromXponHwaLine(str)
			if !ok || onuId == "" {
				continue
			}

			// Initialize ONUid if not exists
			shaper.getOnuId(str)

			// Extract timestamp
			timeStamp := shaper.getTimestamp(str)

			// Store metadata
			shaper.omciMetaDataMap[onuId] = append(shaper.omciMetaDataMap[onuId], OmciMetaData{
				Timestamp: timeStamp,
				RawData:   hexData,
			})
		}
	}

	// Flush all OMCI data to files
	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)

	return shaper.onus
}

func init() {
	shaper := new(xponHwaShaper)
	shaperRegist(XponHwa, shaper)
}
