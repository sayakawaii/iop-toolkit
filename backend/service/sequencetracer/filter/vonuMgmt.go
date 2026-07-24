/*
# ------------------------------------------------------------
# -- vonuMgmt.go
# --
# -- Huang Minghe
# -- 2024-11-15
# ------------------------------------------------------------
*/
package filter

import (
	"bufio"
	"fmt"
	"omciAnalyzer/service/sequencetracer/filter/message"
	"omciAnalyzer/utils"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	VonuMgmt = 2
)

type vonuMgmt struct {
	EventFilterBase
}

func (ev *vonuMgmt) matchLog(line string, timestamp time.Time) message.Meta {
	for _, schema := range ev.schemas {
		matched := true
		for _, event := range schema.Event {
			if !strings.Contains(line, event) {
				matched = false
				break
			}
		}
		if matched {
			dynamicDataRe := regexp.MustCompile(`BEGIN OnOnuRequest: ONU \[ (.*?) \] operation \[ (.*?) \]`)
			dynamicDataMatch := dynamicDataRe.FindStringSubmatch(line)
			dynamicData := ""
			if len(dynamicDataMatch) > 2 {
				dynamicData = fmt.Sprintf("ONU: %s Operation: %s", dynamicDataMatch[1], dynamicDataMatch[2])
			}
			notes := schema.Notes
			if dynamicData != "" {
				notes = dynamicData
			}
			return message.Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
		}
	}
	return message.Meta{}
}

func (ev *vonuMgmt) Extract(filename string) ([]message.Meta, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var extractedLogs []message.Meta
	var lastTimestamp time.Time //cache last time stamp
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lastTimestamp = extractTimestamp(line, lastTimestamp)
		if logEntry := ev.matchLog(line, lastTimestamp); !logEntry.IsEmpty() {
			extractedLogs = append(extractedLogs, logEntry)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return extractedLogs, nil
}

func init() {
	var VonuMgmtSchemaFile string
	if os.Getenv("RUN_MODE") == "test" {
		VonuMgmtSchemaFile = "../../resource/events/schema/vonumgmt.json"
	} else {
		VonuMgmtSchemaFile = "./resource/events/schema/vonumgmt.json"
	}
	ev := new(vonuMgmt)
	schema, err := readJSON(VonuMgmtSchemaFile)
	if err != nil {
		utils.Log("Error reading VonuMgmtSchemaFile JSON file:", err)
		return
	}
	ev.schemas = schema
	eventFilterRegist(VonuMgmt, ev)
}
