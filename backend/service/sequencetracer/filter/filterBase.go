/*
# ------------------------------------------------------------
# -- filterBase.go
# --
# -- Huang Minghe
# -- 2024-11-15
# ------------------------------------------------------------
*/
package filter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"omciAnalyzer/service/sequencetracer/filter/message"
	"omciAnalyzer/utils"
	"os"
	"regexp"
	"strings"
	"time"
)

type EventFilter interface {
	Extract(filename string) ([]message.Meta, error)
}

type EventFilterBase struct {
	schemas []message.Schema
	metas   []message.Meta
}

// ReadJSON reads a JSON file and unmarshals it into a slice of Schemas
func readJSON(filename string) ([]message.Schema, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var schemas []message.Schema
	if err := json.Unmarshal(bytes, &schemas); err != nil {
		return nil, err
	}

	return schemas, nil
}

var formats = []string{
	"02/01/2006-15:04:05.000000",
	"2006/01/02 15:04:05.000000",
}

func extractTimestamp(line string, lastTimestamp time.Time) time.Time {
	timeRe := regexp.MustCompile(`^\[([^\]]+)\]`)
	timeMatch := timeRe.FindStringSubmatch(line)
	if len(timeMatch) > 1 {
		for _, tmpl := range formats {
			// utils.Log(timeMatch[1])
			if t, err := time.Parse(tmpl, timeMatch[1]); err == nil {
				return t
			}
		}
		// utils.Log("time formart error")
		return lastTimestamp
	}
	return lastTimestamp
}

func extractField(line, fieldName string) string {
	fieldRe := regexp.MustCompile(fmt.Sprintf(`\b%s\s*=\s*([^,]+)`, fieldName))
	fieldMatch := fieldRe.FindStringSubmatch(line)
	if len(fieldMatch) > 1 {
		return strings.TrimSpace(fieldMatch[1])
	}
	return ""
}

func removeElement(slice []string, i int) []string {
	if i < 0 || i >= len(slice) {
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

func (ev *EventFilterBase) extractJsonMsg(filename string) ([]string, error) {
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

		if strings.Contains(line, "[yproto],{") {
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

func (ev *EventFilterBase) matchLog(line string, timestamp time.Time) message.Meta {
	for _, schema := range ev.schemas {
		matched := true
		for _, event := range schema.Event {
			if !strings.Contains(line, event) {
				matched = false
				break
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
func (ev *EventFilterBase) Extract(filename string) ([]message.Meta, error) {
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

var (
	EventFilterDef = make(map[uint8]EventFilter)
)

func eventFilterRegist(key uint8, filter EventFilter) {
	EventFilterDef[key] = filter
}
