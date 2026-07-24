/*
# ------------------------------------------------------------
# -- xponHwa.go
# --
# -- Huang Minghe
# -- 2024-11-15
# ------------------------------------------------------------
*/
package filter

import (
	"bufio"
	"omciAnalyzer/service/sequencetracer/filter/message"
	"omciAnalyzer/utils"
	"os"
	"strings"
	"time"
)

const (
	XponHwa = 3
)

type onuInfo_xponHwa struct {
	vaniName string
	onuName  string
	// vaniObjectIndex    string
	// ctObjectIndex      string
	// channelObjectIndex string
}

type xponHwa struct {
	EventFilterBase
	info onuInfo_xponHwa
	msg  []string
}

func (ev *xponHwa) exactOnuInfo(line string) {
	if tmp := extractField(line, "vontani name"); tmp != "" {
		ev.info.vaniName = tmp
	}
	if tmp := extractField(line, "onuName"); tmp != "" {
		ev.info.onuName = tmp
	}
}

func (ev *xponHwa) extractJsonMsg(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	inJSONBlock := false
	var jsonBuilder strings.Builder
	bracketCount := 0
	var jsonList []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "[yproto:main],{") {
			inJSONBlock = true
			bracketCount = 1
			jsonBuilder.Reset()
			jsonBuilder.WriteString("{\n")
			continue
		}

		if inJSONBlock {
			jsonBuilder.WriteString(line)
			jsonBuilder.WriteString("\n")

			bracketCount += strings.Count(line, "{")
			bracketCount -= strings.Count(line, "}")

			if bracketCount == 0 {
				jsonList = append(jsonList, jsonBuilder.String())
				inJSONBlock = false
			}
		}
	}

	if err := scanner.Err(); err != nil {
		utils.Log("Error reading file:", err)
	}
	return jsonList, nil
}

func (ev *xponHwa) matchLog(line string, timestamp time.Time) message.Meta {
	for _, schema := range ev.schemas {
		matched := true
		matchedEvent := ""
		for _, event := range schema.Event {
			if !strings.Contains(line, event) {
				matched = false
				break
			}
			matchedEvent = event
		}
		if matched {
			// utils.Log("msg:", ev.msg[0])
			if handler, ok := message.GetHandler(matchedEvent); ok {
				jsondata := ev.msg[0]
				ev.msg = removeElement(ev.msg, 0)
				return handler.Handle(jsondata, timestamp, schema, line)
			} else {
				return message.Meta{timestamp, schema.From, schema.To, schema.Alias, schema.Notes, line}
			}
		}
	}
	return message.Meta{}
}

func (ev *xponHwa) Extract(filename string) ([]message.Meta, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ev.msg, _ = ev.extractJsonMsg(filename)
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
	var XponHwaSchemaFile string
	if os.Getenv("RUN_MODE") == "test" {
		XponHwaSchemaFile = "../../resource/events/schema/xponhwa.json"
	} else {
		XponHwaSchemaFile = "./resource/events/schema/xponhwa.json"
	}
	ev := new(xponHwa)
	schema, err := readJSON(XponHwaSchemaFile)
	if err != nil {
		utils.Log("Error reading XponHwaSchemaFile JSON file:", err)
		return
	}
	ev.schemas = schema
	eventFilterRegist(XponHwa, ev)
}
