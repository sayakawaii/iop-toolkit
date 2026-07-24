/*
# ------------------------------------------------------------
# -- copperOnuShape.go
# --
# -- Huang Minghe
# -- 2023-9-12
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

type copperOnuShaper struct {
	cachedRegex     *regexp.Regexp
	cachedLayout    string
	outputPath      string
	state           OmciState
	onus            map[string]string
	omciBuff        string
	timeStamp       time.Time
	omciMetaDataMap map[string][]OmciMetaData
}

const (
	CopperOnuLineInit OmciState = iota
	CopperOnuLineData1
	CopperOnuLineData2
	CopperOnuLineData3
	CopperOnuLineData4
	CopperOnuLineInvalid
)

const defaultCopperOnuID = "CopperOnu"

func (shaper *copperOnuShaper) getTimestamp(s string) time.Time {
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

	patterns := []struct {
		regex  string
		layout string
	}{
		{`\[\d{2}/\d{2}/\d{4}-\d{2}:\d{2}:\d{2}\.\d{6}\]`, "02/01/2006-15:04:05.000000"},
		{`\d{2}/\d{2}/\d{4}-\d{2}:\d{2}:\d{2}\.\d{6}`, "02/01/2006-15:04:05.000000"},
		{`\d{2}/\d{2}/\d{2}\s+\d{2}:\d{2}:\d{2}\.\d{6}`, "02/01/06 15:04:05.000000"},
		{`\[\d{4}/\d{2}/\d{2}\s+\d{2}:\d{2}:\d{2}\.\d{6}\]`, "2006/01/02 15:04:05.000000"},
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

func (shaper *copperOnuShaper) initOnu(outputPath string) {
	fileName := outputPath + defaultCopperOnuID
	// Keep output deterministic across repeated shape runs.
	if err := os.WriteFile(fileName, []byte{}, 0666); err != nil {
		utils.Log("truncate file failed: " + err.Error())
	}
	if _, exists := shaper.onus[defaultCopperOnuID]; !exists {
		shaper.onus[defaultCopperOnuID] = fileName
		shaper.omciMetaDataMap[defaultCopperOnuID] = []OmciMetaData{}
	}
}

func (shaper *copperOnuShaper) processLine(s string) {
	if shaper.state != CopperOnuLineInit && strings.Contains(s, "bytes[06-13]") {
		shaper.state = CopperOnuLineInit
		shaper.omciBuff = ""
		return
	}

	switch shaper.state {
	case CopperOnuLineInit:
		shaper.handleInit(s)
	case CopperOnuLineData1:
		shaper.handleData1(s)
	case CopperOnuLineData2:
		shaper.handleData2(s)
	case CopperOnuLineData3:
		shaper.handleData3(s)
	case CopperOnuLineData4:
		shaper.handleData4(s)
	}
}

func (shaper *copperOnuShaper) handleInit(s string) {
	if !strings.Contains(s, "bytes[06-13]") {
		return
	}

	shaper.timeStamp = shaper.getTimestamp(s)
	data, ok := extractSegment(s, "bytes[06-13]:", 30)
	if !ok {
		shaper.state = CopperOnuLineInit
		shaper.omciBuff = ""
		return
	}

	shaper.omciBuff = data
	shaper.state = CopperOnuLineData1
}

func (shaper *copperOnuShaper) handleData1(s string) {
	if !strings.Contains(s, "bytes[14-45]") {
		return
	}

	data, ok := extractSegment(s, "bytes[14-45]:", 42)
	if !ok {
		shaper.state = CopperOnuLineInit
		shaper.omciBuff = ""
		return
	}

	shaper.omciBuff += data
	shaper.state = CopperOnuLineData2
}

func (shaper *copperOnuShaper) handleData2(s string) {
	if !strings.Contains(s, ": ") {
		return
	}

	data, ok := extractSegment(s, ": ", 42)
	if !ok {
		shaper.state = CopperOnuLineInit
		shaper.omciBuff = ""
		return
	}

	shaper.omciBuff += data
	shaper.state = CopperOnuLineData3
}

func (shaper *copperOnuShaper) handleData3(s string) {
	if !strings.Contains(s, ": ") {
		return
	}

	data, ok := extractSegment(s, ": ", 38)
	if !ok {
		shaper.state = CopperOnuLineInit
		shaper.omciBuff = ""
		return
	}

	shaper.omciBuff += data
	shaper.state = CopperOnuLineData4
	shaper.handleData4("")
}

func (shaper *copperOnuShaper) handleData4(_ string) {
	if len(shaper.omciBuff) == 80 {
		shaper.omciMetaDataMap[defaultCopperOnuID] = append(shaper.omciMetaDataMap[defaultCopperOnuID], OmciMetaData{
			Timestamp: shaper.timeStamp,
			RawData:   shaper.omciBuff,
		})
	}

	shaper.omciBuff = ""
	shaper.state = CopperOnuLineInit
}

func (shaper *copperOnuShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("copperOnu omci shape")
	shaper.outputPath = outputPath
	shaper.state = CopperOnuLineInit
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
		if len(line) >= 400 {
			continue
		}

		shaper.processLine(string(line))
	}

	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)

	return shaper.onus
}

func init() {
	shaper := new(copperOnuShaper)
	shaperRegist(CopperOnu, shaper)
}
