/*
# ------------------------------------------------------------
# -- mermaidApi.go
# --
# -- Huang Minghe
# -- 2024-11-15
# ------------------------------------------------------------
*/
package sequencetracer

import (
	"fmt"
	"omciAnalyzer/service/sequencetracer/filter/message"
)

type RGBA struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha float64
}

var colors []RGBA = []RGBA{
	{Red: 0, Green: 255, Blue: 0, Alpha: 0.1},
	{Red: 0, Green: 0, Blue: 255, Alpha: 0.1},
	{Red: 255, Green: 0, Blue: 0, Alpha: 0.1},
	{Red: 0, Green: 255, Blue: 255, Alpha: 0.1},
	{Red: 255, Green: 255, Blue: 0, Alpha: 0.1},
	{Red: 255, Green: 0, Blue: 255, Alpha: 0.1},
	{Red: 255, Green: 128, Blue: 0, Alpha: 0.1},
}

func ColorToString(color RGBA) string {
	return fmt.Sprintf("rgba(%d, %d, %d, %.1f)", color.Red, color.Green, color.Blue, color.Alpha)
}

// GenerateMermaid generates a Mermaid diagram from the extracted logs
func GenerateMermaid(logs []message.Meta) string {
	mermaid := "sequenceDiagram\n"
	mermaid += `
            participant dmsp as DMSP
            participant alarmmgr as AlarmMgr
            participant vonumgmt as VonuMgmt
            participant hypervisor as Hypervisor
            participant onumgntolt as OnuMgntOlt
            participant xponhwa as XponHwa
            participant glob as GLOB
            participant onu as ONU Device` + "\n"
	colorIndex := 0
	for _, log := range logs {
		timeStamp := log.Time.Format("2006-01-02 15:04:05.000000")

		color := ColorToString(colors[colorIndex])
		colorIndex = (colorIndex + 1) % len(colors)
		mermaid += fmt.Sprintf("%s", "rect "+color+"\n")
		mermaid += fmt.Sprintf("%s->>%s: %s [%s]\n", log.From, log.To, log.Event, timeStamp)
		if log.Notes != "" {
			mermaid += fmt.Sprintf("Note right of %s: %s\n", log.To, log.Notes)
		}
		mermaid += "end\n"
	}
	return mermaid
}
