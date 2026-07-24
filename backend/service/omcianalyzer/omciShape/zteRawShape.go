/*
# ------------------------------------------------------------
# -- zteRawShape.go
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

type zteRawShaper struct {
	outputPath      string
	onus            map[string]string
	omciMetaDataMap map[string][]OmciMetaData
}

func (shaper *zteRawShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("ZTE RAW OMCI shape processing started")
	shaper.outputPath = outputPath
	shaper.onus = map[string]string{}
	shaper.omciMetaDataMap = map[string][]OmciMetaData{}
	ontid := "ZTE-ONU"
	file, err := os.Open(logPath)
	if err != nil {
		utils.Log("open file failed: " + fmt.Sprintf("%s", err))
		return nil
	}

	defer file.Close()
	if _, exists := shaper.onus[ontid]; !exists {
		fileName := shaper.outputPath + ontid
		shaper.onus[ontid] = fileName
		shaper.omciMetaDataMap[ontid] = []OmciMetaData{}
	}

	reHex := regexp.MustCompile(`[0-9A-Fa-f]+`)
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
		matches := reHex.FindAllString(str, -1)
		if len(matches) < 40 {
			continue
		}

		hexString := strings.Join(matches[:40], "")
		raw := models.StringStrip(hexString)
		if len(raw) < 80 {
			continue
		}

		shaper.omciMetaDataMap[ontid] = append(shaper.omciMetaDataMap[ontid], OmciMetaData{
			Timestamp: time.Time{},
			RawData:   raw,
		})
	}

	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)
	utils.Log("ZTE RAW OMCI shape processing completed, processed " + fmt.Sprintf("%d", len(shaper.onus)) + " ONUs")

	return shaper.onus
}

func init() {
	shaper := new(zteRawShaper)
	shaperRegist(ZteRaw, shaper)
}
