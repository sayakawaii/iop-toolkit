/*
# ------------------------------------------------------------
# -- sstRawShape.go
# --
# -- zixhu
# -- 2023-5-31
# ------------------------------------------------------------
*/

package omciShape

import (
	"fmt"
	"io"
	"omciAnalyzer/models"
	"omciAnalyzer/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

type sstRawShaper struct {
	outputPath      string
	onus            map[string]string
	omciMetaDataMap map[string][]OmciMetaData
}

func (shaper *sstRawShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("SST RAW OMCI shape processing started")
	shaper.outputPath = outputPath
	shaper.onus = map[string]string{}
	shaper.omciMetaDataMap = map[string][]OmciMetaData{}

	// set default ONU name
	ontid := "AONT"
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

		row := strings.TrimSpace(string(line))
		if row == "" {
			continue
		}

		row = strings.TrimPrefix(row, "[")
		row = strings.TrimSuffix(row, "]")
		row = strings.TrimSuffix(row, ",")

		parts := strings.SplitN(row, ",", 3)
		if len(parts) < 3 {
			continue
		}

		timestamp := parseSstTimestamp(strings.TrimSpace(parts[0]))
		if timestamp.IsZero() {
			timestamp = time.Now()
		}

		rawText := strings.TrimSpace(parts[2])
		rawText = strings.Trim(rawText, "\"")
		if len(rawText) < 80 {
			continue
		}
		raw := models.StringStrip(rawText[:80])

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
	utils.Log("SST RAW OMCI shape processing completed, processed " + fmt.Sprintf("%d", len(shaper.onus)) + " ONUs")

	return shaper.onus
}

func parseSstTimestamp(timestampStr string) time.Time {
	timestampStr = strings.Trim(timestampStr, "\"")
	if timestampStr == "" {
		return time.Time{}
	}

	seconds, err := strconv.ParseFloat(timestampStr, 64)
	if err != nil {
		return time.Time{}
	}

	if seconds == 0 {
		return time.Time{}
	}

	secPart := int64(seconds)
	nanoPart := int64((seconds - float64(secPart)) * float64(time.Second))
	return time.Unix(secPart, nanoPart)
}

func init() {
	shaper := new(sstRawShaper)
	shaperRegist(sstRaw, shaper)
}
