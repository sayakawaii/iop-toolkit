package api

import (
	"github.com/iancoleman/orderedmap"
)

const (
	Success           uint8 = 0
	CmdProcError      uint8 = 1
	CmdNotSupport     uint8 = 2
	ParameterError    uint8 = 3
	UnknownME         uint8 = 4
	UnknownMEInst     uint8 = 5
	DeviceBusy        uint8 = 6
	InstExists        uint8 = 7
	Reserv2           uint8 = 8
	AttrFailOrUnknown uint8 = 9
)

func OmciContentAttrsToJSON(spValue map[string]any) *orderedmap.OrderedMap {
	result := orderedmap.New()

	for k, v := range spValue {
		switch value := v.(type) {
		case map[string]any:
			result.Set(k, OmciContentAttrsToJSON(value))
		case []any:
			if k == "Received frame VLAN tagging operation data" {
				result.Set(k, value[0])
				continue
			}
			processedArray := make([]any, 0, len(value))
			for _, item := range value {
				if nestedMap, ok := item.(map[string]any); ok {
					processedArray = append(processedArray, OmciContentAttrsToJSON(nestedMap))
				} else {
					processedArray = append(processedArray, item)
				}
			}
			result.Set(k, processedArray)
		default:
			result.Set(k, value)
		}
	}

	return result
}
