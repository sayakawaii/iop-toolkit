/*
# ------------------------------------------------------------
# -- iidentifyLogType.go
# --
# -- Huang Minghe
# -- 2024-11-15
# ------------------------------------------------------------
*/
package sequencetracer

import (
	"bufio"
	"os"
	"regexp"
)

const (
	Hypervisor = 1
	VonuMgmt   = 2
	XponHwa    = 3
	OnuMgntOlt = 4
	Glob       = 5
	Invalid    = 6
)

var logTypeDef map[uint8]string = map[uint8]string{
	1: "hypervisor",
	2: "vonumgmt",
	3: "xponhwa",
	4: "onumgntolt",
	5: "glob",
	6: "Invalid",
}

// log type regexp
var logTypeRules = []struct {
	regex   *regexp.Regexp
	logType uint8
}{
	{regexp.MustCompile(`(?i)\[hypervisor\]`), Hypervisor},
	{regexp.MustCompile(`(?i)\[omci-default\]`), VonuMgmt},
	{regexp.MustCompile(`(?i)\[xponhwa(?::main)?\]`), XponHwa},
	{regexp.MustCompile(`(?i)\[onumgntolt(?:_app)?\]`), OnuMgntOlt}, // match [OnuMgntOlt] or [OnuMgntOlt_app]
	{regexp.MustCompile(`(?i)\[glob\]`), Glob},
}

func IdentifyLogType(filePath string) (uint8, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Invalid, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logLine := scanner.Text()
		for _, rule := range logTypeRules {
			if rule.regex.MatchString(logLine) {
				return rule.logType, nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return Invalid, err
	}

	return Invalid, nil
}
