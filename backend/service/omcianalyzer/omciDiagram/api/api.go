/*
# ------------------------------------------------------------
# -- diagramDefine.go
# --
# -- Huang Minghe
# -- 2022-8-23
# ------------------------------------------------------------
*/
package api

import (
	"fmt"
	"strconv"
)

var VoipProtocol uint64

const (
	Protocol_SIP  = 1
	Protocol_H248 = 2
)

type MeRelationship struct {
	ShortName        string
	AssociationType  []uint64
	RelationshipType uint // 0 == ..> means Implicit, 1 == --> means Explicit
}

type MeNode struct {
	MeUID  string
	MeName string
	MeInst uint64
}

func (node *MeNode) MeNodeSet(name string, inst uint64) {
	node.MeName = name
	node.MeInst = inst
	node.MeUID = name + "_" + strconv.FormatUint(inst, 16)
}

type MeEdges struct {
	EdgeType       uint
	EdgeDirection  string
	EdgeLinkObject string
	EdgeNote       []string
}

func (edge *MeEdges) MeEdgesSet(eType uint, eDir string, name string, inst uint64, eNote []string) {
	edge.EdgeType = eType
	edge.EdgeDirection = eDir
	edge.EdgeLinkObject = name + "_" + strconv.FormatUint(inst, 16)
	edge.EdgeNote = eNote
}

type MeChart struct {
	Node  MeNode
	Edge  map[string]MeEdges
	Note  map[string]string
	Attrs map[string]interface{}
}

func (c *MeChart) Init(node MeNode, old *MeChart) *MeChart {
	if old != nil {
		c = old
	}

	if c.Edge == nil {
		c.Edge = make(map[string]MeEdges)
	}
	if c.Note == nil {
		c.Note = make(map[string]string)
	}
	if c.Attrs == nil {
		c.Attrs = make(map[string]interface{})
	}

	c.Node = node
	return c
}

func (c *MeChart) SetAttr(key string, value interface{}) {
	if c.Attrs == nil {
		c.Attrs = make(map[string]interface{})
	}

	switch v := value.(type) {
	case string:
		c.Attrs[key] = v
	case int, int32, int64:
		c.Attrs[key] = fmt.Sprintf("%d", v)
	case uint, uint32, uint64:
		c.Attrs[key] = fmt.Sprintf("%d", v)
	case float32, float64:
		c.Attrs[key] = fmt.Sprintf("%f", v)
	case fmt.Stringer:
		c.Attrs[key] = v.String()
	default:
		c.Attrs[key] = fmt.Sprintf("%v", v)
	}
}

func NodeWithoutEdge(old *MeChart, node MeNode, attrs map[string]interface{}) MeChart {
	chart := new(MeChart)
	if old != nil {
		chart = old
	} else {
		chart.Node = node
	}
	chart.Edge = nil
	chart.Note = nil
	return *chart
}

type MeChartHandler interface {
	GetMeChart(old *MeChart, node MeNode, attrs map[string]interface{}) MeChart
}

type MeChartHandlerBase struct{}

func (me *MeChartHandlerBase) GetMeChart(old *MeChart, node MeNode, attrs map[string]interface{}) MeChart {
	return NodeWithoutEdge(old, node, attrs)
}

type MeRelShip interface {
	GetRelShip() MeRelationship
}

var (
	MeRelationshipDef = make(map[uint64]MeRelShip)
	MeChartHandlerDef = make(map[uint64]MeChartHandler)
)

func MeRelationshipRegist(key uint64, relShip MeRelShip) {
	MeRelationshipDef[key] = relShip
}

func MeHandlerRegist(key uint64, handler MeChartHandler) {
	MeChartHandlerDef[key] = handler
}

func GetMeChartHandler(key uint64) MeChartHandler {
	if handler, ok := MeChartHandlerDef[key]; ok {
		return handler
	}
	return nil
}

func GetMeRelShip(key uint64) MeRelShip {
	if relShip, ok := MeRelationshipDef[key]; ok {
		return relShip
	}
	return nil
}
