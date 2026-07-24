/*
# ------------------------------------------------------------
# -- isamShape.go
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

type isamShaper struct {
	cachedRegex     *regexp.Regexp
	cachedLayout    string
	outputPath      string
	state           OmciState
	onus            map[string]string
	ontid           string
	omciBuff        string
	timeStamp       time.Time
	omciMetaDataMap map[string][]OmciMetaData
}

const (
	IsamLineInit OmciState = iota
	IsamLineData0
	IsamLineData1
	IsamLineData2
	IsamLineData3
	IsamLineData4
	IsamLineInvalid
)

func (shaper *isamShaper) getTimestamp(s string) time.Time {
	// example: [OMFT(omft):100 27/10/20 01:25:42.059012 ]<OMCI MSG> Rx <-- OntId : 222 -  637:35:50.849
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
		{`\d{2}/\d{2}/\d{2}\s+\d{2}:\d{2}:\d{2}\.\d{6}`, "02/01/06 15:04:05.000000"},
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

func (shaper *isamShaper) getOnuId(s string) string {
	gponOnuId := "OntId : "
	ng2OnuId := "OntId(NG2) : "
	prefix := ""
	if strings.Contains(s, gponOnuId) {
		prefix = gponOnuId
	} else if strings.Contains(s, ng2OnuId) {
		prefix = ng2OnuId
	} else {
		return ""
	}

	leftIndex := strings.Index(s, prefix)
	rightIndex := strings.Index(s, " - ")
	if leftIndex == -1 || rightIndex == -1 || rightIndex <= leftIndex {
		return ""
	}

	ontid := string(s[leftIndex+len(prefix) : rightIndex])
	ontid = models.StringStrip(ontid)

	fileName := shaper.outputPath + ontid
	if _, exists := shaper.onus[ontid]; !exists {
		shaper.onus[ontid] = fileName
	}

	return ontid
}

func (shaper *isamShaper) matchLineInit(s string) bool {
	if !strings.Contains(s, "<OMCI MSG>") ||
		!strings.Contains(s, "OntId : ") && !strings.Contains(s, "OntId(NG2) : ") {
		return false
	}

	return strings.Contains(s, "-->") ||
		strings.Contains(s, "<--")
}

func (shaper *isamShaper) processLine(s string) {
	switch shaper.state {
	case IsamLineInit:
		shaper.handleInit(s)
	case IsamLineData0:
		shaper.handleData0(s)
	case IsamLineData1:
		shaper.handleData(s, IsamLineData2)
	case IsamLineData2:
		shaper.handleData(s, IsamLineData3)
	case IsamLineData3:
		shaper.handleData(s, IsamLineData4)
	case IsamLineData4:
		shaper.handleData(s, IsamLineInit)
	}
}

func (shaper *isamShaper) handleInit(s string) {
	if !shaper.matchLineInit(s) {
		return
	}
	shaper.timeStamp = shaper.getTimestamp(s)
	shaper.state = IsamLineData0
	shaper.ontid = shaper.getOnuId(s)
	fileName := shaper.outputPath + shaper.ontid
	if _, res := shaper.onus[shaper.ontid]; !res {
		shaper.onus[shaper.ontid] = fileName
		shaper.omciMetaDataMap[shaper.ontid] = []OmciMetaData{}
	}
}

// hexDataLineRegex matches lines starting with hex byte pairs like "04 a8 49 ..."
var hexDataLineRegex = regexp.MustCompile(`^[0-9a-fA-F]{2} [0-9a-fA-F]{2} `)

func (shaper *isamShaper) handleData0(s string) {
	if strings.Contains(s, "]") {
		data, ok := extractSegment(s, "]", 23)
		if !ok {
			shaper.state = IsamLineInit
			return
		}
		shaper.omciBuff += data
		shaper.state = IsamLineData1
		return
	}

	// Handle bare hex data line without ']' prefix (e.g. 5402z ISAM variant)
	trimmed := strings.TrimSpace(s)
	if len(trimmed) >= 23 && hexDataLineRegex.MatchString(trimmed) {
		shaper.omciBuff += models.StringStrip(trimmed[0:23])
		shaper.state = IsamLineData1
	}
}

func (shaper *isamShaper) handleData(s string, nextState OmciState) {
	if len(s) < 23 {
		shaper.state = IsamLineInit
		return
	}

	var data string
	var ok bool

	switch {
	case len(s) <= 40:
		data = models.StringStrip(s[0:23])
	case len(s) > 40:
		if strings.Contains(s, "}") {
			data, ok = extractSegment(s, "}", 24)
		} else if strings.Contains(s, "]") {
			data, ok = extractSegment(s, "]", 24)
		}
		if !ok {
			shaper.state = IsamLineInit
			return
		}
	}
	if data == "" {
		shaper.state = IsamLineInit
		return
	}

	shaper.omciBuff += data

	if nextState == IsamLineInit {
		if len(shaper.omciBuff) == 80 {
			shaper.omciMetaDataMap[shaper.ontid] = append(shaper.omciMetaDataMap[shaper.ontid], OmciMetaData{
				Timestamp: shaper.timeStamp,
				RawData:   shaper.omciBuff,
			})
		}
		shaper.omciBuff = ""
	}

	shaper.state = nextState
}

func (shaper *isamShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("isam omci shape")
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
		shaper.processLine(string(line))
	}

	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)

	return shaper.onus
}

func init() {
	shaper := new(isamShaper)
	shaperRegist(Isam, shaper)
}
