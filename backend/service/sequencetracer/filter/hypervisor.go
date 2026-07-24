/*
# ------------------------------------------------------------
# -- hypervisor.go
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
	"time"
)

const (
	Hypervisor = 1
)

type hypervisor struct {
	EventFilterBase
}

func (ev *hypervisor) matchLog(line string, timestamp time.Time) message.Meta {
	for _, schema := range ev.schemas {
		matched := true
		for _, regexStr := range schema.Event {
			// if !strings.Contains(line, event) {
			// 	matched = false
			// 	break
			// }
			re := regexp.MustCompile(regexStr)
			matches := re.FindStringSubmatch(line)
			if len(matches) > 0 {
				fmt.Printf("Matched Event: %s\n", schema.Alias)
				fmt.Printf("Extracted Data: %+v\n", matches[1:])
			}
		}
		if matched {
			notes := schema.Notes
			return message.Meta{timestamp, schema.From, schema.To, schema.Alias, notes, line}
		}
	}
	return message.Meta{}
}

// ExtractLogs extracts relevant information from the log file based on the schemas
func (ev *hypervisor) Extract(filename string) ([]message.Meta, error) {
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
	var HypervisorSchemaFile string
	if os.Getenv("RUN_MODE") == "test" {
		HypervisorSchemaFile = "../../resource/events/schema/hypervisor.json"
	} else {
		HypervisorSchemaFile = "./resource/events/schema/hypervisor.json"
	}
	ev := new(hypervisor)
	schema, err := readJSON(HypervisorSchemaFile)
	if err != nil {
		utils.Log("Error reading HypervisorSchemaFile JSON file:", err)
		return
	}
	ev.schemas = schema
	eventFilterRegist(Hypervisor, ev)
}
