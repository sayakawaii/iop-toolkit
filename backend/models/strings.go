package models

import (
	"regexp"
	"strings"
)

func StringStrip(str string) string {
	line := strings.TrimSpace(str)
	line = strings.Trim(line, "\n")
	line = strings.ReplaceAll(line, " ", "")
	line = strings.ReplaceAll(line, ".", "_")
	line = strings.ReplaceAll(line, "=", "_")
	line = strings.ReplaceAll(line, "<", "_")
	line = strings.ReplaceAll(line, ">", "_")
	line = strings.ReplaceAll(line, "|", "_")
	line = strings.ReplaceAll(line, "&", "_")
	line = strings.ReplaceAll(line, "/", "_")

	return line
}

func GetSubstringBetween(rawString string, startString string, endString string) (string, bool) {
	re := regexp.MustCompile(startString + "(.*)" + endString)
	subMatch := re.FindStringSubmatch(rawString)
	if len(subMatch) < 2 {
		return "", false
	}
	return subMatch[1], true
}
