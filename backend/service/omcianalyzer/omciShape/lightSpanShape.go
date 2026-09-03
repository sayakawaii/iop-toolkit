/*
# ------------------------------------------------------------
# -- lightSpanShape.go
# --
# -- Huang Minghe
# -- 2022-4-30
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

type lightSpanShaper struct {
	cachedRegex     *regexp.Regexp
	cachedLayout    string
	outputPath      string
	state           OmciState
	onus            map[string]string
	ontid           string
	omciBuff        string
	timeStamp       time.Time
	direction       string
	sstOmciFileName string
	omciMetaDataMap map[string][]OmciMetaData
}

const (
	LightSpanLineInit OmciState = iota
	LightSpanLineData0
	LightSpanLineData1
	LightSpanLineData2
	LightSpanLineData3
	LightSpanLineInvalid
)

func (shaper *lightSpanShaper) getTimestamp(s string) time.Time {
	// example: 11:33:47.604471 Dir: Rx <-- Onu: V-ANI-1-1-1
	// example: [001bdc][2026/02/10 03:02:44.220305][omci-ONT1][TRACE] Dir: Rx <-- Onu: ONT1
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
		{`^\d{2}:\d{2}:\d{2}\.\d{6}`, "15:04:05.000000"},
		{`\[\d{4}/\d{2}/\d{2}\s+\d{2}:\d{2}:\d{2}\.\d{6}\]`, "2006/01/02 15:04:05.000000"},
		{`\d{4}-\d{2}-\d{2}\s+\d{2}:\d{2}:\d{2}`, "2006-01-02 15:04:05"},
		{`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:Z|[+-]\d{2}:\d{2})`, time.RFC3339},
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

func (shaper *lightSpanShaper) getOnuId(s string) string {
	leftIndex := strings.LastIndex(s, "Onu: ")
	rightIndex := strings.LastIndex(s, "Retry:")
	if rightIndex > 0 && rightIndex < leftIndex {
		utils.Log("invalid line:\n" + s)
		return ""
	}
	ontidstr := ""
	if rightIndex >= 0 {
		ontidstr = s[leftIndex:rightIndex]
	} else {
		ontidstr = s[leftIndex:]
	}
	ontid := ontidstr[5:]
	ontid = models.StringStrip(ontid)
	return ontid
}

func (shaper *lightSpanShaper) matchLineInit(s string) bool {
	if !strings.Contains(s, "Dir: ") ||
		!strings.Contains(s, "Onu: ") {
		return false
	}

	return strings.Contains(s, "-->") ||
		strings.Contains(s, "<--")
}

func (shaper *lightSpanShaper) parseDirection(s string) (string, bool) {
	switch {
	case strings.Contains(s, "Tx -->"):
		return "TX", true
	case strings.Contains(s, "Rx <--"):
		return "RX", true
	default:
		return "", false
	}
}

func (shaper *lightSpanShaper) processLine(s string) {
	switch shaper.state {
	case LightSpanLineInit:
		shaper.handleInit(s)
	case LightSpanLineData0:
		shaper.handleData0(s)
	case LightSpanLineData1:
		shaper.handleData1(s)
	case LightSpanLineData2:
		shaper.handleData2(s)
	}
}

func (shaper *lightSpanShaper) handleInit(s string) {
	if !shaper.matchLineInit(s) {
		return
	}
	shaper.timeStamp = shaper.getTimestamp(s)
	shaper.state = LightSpanLineData0
	shaper.ontid = shaper.getOnuId(s)
	fileName := shaper.outputPath + shaper.ontid
	shaper.sstOmciFileName = shaper.outputPath + shaper.ontid + ".sst.txt"
	if _, res := shaper.onus[shaper.ontid]; !res {
		shaper.onus[shaper.ontid] = fileName
		shaper.omciMetaDataMap[shaper.ontid] = []OmciMetaData{}
	}
	// get direction for sst
	direction, ok := shaper.parseDirection(s)
	if !ok {
		return
	}
	shaper.direction = direction
}

func (shaper *lightSpanShaper) handleData0(s string) {
	marker := "00000000  "
	if !strings.Contains(s, marker) {
		return
	}
	data, ok := extractSegment(s, marker, 48)
	if !ok {
		shaper.state = LightSpanLineInit
		return
	}

	shaper.omciBuff += data
	shaper.state = LightSpanLineData1
}

func (shaper *lightSpanShaper) handleData1(s string) {
	marker := "00000010  "
	if !strings.Contains(s, marker) {
		return
	}
	data, ok := extractSegment(s, marker, 48)
	if !ok {
		shaper.state = LightSpanLineInit
		return
	}

	shaper.omciBuff += data
	shaper.state = LightSpanLineData2
}

func (shaper *lightSpanShaper) handleData2(s string) {
	marker := "00000020  "
	if !strings.Contains(s, marker) {
		return
	}
	sstOmciBuff := shaper.omciBuff
	data, ok := extractSegment(s, marker, 24)
	if !ok {
		shaper.state = LightSpanLineInit
		return
	}

	shaper.omciBuff += data
	if len(shaper.omciBuff) == 80 {
		shaper.omciMetaDataMap[shaper.ontid] = append(shaper.omciMetaDataMap[shaper.ontid], OmciMetaData{
			Timestamp: shaper.timeStamp,
			RawData:   shaper.omciBuff,
		})
	}
	len := 0
	switch shaper.direction {
	case "TX":
		len = 36
	case "RX":
		len = 48
	}
	data, ok = extractSegment(s, marker, len)
	if !ok {
		shaper.state = LightSpanLineInit
		return
	}
	sstOmciBuff += data
	sstWriteFile(shaper.sstOmciFileName, sstOmciBuff, shaper.direction)
	shaper.direction = ""
	shaper.omciBuff = ""
	shaper.state = LightSpanLineInit
}

func (shaper *lightSpanShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("lightSpan omci shape")
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
		/*
			for lightSpan onu name max 128 bytes
			example: [2023/03/06 20:27:06.618163][omci/XS-110G-A][DEBUG] Dir: Rx <-- Onu: XS-110G-A
			above example there are 2 onu name string, so onu name will take max 256 bytes
			in another log formart, v-onu log, there is a spcific timestamp at each line start
			example: 2022-04-12T12:53:27.727622137Z [2022/04/12 12:53:27.727444][omci/device=ont1-fwlt-c][DEBUG] Dir: Rx <-- Onu: device=ont1-fwlt-c
			so, base above 2 case, we set valid log length less than 400 bytes per line
		*/
		if len(line) == 0 || len(line) >= 400 {
			continue
		}
		shaper.processLine(string(line))
	}

	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)

	return shaper.onus
}

func init() {
	shaper := new(lightSpanShaper)
	shaperRegist(LightSpan, shaper)
}
