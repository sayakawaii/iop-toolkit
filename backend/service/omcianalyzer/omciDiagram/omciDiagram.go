package omciDiagram

import (
	"encoding/json"
	_ "omciAnalyzer/service/omcianalyzer/omciDiagram/ALCL"
	"omciAnalyzer/service/omcianalyzer/omciDiagram/api"
	"omciAnalyzer/service/omcianalyzer/omciDiagram/def"
	"omciAnalyzer/service/omcianalyzer/omciSchema"
	"omciAnalyzer/utils"
	"os"
	"strconv"
	"strings"
)

type edgeResolveRule struct {
	sourceMeName string
	edgeKey      string
}

func serviceProc(omciLine omciSchema.OmciContext, relship api.MeRelationship, chart map[string]api.MeChart, pqMap map[uint64]uint64) {
	contents := omciSchema.OmciArContents{}
	header := omciLine.Header
	err := omciSchema.GetContentFromInterface(omciLine.Contents, &contents)
	if err != nil {
		return
	}
	meId := header.MeClass
	MeInst := header.MeInst
	action := header.MsgTypeName
	handler := api.GetMeChartHandler(meId)
	MeName := relship.ShortName
	attrsList := omciSchema.SliceMapConvToMap(contents.Attrs)
	node := new(api.MeNode)
	node.MeNodeSet(MeName, MeInst)
	key := node.MeUID
	//if action is delete, delete object from map
	if action == "Delete" {
		delete(chart, key)
		return
	}
	//TODO: mark ME object set before create if need create by OLT
	if v, err := chart[key]; err {
		chart[key] = handler.GetMeChart(&v, *node, attrsList)
	} else {
		chart[key] = handler.GetMeChart(nil, *node, attrsList)
	}
	//flexible PQ workaround
	//TODO:improve flexible PQ handle
	if meId == def.PQ {
		if v, err := attrsList["Traffic scheduler pointer"]; err {
			trafficPointer := (uint64)(v.(float64))
			if trafficPointer == 0 {
				return
			}
			//create TS
			meId := (uint64)(def.TS)
			MeInst := trafficPointer
			handler := api.GetMeChartHandler(meId)
			relship := api.GetMeRelShip(meId).GetRelShip()
			MeName := relship.ShortName
			attrsList := map[string]interface{}{"T-CONT pointer": (float64)(trafficPointer)}
			node := new(api.MeNode)
			node.MeNodeSet(MeName, MeInst)
			key := node.MeUID
			if v, err := chart[key]; err {
				chart[key] = handler.GetMeChart(&v, *node, attrsList)
			} else {
				chart[key] = handler.GetMeChart(nil, *node, attrsList)
			}
			//create TCONT
			meId = (uint64)(def.TCONT)
			handler = api.GetMeChartHandler(meId)
			relship = api.GetMeRelShip(meId).GetRelShip()
			MeName = relship.ShortName
			node.MeNodeSet(MeName, MeInst)
			key = node.MeUID
			if v, err := chart[key]; err {
				chart[key] = handler.GetMeChart(&v, *node, nil)
			} else {
				chart[key] = handler.GetMeChart(nil, *node, nil)
			}
		}
	}
}

func mibuploadProc(omciLine omciSchema.OmciContext, pqMap map[uint64]uint64) {
	contents := omciSchema.OmciMibuploadNextAkContents{}
	err := omciSchema.GetContentFromInterface(omciLine.Contents, &contents)
	if err != nil {
		return
	}
	if contents.UploadMeName != "PriorityQueue" {
		return
	}
	MeInst := contents.UploadMeInstance
	attrsList := omciSchema.SliceMapConvToMap(contents.Attrs)
	if v, err := attrsList["Related port"]; err {
		tcontPointer := ((uint64)(v.(float64)) & 0x00000000FFFF0000) >> 16
		if tcontPointer != 0 {
			pqMap[MeInst] = tcontPointer
		}
	}
}

func softwareProc(omciLine omciSchema.OmciContext, ont *OntInfo) {
	contents := omciSchema.OmciGetAkContents{}
	err := omciSchema.GetContentFromInterface(omciLine.Contents, &contents)
	if err != nil {
		return
	}
	attrsList := omciSchema.SliceMapConvToMap(contents.Attrs)
	if v, err := attrsList["Is active"]; err {
		if (int)(v.(float64)) == 1 {
			if v, err := attrsList["Version"]; err {
				ont.SwVersion = v.(string)
			}
		}
	}
}

func hardwareProc(omciLine omciSchema.OmciContext, ont *OntInfo) {
	contents := omciSchema.OmciGetAkContents{}
	err := omciSchema.GetContentFromInterface(omciLine.Contents, &contents)
	if err != nil {
		return
	}
	attrsList := omciSchema.SliceMapConvToMap(contents.Attrs)
	if v, err := attrsList["Version"]; err {
		ont.HwVersion = v.(string)
	}
}

