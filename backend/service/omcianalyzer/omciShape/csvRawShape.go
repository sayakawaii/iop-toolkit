/*
# ------------------------------------------------------------
# -- csvRawShape.go
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
	"strconv"
	"strings"
	"time"
)

type csvRawShaper struct {
	outputPath      string
	onus            map[string]string
	omciMetaDataMap map[string][]OmciMetaData
}

func (shaper *csvRawShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("CSV RAW OMCI shape processing started")
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
		if row == "" || strings.HasPrefix(row, "Row,") {
			continue
		}

		columns := strings.SplitN(row, ",", 3)
		if len(columns) < 3 {
			continue
		}

		timestamp := parseCsvRawTimestamp(strings.TrimSpace(columns[1]))
		if timestamp.IsZero() {
			timestamp = time.Now()
		}

		rawText := strings.TrimSpace(columns[2])
		rawText = strings.TrimPrefix(rawText, "0x")
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
	utils.Log("CSV RAW OMCI shape processing completed, processed " + fmt.Sprintf("%d", len(shaper.onus)) + " ONUs")

	return shaper.onus
}

func parseCsvRawTimestamp(timestampStr string) time.Time {
	if timestampStr == "" {
		return time.Time{}
	}

	// CSV RAW format uses milliseconds since capture start.
	ms, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return time.Time{}
	}

	if ms == 0 {
		return time.Time{}
	}

	return time.Unix(0, ms*int64(time.Millisecond))
}

func init() {
	shaper := new(csvRawShaper)
	shaperRegist(CsvRaw, shaper)
}
