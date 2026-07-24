/*
# ------------------------------------------------------------
# -- zteOnuShape.go
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

type zteOnuShaper struct {
	outputPath      string
	onus            map[string]string
	omciMetaDataMap map[string][]OmciMetaData
}

func (shaper *zteOnuShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("ZTE ONU OMCI shape processing started")
	shaper.outputPath = outputPath
	shaper.onus = map[string]string{}
	shaper.omciMetaDataMap = map[string][]OmciMetaData{}

	// set default ONU name
	ontid := "ZTE-ONU"
	omciBuff := ""

	file, err := os.Open(logPath)
	if err != nil {
		utils.Log("open file failed: " + fmt.Sprintf("%s", err))
		return nil
	}
	defer file.Close()

	// Initialize ONU entry
	fileName := shaper.outputPath + ontid
	if _, exists := shaper.onus[ontid]; !exists {
		shaper.onus[ontid] = fileName
		shaper.omciMetaDataMap[ontid] = []OmciMetaData{}
	}

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
		if !strings.Contains(str, "Recive packet:") && !strings.Contains(str, "Send packet:") {
			continue
		}

		startIndex := strings.LastIndex(str, "packet:")
		if startIndex == -1 {
			continue
		}
		startIndex = startIndex + 7

		matches := regexp.MustCompile(`[0-9A-Fa-f]+`).FindAllString(string(str[startIndex:]), -1)
		if len(matches) < 1 {
			continue
		}

		hexString := matches[0]
		if len(hexString) < 80 {
			continue
		}
		hexString = hexString[:80]

		omciBuff = omciBuff + models.StringStrip(hexString)
		if len(omciBuff) == 80 {
			// TODO: Parse timestamp from log line in the future
			// For now, use zero time as default
			timestamp := time.Time{}

			shaper.omciMetaDataMap[ontid] = append(shaper.omciMetaDataMap[ontid], OmciMetaData{
				Timestamp: timestamp,
				RawData:   omciBuff,
			})
		}
		omciBuff = ""
	}

	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)
	utils.Log("ZTE ONU OMCI shape processing completed, processed " + fmt.Sprintf("%d", len(shaper.onus)) + " ONUs")

	return shaper.onus
}

func init() {
	shaper := new(zteOnuShaper)
	shaperRegist(ZteOnu, shaper)
}