func omciJsonProc(omciLine omciSchema.OmciContext, chart map[string]api.MeChart, pqMap map[uint64]uint64) {

	meId := omciLine.Header.MeClass
	meAction := omciLine.Header.MsgTypeName
	meClass := omciLine.Header.MeClass
	meInst := omciLine.Header.MeInst
	if meAction == "MIBUploadNext" && omciLine.Header.AK == 1 {
		// handle mibupload
		mibuploadProc(omciLine, pqMap)
	} else if meAction == "Get" {
		if meClass == 7 && omciLine.Header.AK == 1 && (meInst&0x00FF == 0 || meInst&0x00FF == 1) {
			softwareProc(omciLine, &ontInfo)
		} else if meClass == 256 && omciLine.Header.AK == 1 {
			hardwareProc(omciLine, &ontInfo)
		}
	} else if meAction == "Alarm" {
		// handle alarm
		// alarmProc(omciLine, pqMap)
	} else if v := api.GetMeRelShip(meId); v != nil && api.GetMeChartHandler(meId) != nil && omciLine.Header.AR == 1 {
		// handle service
		serviceProc(omciLine, v.GetRelShip(), chart, pqMap)
	}
}

func isNodeReferenced(chart map[string]api.MeChart, node api.MeNode) bool {
	//chart
	for _, c := range chart {
		//edge
		for _, e := range c.Edge {
			if e.EdgeLinkObject == node.MeUID {
				return true
			}
		}
	}
	return false
}

func pqToTcontMapProc(chart map[string]api.MeChart, pqMap map[uint64]uint64) {
	var meId, MeInst, tpPointer uint64
	node := new(api.MeNode)
	var relship api.MeRelationship
	var MeName, tpMeName string
	tmp := new(api.MeChart)
	edgeTP := new(api.MeEdges)
	for k, v := range pqMap {
		//skip downstream queue
		if k < 32768 || v < 32768 {
			continue
		}
		//PQ
		meId = def.PQ
		MeInst = k
		tpPointer = v
		relship = api.GetMeRelShip(meId).GetRelShip()
		MeName = relship.ShortName
		node.MeNodeSet(MeName, MeInst)
		key := node.MeUID
		//don't display non-referenced PQ node
		if !isNodeReferenced(chart, *node) {
			delete(chart, key)
			continue
		}
		if v, err := chart[key]; err && len(v.Edge) > 0 {
			continue
		} else {
			tmp.Node = *node
			tmp.Edge = make(map[string]api.MeEdges)
			tmp.Note = make(map[string]string)
			tpMeName = api.GetMeRelShip(relship.AssociationType[1]).GetRelShip().ShortName
			eType := relship.RelationshipType
			eNote := []string(nil)
			eDir := "up"
			edgeTP.MeEdgesSet(eType, eDir, tpMeName, tpPointer, eNote)
			tmp.Edge["Related Port"] = *edgeTP
			tmp.Note = nil
			chart[key] = *tmp
		}
		//TCONT
		meId = def.TCONT
		MeInst = v
		relship = api.GetMeRelShip(meId).GetRelShip()
		MeName = relship.ShortName
		node.MeNodeSet(MeName, MeInst)
		key = node.MeUID
		if _, err := chart[key]; err {
			continue
		} else {
			chart[key] = api.NodeWithoutEdge(nil, *node, nil)
		}
	}
}

func tcontChartProc(pqMap map[uint64]uint64) []Diagram {
	var meId, MeInst, tpPointer uint64
	node := new(api.MeNode)
	var relship api.MeRelationship
	var MeName, tpMeName string
	tmp := new(api.MeChart)
	edgeTP := new(api.MeEdges)
	var diagramSet []Diagram
	diagram := new(Diagram)
	diagram.Category = Tcont
	tcontPQList := make(map[uint64][]uint64)
	//because GO map no senquence, must load all data first
	for k, v := range pqMap {
		tcontPQList[v] = append(tcontPQList[v], k)
	}
	for k, v := range tcontPQList {
		diagram.Chart = make(map[string]api.MeChart)
		//TCONT
		meId = def.TCONT
		MeInst = k
		relship = api.GetMeRelShip(meId).GetRelShip()
		MeName = relship.ShortName
		node.MeNodeSet(MeName, MeInst)
		key := node.MeUID
		if _, err := diagram.Chart[key]; err {
			continue
		} else {
			diagram.Name = strconv.FormatUint(MeInst, 16)
			diagram.Chart[key] = api.NodeWithoutEdge(nil, *node, nil)
		}
		//PQ
		meId = def.PQ
		tpPointer = k
		for _, v := range v {
			MeInst = v
			relship = api.GetMeRelShip(meId).GetRelShip()
			MeName = relship.ShortName
			node.MeNodeSet(MeName, MeInst)
			key := node.MeUID
			if v, err := diagram.Chart[key]; err && len(v.Edge) > 0 {
				continue
			} else {
				tmp.Node = *node
				tmp.Edge = make(map[string]api.MeEdges)
				tmp.Note = make(map[string]string)
				tpMeName = api.GetMeRelShip(relship.AssociationType[1]).GetRelShip().ShortName
				eType := relship.RelationshipType
				eNote := []string(nil)
				eDir := "up"
				edgeTP.MeEdgesSet(eType, eDir, tpMeName, tpPointer, eNote)
				tmp.Edge["Related Port"] = *edgeTP
				tmp.Note = nil
				diagram.Chart[key] = *tmp
			}
		}
		//save one TCONT&PQ data
		diagramSet = append(diagramSet, *diagram)
	}
	return diagramSet
}

