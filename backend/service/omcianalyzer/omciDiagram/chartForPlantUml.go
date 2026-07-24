/*
# ------------------------------------------------------------
# -- chartForPlantUml.go
# --
# -- Huang Minghe
# -- 2022-6-28
# ------------------------------------------------------------
*/
package omciDiagram

import (
	"omciAnalyzer/global"
	"omciAnalyzer/service/omcianalyzer/omciDiagram/api"
	"omciAnalyzer/utils"
	"os"
	"path"
	"strconv"
	"strings"
)

var planUmlHead string = `@startuml
' Split into 4 pages
' page 2x2
skinparam pageMargin 10
skinparam pageExternalColor gray
skinparam pageBorderColor black

scale 1

header Transport2 eONU
' footer Page %page% of %lastpage%
footer Nokia Shanghai Bell
title Auto Generated OMCI Model

`
var plantUmlTail string = `

@enduml
`

func plantUmlNoteLinkCreate(pathStr string, note string) string {

	ext := path.Ext(pathStr)
	name := path.Base(pathStr)
	fileName := pathStr
	fileName = strings.ReplaceAll(fileName, name, "MOP_DACL")
	fileName = strings.ReplaceAll(fileName, ext, "txt")
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		utils.Log("open file failed")
	} else {
		defer file.Close()
		file.WriteString(note)
	}
	link := "http://" + global.AppConf.Server.Addr + "/" + fileName

	return link
}

func plantUmlObjectCreate(fileName string, chart api.MeChart) string {
	node := chart.Node
	var note string
	lines := "object " + "\"" + node.MeName + "\"" + " as " + node.MeUID + "{\n"
	lines = lines + "0x" + strconv.FormatUint(node.MeInst, 16) + "\n"
	if chart.Attrs != nil {
		for k, v := range chart.Attrs {
			lines = lines + k + ": " + v.(string) + "\n"
		}
	}
	lines = lines + "}\n"
	if len(chart.Note) > 0 {
		for _, v := range chart.Note {
			note = note + v + "\n"
		}
		//add notes
		if chart.Node.MeName == "MOP" {
			link := plantUmlNoteLinkCreate(fileName, note)
			lines = lines + "note left of " + node.MeUID + "\n"
			lines = lines + "[[" + link + " DACL]]\n"
			lines = lines + "end note\n"
		} else {
			lines = lines + "note right of " + node.MeUID + "\n"
			lines = lines + note
			lines = lines + "end note\n"
		}
	}
	return lines
}

func plantUmlRelationshipCreate(chart api.MeChart) string {
	node := chart.Node
	var edge api.MeEdges
	var lines, direction, note string
	for _, v := range chart.Edge {
		edge = v
		if edge.EdgeDirection != "auto" {
			direction = edge.EdgeDirection
		} else {
			direction = ""
		}
		if edge.EdgeNote != nil {
			for index := range edge.EdgeNote {
				note = edge.EdgeNote[index]
			}
			if edge.EdgeType == 1 {
				lines = lines + node.MeUID + " \"" + note + "\" -" + direction + "-> " + edge.EdgeLinkObject + "\n"
			} else {
				lines = lines + node.MeUID + " \"" + note + "\" ." + direction + ".> " + edge.EdgeLinkObject + "\n"
			}
		} else {
			if edge.EdgeType == 1 {
				lines = lines + node.MeUID + " -" + direction + "-> " + edge.EdgeLinkObject + "\n"
			} else {
				lines = lines + node.MeUID + " ." + direction + ".> " + edge.EdgeLinkObject + "\n"
			}
		}
	}
	return lines
}

func PlantUmlGenerate(fileName string, chart map[string]api.MeChart) {
	var lines string
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		utils.Log("open file failed")
	} else {
		defer file.Close()
		file.WriteString(planUmlHead)
		for _, v := range chart {
			lines = lines + plantUmlObjectCreate(fileName, v)
		}
		for _, v := range chart {
			lines = lines + plantUmlRelationshipCreate(v)
		}
		file.WriteString(lines)
		file.WriteString(plantUmlTail)
	}
}
