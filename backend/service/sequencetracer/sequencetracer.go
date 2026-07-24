/*
# ------------------------------------------------------------
# -- sequencetracer.go
# --
# -- Huang Minghe
# -- 2024-8-2
# ------------------------------------------------------------
*/
package sequencetracer

import (
	"omciAnalyzer/service/sequencetracer/filter"
	"omciAnalyzer/service/sequencetracer/filter/message"
	"omciAnalyzer/utils"
	"sort"
)

func SequenceTracer(logs []string) string {
	// Extract logs
	var metas []message.Meta
	for _, log := range logs {
		logType, err := IdentifyLogType(log)
		if err != nil || logType == Invalid {
			utils.Log("Invalid log type: ", err)
			return ""
		}
		utils.Log("logType is:", logTypeDef[logType])
		// meta, err := ExtractLogs(log, schemas[logTypeDef[logType]])
		filter := filter.EventFilterDef[logType]
		meta, err := filter.Extract(log)
		if err != nil {
			utils.Log("Error extracting logs:", err)
			return ""
		}
		metas = append(metas, meta...)
	}
	// utils.Log("metas: ", metas)
	//sort by time
	sort.Slice(metas, func(i, j int) bool {
		return metas[i].Time.Before(metas[j].Time)
	})

	// Generate Mermaid diagram
	return GenerateMermaid(metas)
}