func resolveHostTpTypeEdges(chart map[string]api.MeChart) {
	rules := []edgeResolveRule{
		{sourceMeName: "MBPCD", edgeKey: "TP pointer"},
		{sourceMeName: "VTOCD", edgeKey: "Associated ME pointer"},
		{sourceMeName: "EVTOCD", edgeKey: "Associated me pointer"},
		{sourceMeName: "TcpUdp", edgeKey: "IP host pointer"},
	}

	for key, meChart := range chart {
		for _, rule := range rules {
			if meChart.Node.MeName != rule.sourceMeName {
				continue
			}

			edge, ok := meChart.Edge[rule.edgeKey]
			if !ok || !strings.HasPrefix(edge.EdgeLinkObject, "IPHost_") {
				continue
			}

			resolved := resolveHostTpTypeEdgeLink(chart, edge.EdgeLinkObject)
			if resolved == edge.EdgeLinkObject {
				continue
			}

			edge.EdgeLinkObject = resolved
			meChart.Edge[rule.edgeKey] = edge
			chart[key] = meChart
		}
	}
}

func resolveHostTpTypeEdgeLink(chart map[string]api.MeChart, linkObject string) string {
	_, instHex, found := strings.Cut(linkObject, "_")
	if !found {
		return linkObject
	}

	if _, err := strconv.ParseUint(instHex, 16, 64); err != nil {
		return linkObject
	}

	ipv6LinkObject := "IPV6Host_" + instHex
	if _, ok := chart[ipv6LinkObject]; ok {
		return ipv6LinkObject
	}

	ipv4LinkObject := "IPHost_" + instHex
	if _, ok := chart[ipv4LinkObject]; ok {
		return ipv4LinkObject
	}

	return linkObject
}

const (
	Latest   = "latest"
	Tcont    = "tcont"
	Update   = "update"
	Invalide = "unknow"
)

type Diagram struct {
	Name     string
	Category string //latest, tcont, update
	Chart    map[string]api.MeChart
}

type OntInfo struct {
	SwVersion string
	HwVersion string
}

var ontInfo OntInfo

func DiagramProc(jsonPath string) ([]Diagram, OntInfo) {
	utils.Log("omci diagram process")
	ontInfo.SwVersion = ""
	ontInfo.SwVersion = ""
	file, err := os.Open(jsonPath)
	if err != nil {
		utils.Log("open file failed")
		return nil, ontInfo
	}
	defer file.Close()
	// NewDecoder that reads from file (Recommended when handling big files)
	// It doesn't keeps the whole in memory, and hence use less resources
	decoder := json.NewDecoder(file)
	// var omciLine map[string]interface{}
	var omciLine omciSchema.OmciContext
	latest := new(Diagram)
	latest.Name = Latest
	latest.Category = Latest
	latest.Chart = make(map[string]api.MeChart)
	pqMap := make(map[uint64]uint64)
	// Reads the array open bracket
	decoder.Token()
	// Decode reads the next JSON-encoded value from its input and stores it
	var omciMeta omciSchema.OmciMetaInfo
	for decoder.More() {
		err = decoder.Decode(&omciMeta)
		if err != nil {
			utils.Log("Decode failed")
			return nil, ontInfo
		}
		// Prints the first single JSON object
		err = json.Unmarshal([]byte(omciMeta.Omci), &omciLine)
		if err != nil {
			utils.Log("Decode omciMeta.Omci failed")
			return nil, ontInfo
		}
		omciJsonProc(omciLine, latest.Chart, pqMap)
	}
	resolveHostTpTypeEdges(latest.Chart)
	pqToTcontMapProc(latest.Chart, pqMap)
	diagramSet := tcontChartProc(pqMap)
	diagramSet = append(diagramSet, *latest)
	return diagramSet, ontInfo
}
