/*
# ------------------------------------------------------------
# -- hisenseOnuSahpe.go
# --
# -- Huang Minghe
# -- 2022-8-16
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

type hisenseOnuShaper struct {
	outputPath      string
	omciMetaDataMap map[string][]OmciMetaData
}

// getTimestamp extracts timestamp from log line
// Currently returns zero time as the log format doesn't contain timestamp information
// Reserved for future timestamp extraction when log format is updated
func (shaper *hisenseOnuShaper) getTimestamp(s string) time.Time {
	// Placeholder for future timestamp extraction logic
	// When timestamp format is available in the log, implement pattern matching here
	// Example patterns to consider:
	// - [DD/MM/YY HH:MM:SS.mmmmmm]
	// - [YYYY-MM-DD HH:MM:SS]
	// etc.

	// For now, return zero time
	return time.Time{}
}

func (shaper *hisenseOnuShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("hisense onu omci shape")
	shaper.outputPath = outputPath
	onus := map[string]string{}
	shaper.omciMetaDataMap = map[string][]OmciMetaData{}

	var ontid string = "HisenseOnu"
	var omciBuff string = ""

	file, err := os.Open(logPath)
	if err != nil {
		utils.Log("open file failed: " + fmt.Sprintf("%s", err))
		return nil
	}
	defer file.Close()

	// Initialize file mapping (without .log extension to match pattern)
	fileName := outputPath + ontid
	onus[ontid] = fileName
	shaper.omciMetaDataMap[ontid] = []OmciMetaData{}

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
		if strings.Contains(str, "[INFO]: ") {
			// Extract hex data after "[INFO]: "
			startIndex := strings.Index(str, "[INFO]: ")
			if startIndex == -1 {
				continue
			}
			startIndex = startIndex + len("[INFO]: ")

			// Get all hex characters from this position
			hexData := strings.TrimSpace(str[startIndex:])
			matches := regexp.MustCompile(`[0-9A-Fa-f]+`).FindAllString(hexData, -1)
			if len(matches) < 1 {
				continue
			}

			// Concatenate all hex strings
			hexString := strings.Join(matches, "")
			if len(hexString) < 80 {
				continue
			}
			hexString = hexString[:80]
			omciBuff = models.StringStrip(hexString)

			if len(omciBuff) == 80 {
				// Get timestamp (currently returns zero time)
				timestamp := shaper.getTimestamp(str)

				// Store metadata
				shaper.omciMetaDataMap[ontid] = append(shaper.omciMetaDataMap[ontid], OmciMetaData{
					Timestamp: timestamp,
					RawData:   omciBuff,
				})
			}
			omciBuff = ""
		}
	}

	// Flush all OMCI data to JSON files
	flushOmciDataToFile(onus, shaper.omciMetaDataMap)

	return onus
}

func init() {
	shaper := new(hisenseOnuShaper)
	shaperRegist(HisenseOnu, shaper)
}
