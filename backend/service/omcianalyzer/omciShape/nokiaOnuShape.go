/*
# ------------------------------------------------------------
# -- nokiaOnuShape.go
# --
# -- Huang Minghe
# -- 2022-8-16
# ------------------------------------------------------------
*/

package omciShape

import (
	"bufio"
	"fmt"
	"io"
	"omciAnalyzer/models"
	"omciAnalyzer/utils"
	"os"
	"regexp"
	"strings"
	"time"
)

type nokiaOnuShaper struct {
	cachedRegex     *regexp.Regexp
	cachedLayout    string
	outputPath      string
	onus            map[string]string
	omciMetaDataMap map[string][]OmciMetaData
}

func (shaper *nokiaOnuShaper) getTimestamp(s string) time.Time {
	// example: [01-19 01:39:47] or [14:20:11:714] or [12-31 20:01:50]
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

	// Pattern for Nokia ONU timestamps
	patterns := []struct {
		regex  string
		layout string
	}{
		{`\[\d{2}-\d{2}\s+\d{2}:\d{2}:\d{2}\]`, "01-02 15:04:05"},
		{`\[\d{2}:\d{2}:\d{2}:\d{3}\]`, "15:04:05:000"},
		{`\[\d{2}-\d{2}-\d{2}\s+\d{2}:\d{2}:\d{2}\]`, "01-02-06 15:04:05"},
		{`\[\d{2}/\d{2}/\d{2}\s+\d{2}:\d{2}:\d{2}\.\d{6}\]`, "02/01/06 15:04:05.000000"},
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

type ProcessAction int

const (
	ActionNormal ProcessAction = iota
	ActionContinue
	ActionBreak
)

func getExtendedOmciContentLength(matches []string) int {
	var contentLength int
	hexString := strings.Join(matches[:], "")
	hexString = models.StringStrip(hexString)
	contentLengthHex := hexString[16:20]
	fmt.Sscanf(contentLengthHex, "%x", &contentLength)
	expectedLen := 20 + contentLength*2

	return expectedLen
}

func processOMCILine(str string, reader *bufio.Reader) (hexOut string, action ProcessAction, omciLen int, ok bool) {
	if !strings.Contains(str, "ms-") {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			utils.Log("end of read")
			return "", ActionBreak, 0, false
		}
		str += string(line)
		if !strings.Contains(str, "ms-") {
			return "", ActionContinue, 0, false
		}
	}

	var matches []string
	searchStartIndex := 0

	for {
		msIndex := strings.Index(str[searchStartIndex:], "ms-")
		if msIndex == -1 {
			return "", ActionContinue, 0, false
		}

		startIndex := searchStartIndex + msIndex + 3
		endMarkIndex := strings.Index(str[startIndex:], "[")
		var searchStr string
		if endMarkIndex != -1 {
			searchStr = str[startIndex : startIndex+endMarkIndex]
		} else {
			searchStr = str[startIndex:]
		}
		matches = regexp.MustCompile(`[0-9A-Fa-f]+`).FindAllString(searchStr, -1)

		if len(matches) > 12 {
			if (matches[3] == "0A" || matches[3] == "0a") && len(matches) >= 40 {
				break
			} else if matches[3] == "0B" || matches[3] == "0b" {
				break
			}
		}

		searchStartIndex = startIndex + len(searchStr)
	}

	hexString := strings.Join(matches[:], "")
	hexString = models.StringStrip(hexString)

	if len(hexString) < 8 {
		return "", ActionContinue, 0, false
	}

	omciType := hexString[6:8]
	var expectedLen int

	switch omciType {
	case "0B", "0b":
		if len(hexString) < 20 {
			return "", ActionContinue, 0, false
		}
		expectedLen = getExtendedOmciContentLength(matches)
		if expectedLen < 80 {
			expectedLen = 80
		}
	case "0A", "0a":
		expectedLen = 80
	default:
		return "", ActionContinue, 0, false
	}

	if len(hexString) < expectedLen {
		//pading to expected length
		paddingLen := expectedLen - len(hexString)
		hexString += strings.Repeat("0", paddingLen)
	}

	hexString = hexString[:expectedLen]
	return hexString, ActionNormal, expectedLen, true
}

func (shaper *nokiaOnuShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("onu omci shape")
	shaper.outputPath = outputPath
	shaper.onus = map[string]string{}
	shaper.omciMetaDataMap = map[string][]OmciMetaData{}

	var ontid string = "NOKIA-ONU"
	var omciBuff string = ""
	var fileName string = ""
	file, err := os.Open(logPath)
	if err != nil {
		utils.Log("open file failed: " + fmt.Sprintf("%s", err))
		return nil
	}

	fileName = outputPath + ontid
	if _, res := shaper.onus[ontid]; !res {
		shaper.onus[ontid] = fileName
		shaper.omciMetaDataMap[ontid] = []OmciMetaData{}
	}

	defer file.Close()
	reader := models.GetReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			utils.Log("end of read")
			break
		}
		if len(line) != 0 {
			str := string(line)
			var (
				hexProcessed string
				action       ProcessAction
				omciLen      int
				ok           bool
			)

			if strings.Contains(str, "[OMCI]OMCI_TX#") || strings.Contains(str, "OMCI_TX#") || strings.Contains(str, "parser_TX#") {
				hexProcessed, action, omciLen, ok = processOMCILine(str, reader)
			} else if strings.Contains(str, "[OMCI]OMCI_RX#") || strings.Contains(str, "OMCI_RX#") || strings.Contains(str, "parser_RX#") {
				hexProcessed, action, omciLen, ok = processOMCILine(str, reader)
			} else {
				continue
			}
			if action == ActionBreak {
				break
			}
			if action == ActionContinue || !ok {
				continue
			}
			omciBuff += hexProcessed
			if len(omciBuff) == omciLen {
				// Extract timestamp from current line
				timeStamp := shaper.getTimestamp(str)

				// Store metadata with timestamp
				shaper.omciMetaDataMap[ontid] = append(shaper.omciMetaDataMap[ontid], OmciMetaData{
					Timestamp: timeStamp,
					RawData:   omciBuff,
				})
				omciBuff = ""
			}
		}
	}

	// Flush all OMCI data to files
	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)

	return shaper.onus
}

func init() {
	shaper := new(nokiaOnuShaper)
	shaperRegist(NokiaOnu, shaper)
}
