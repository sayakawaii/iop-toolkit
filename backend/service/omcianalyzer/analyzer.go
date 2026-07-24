/*
# ------------------------------------------------------------
# -- analyzer.go
# --
# -- Huang Minghe
# -- 2022-7-8
# ------------------------------------------------------------
*/
package omcianalyzer

import (
	"encoding/json"

	"omciAnalyzer/service/omcianalyzer/omciDiagram"
	"omciAnalyzer/service/omcianalyzer/omciSchema"
	"omciAnalyzer/service/omcianalyzer/omciShape"
	"omciAnalyzer/service/plantUml"
	"omciAnalyzer/utils"
)

type DiagramData struct {
	Category string `json:"category"`
	Path     string `json:"path"`
}

type PlantumlData struct {
	Category string `json:"category"`
	Path     string `json:"path"`
}

// UnmarshalJSON handles backward compatibility for PlantumlData.
// Old format stored a plain string (file path), new format is {"category":"...","path":"..."}.
func (p *PlantumlData) UnmarshalJSON(data []byte) error {
	// Try old format: plain string
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		p.Path = s
		p.Category = ""
		return nil
	}
	// New format: object
	type plantumlDataAlias PlantumlData
	var alias plantumlDataAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	*p = PlantumlData(alias)
	return nil
}

type AnalyzerContent struct {
	SwVersion string                  `json:"swVersion"`
	HwVersion string                  `json:"hwVersion"`
	JsonName  string                  `json:"jsonName"`
	JsonPath  string                  `json:"jsonPath"`
	Plantuml  map[string]PlantumlData `json:"plantuml"`
	Diagram   map[string]DiagramData  `json:"diagram"`
	Charts    []omciDiagram.Diagram   `json:"charts"`
}

type LogContent struct {
	LogType    string                     `json:"logType"`
	OnuContent map[string]AnalyzerContent `json:"onuContent"`
}

type RespStatusContent struct {
	Progress uint8  `json:"progress"`
	Status   string `json:"status"`
	ErrorMsg string `json:"errorMsg"`
}

func Analyzer(logPath string, outputPath string, c chan LogContent, sc chan RespStatusContent) {
	defer close(c)
	defer close(sc)
	utils.Log("Analyzer start")
	//step1 shape log
	logType := omciShape.LogTypeAutoIdentify(logPath)
	if logType == omciShape.Invalid {
		utils.Log("invalid log file")
		sc <- RespStatusContent{100, "error", "invalid log file"}
		return
	}
	/*
		improve for 1k onu log
		big log file will take long time at identify and shape step
		so immediately set progress as 5 while analyzer started
		this will trig create a log in redis DB, when AJAX query come from web, can get a normal response
	*/
	sc <- RespStatusContent{10, "processing", ""}
	content := new(LogContent)
	content.OnuContent = make(map[string]AnalyzerContent)
	content.LogType = omciShape.LogTypeNameGet(logType)
	shaper := omciShape.ShaperDef[logType]
	onus := shaper.Shape(logPath, outputPath)
	sc <- RespStatusContent{20, "processing", ""}
	utils.Log("OMCI Shape successed")
	for k, v := range onus {
		//step2 convert log to json
		ret := omciSchema.OmciRawDataParse(v, outputPath+k+".json", logType)
		if ret == "" {
			continue
		}
		//step3 convert omci to diagram

		analyzerContent := new(AnalyzerContent)
		analyzerContent.JsonName = k + ".json"
		analyzerContent.JsonPath = outputPath + k + ".json"
		analyzerContent.Plantuml = make(map[string]PlantumlData)
		analyzerContent.Diagram = make(map[string]DiagramData)

		diagramSet, ontInfo := omciDiagram.DiagramProc(ret)
		analyzerContent.Charts = diagramSet
		analyzerContent.SwVersion = ontInfo.SwVersion
		analyzerContent.HwVersion = ontInfo.HwVersion
		for _, v := range diagramSet {
			fileName := outputPath + k + "-" + v.Category + "-" + v.Name
			omciDiagram.PlantUmlGenerate(fileName+".wsd", v.Chart)
			p := PlantumlData{v.Category, fileName + ".wsd"}
			analyzerContent.Plantuml[v.Name] = p
			d := DiagramData{v.Category, fileName + ".svg"}
			analyzerContent.Diagram[v.Name] = d
		}
		content.OnuContent[k] = *analyzerContent
	}
	sc <- RespStatusContent{50, "processing", ""}
	plantUml.Draw(outputPath+"*.wsd", outputPath)
	utils.Log("OMCI Diagram successed")
	//step4 save DB
	sc <- RespStatusContent{100, "success", ""}
	c <- *content
	//step5 return to web
}
