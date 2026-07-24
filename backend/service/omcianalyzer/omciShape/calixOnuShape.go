/*
# ------------------------------------------------------------
# -- calixOnuShape.go
# --
# -- Calix ONU OMCI log shaper
# -- Log format:
# --   OMCCRX[<seconds>]:
# --   <16 bytes hex>
# --   <16 bytes hex>
# --   <8 bytes hex>
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

type calixOnuShaper struct {
	outputPath      string
	state           OmciState
	onus            map[string]string
	omciBuff        string
	timeStamp       time.Time
	omciMetaDataMap map[string][]OmciMetaData
}

const (
	CalixOnuLineInit OmciState = iota
	CalixOnuLineData1
	CalixOnuLineData2
	CalixOnuLineData3
)

const defaultCalixOnuID = "CalixOnu"

func (shaper *calixOnuShaper) getTimestamp(s string) time.Time {
	// Format: OMCCRX[83.156922]: or OMCCTX[83.157627]:
	leftIdx := strings.Index(s, "[")
	rightIdx := strings.Index(s, "]")
	if leftIdx < 0 || rightIdx < 0 || rightIdx <= leftIdx {
		return time.Time{}
	}
	tsStr := s[leftIdx+1 : rightIdx]
	seconds, err := strconv.ParseFloat(tsStr, 64)
	if err != nil {
		return time.Time{}
	}
	return time.Unix(0, int64(seconds*1e9))
}

func (shaper *calixOnuShaper) matchHeader(s string) bool {
	return strings.HasPrefix(s, "OMCCRX[") || strings.HasPrefix(s, "OMCCTX[")
}

func (shaper *calixOnuShaper) initOnu(outputPath string) {
	fileName := outputPath + defaultCalixOnuID
	if err := os.WriteFile(fileName, []byte{}, 0666); err != nil {
		utils.Log("truncate file failed: " + err.Error())
	}
	if _, exists := shaper.onus[defaultCalixOnuID]; !exists {
		shaper.onus[defaultCalixOnuID] = fileName
		shaper.omciMetaDataMap[defaultCalixOnuID] = []OmciMetaData{}
	}
}

func (shaper *calixOnuShaper) processLine(s string) {
	switch shaper.state {
	case CalixOnuLineInit:
		shaper.handleInit(s)
	case CalixOnuLineData1:
		shaper.handleData(s, CalixOnuLineData2)
	case CalixOnuLineData2:
		shaper.handleData(s, CalixOnuLineData3)
	case CalixOnuLineData3:
		shaper.handleData(s, CalixOnuLineInit)
	}
}

func (shaper *calixOnuShaper) handleInit(s string) {
	trimmed := strings.TrimSpace(s)
	if !shaper.matchHeader(trimmed) {
		return
	}
	shaper.timeStamp = shaper.getTimestamp(trimmed)
	shaper.state = CalixOnuLineData1
	shaper.omciBuff = ""
}

func (shaper *calixOnuShaper) handleData(s string, nextState OmciState) {
	trimmed := strings.TrimSpace(s)
	if len(trimmed) == 0 || shaper.matchHeader(trimmed) {
		// Unexpected header or empty line during data collection, reset
		shaper.state = CalixOnuLineInit
		shaper.omciBuff = ""
		if shaper.matchHeader(trimmed) {
			shaper.handleInit(s)
		}
		return
	}

	data := models.StringStrip(trimmed)
	shaper.omciBuff += data

	if nextState == CalixOnuLineInit {
		if len(shaper.omciBuff) == 80 {
			shaper.omciMetaDataMap[defaultCalixOnuID] = append(shaper.omciMetaDataMap[defaultCalixOnuID], OmciMetaData{
				Timestamp: shaper.timeStamp,
				RawData:   shaper.omciBuff,
			})
		}
		shaper.omciBuff = ""
	}

	shaper.state = nextState
}

func (shaper *calixOnuShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("calix onu omci shape")
	shaper.outputPath = outputPath
	shaper.state = CalixOnuLineInit
	shaper.omciBuff = ""
	shaper.onus = map[string]string{}
	shaper.omciMetaDataMap = map[string][]OmciMetaData{}
	shaper.initOnu(outputPath)

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
		shaper.processLine(string(line))
	}

	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)

	return shaper.onus
}

func init() {
	shaper := new(calixOnuShaper)
	shaperRegist(CalixOnu, shaper)
}
