package main

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

// Node 代表 PlantUML 中的对象节点
type Node struct {
	id          int64
	alias       string // 例如 MBSP_101
	displayName string // 例如 "MBSP"
}

func (n Node) ID() int64 { return n.id }

func main() {
	// 注意：这里只使用了第一个 @startuml ... @enduml 部分内容
	plantUML := `
@startuml
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

object "MBSP" as MBSP_101
MBSP_101 : 0x101
object "PPTP" as PPTP_104
PPTP_104 : 0x104
object "VTFD" as VTFD_2
VTFD_2 : 0x2
note right of VTFD_2
FWDop: 16
numOfEntr: 1
vlanID: 2000
end note
object "MBPCD" as MBPCD_104
MBPCD_104 : 0x104
object "MBPCD" as MBPCD_601
MBPCD_601 : 0x601
object "TCONT" as TCONT_8001
TCONT_8001 : 0x8001
object "PQ" as PQ_8008
PQ_8008 : 0x8008
object "MBPCD" as MBPCD_1
MBPCD_1 : 0x1
object "MBPCD" as MBPCD_103
MBPCD_103 : 0x103
object "MBSP" as MBSP_104
MBSP_104 : 0x104
object "TCONT" as TCONT_8000
TCONT_8000 : 0x8000
object "PMapper" as PMapper_1
PMapper_1 : 0x1
object "EVTOCD" as EVTOCD_101
EVTOCD_101 : 0x101
note right of EVTOCD_101
(14, 4096, 5), (14, 4096, 0), 0 || 3, (15, 0, 0), (15, 4096, 3)
(15, 4096, 0), (15, 0, 5), 0 || 3, (15, 0, 0), (15, 0, 2)
(15, 4096, 0), (14, 4096, 5), 0 || 3, (15, 0, 0), (15, 4096, 2)
end note
object "PPTP" as PPTP_102
PPTP_102 : 0x102
object "EVTOCD" as EVTOCD_102
EVTOCD_102 : 0x102
note right of EVTOCD_102
(15, 4096, 0), (14, 4096, 5), 0 || 3, (15, 0, 0), (15, 4096, 2)
(14, 4096, 5), (14, 4096, 0), 0 || 3, (15, 0, 0), (15, 4096, 3)
(15, 4096, 0), (15, 0, 5), 0 || 3, (15, 0, 0), (15, 0, 2)
end note
object "PPTP" as PPTP_103
PPTP_103 : 0x103
object "MBSP" as MBSP_103
MBSP_103 : 0x103
object "EVTOCD" as EVTOCD_104
EVTOCD_104 : 0x104
note right of EVTOCD_104
(14, 4096, 5), (14, 4096, 0), 0 || 3, (15, 0, 0), (15, 4096, 3)
(15, 4096, 0), (15, 2000, 0), 0 || 0, (15, 0, 0), (3, 2000, 2)
(15, 4096, 0), (14, 4096, 5), 0 || 3, (15, 0, 0), (15, 4096, 2)
end note
object "GemITP" as GemITP_107
GemITP_107 : 0x107
object "GAL" as GAL_0
GAL_0 : 0x0
note right of GAL_0
MTU:48
end note
object "MBPCD" as MBPCD_101
MBPCD_101 : 0x101
object "MBSP" as MBSP_601
MBSP_601 : 0x601
object "PQ" as PQ_8000
PQ_8000 : 0x8000
object "PPTP" as PPTP_101
PPTP_101 : 0x101
object "MBPCD" as MBPCD_2
MBPCD_2 : 0x2
object "GemCTP" as GemCTP_107
GemCTP_107 : 0x107
object "GemITP" as GemITP_10a
GemITP_10a : 0x10a
object "MBSP" as MBSP_102
MBSP_102 : 0x102
object "EVTOCD" as EVTOCD_103
EVTOCD_103 : 0x103
note right of EVTOCD_103
(15, 4096, 0), (14, 4096, 5), 0 || 3, (15, 0, 0), (15, 4096, 2)
(14, 4096, 5), (14, 4096, 0), 0 || 3, (15, 0, 0), (15, 4096, 3)
(15, 4096, 0), (15, 0, 5), 0 || 3, (15, 0, 0), (15, 0, 2)
end note
object "PMapper" as PMapper_2
PMapper_2 : 0x2
object "MBPCD" as MBPCD_102
MBPCD_102 : 0x102
object "VEIP" as VEIP_601
VEIP_601 : 0x601
object "EVTOCD" as EVTOCD_601
EVTOCD_601 : 0x601
note right of EVTOCD_601
(14, 4096, 5), (14, 4096, 0), 0 || 3, (15, 0, 0), (15, 4096, 3)
(15, 4096, 0), (15, 0, 5), 0 || 3, (15, 0, 0), (15, 0, 2)
(15, 4096, 0), (8, 1500, 0), 0 || 1, (15, 0, 0), (8, 1500, 2)
(8, 1500, 0), (8, 4096, 0), 0 || 1, (15, 0, 0), (9, 1500, 3)
(15, 4096, 0), (14, 4096, 5), 0 || 3, (15, 0, 0), (15, 4096, 2)
end note
object "VTFD" as VTFD_1
VTFD_1 : 0x1
note right of VTFD_1
vlanID: 1500
FWDop: 16
numOfEntr: 1
end note
object "GemCTP" as GemCTP_10a
GemCTP_10a : 0x10a
VTFD_2 ..> MBPCD_2
MBPCD_1 -down-> MBSP_601
MBPCD_1 -up-> PMapper_1
MBPCD_104 -down-> PPTP_104
MBPCD_104 -up-> MBSP_104
MBPCD_601 -up-> MBSP_601
MBPCD_601 -down-> VEIP_601
PMapper_1 "6" -up-> GemITP_107
PMapper_1 "7" -up-> GemITP_107
PMapper_1 "0" -up-> GemITP_107
PMapper_1 "1" -up-> GemITP_107
PMapper_1 "2" -up-> GemITP_107
PMapper_1 "3" -up-> GemITP_107
PMapper_1 "4" -up-> GemITP_107
PMapper_1 "5" -up-> GemITP_107
MBPCD_103 -up-> MBSP_103
MBPCD_103 -down-> PPTP_103
EVTOCD_102 -left-> PPTP_102
EVTOCD_104 -left-> PPTP_104
GemITP_107 -up-> GemCTP_107
EVTOCD_101 -left-> PPTP_101
MBPCD_101 -up-> MBSP_101
MBPCD_101 -down-> PPTP_101
GemCTP_107 -up-> PQ_8008
GemITP_10a -up-> GemCTP_10a
MBPCD_2 -down-> MBSP_104
MBPCD_2 -up-> PMapper_2
PMapper_2 "7" -up-> GemITP_10a
PMapper_2 "0" -up-> GemITP_10a
PMapper_2 "1" -up-> GemITP_10a
PMapper_2 "2" -up-> GemITP_10a
PMapper_2 "3" -up-> GemITP_10a
PMapper_2 "4" -up-> GemITP_10a
PMapper_2 "5" -up-> GemITP_10a
PMapper_2 "6" -up-> GemITP_10a
EVTOCD_103 -left-> PPTP_103
EVTOCD_601 -left-> VEIP_601
VTFD_1 ..> MBPCD_1
GemCTP_10a -up-> PQ_8000
MBPCD_102 -down-> PPTP_102
MBPCD_102 -up-> MBSP_102

@enduml
`

	// 正则表达式分别用于提取节点和边（此处简化处理，不考虑 note 块等其他语法）
	nodeRe := regexp.MustCompile(`object\s+"([^"]+)"\s+as\s+(\w+)`)
	edgeRe := regexp.MustCompile(`(\w+)(?:\s+"[^"]+")?\s+(?:(?:-(?:up|down|left|right|auto)-)|(?:[-.]+))>\s*(\w+)`)

	// 构造有向图
	g := simple.NewDirectedGraph()
	nodeMap := make(map[string]graph.Node)
	var nextID int64 = 1

	// 第一遍扫描：提取所有节点定义
	scanner := bufio.NewScanner(strings.NewReader(plantUML))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if matches := nodeRe.FindStringSubmatch(line); matches != nil {
			displayName := matches[1]
			alias := matches[2]
			if _, exists := nodeMap[alias]; !exists {
				n := Node{id: nextID, alias: alias, displayName: displayName}
				nextID++
				nodeMap[alias] = n
				g.AddNode(n)
			}
		}
	}

	// 第二遍扫描：提取边的定义
	scanner = bufio.NewScanner(strings.NewReader(plantUML))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if matches := edgeRe.FindStringSubmatch(line); matches != nil {
			fromAlias := matches[1]
			toAlias := matches[2]
			fromNode, ok := nodeMap[fromAlias]
			if !ok {
				// 如果边中出现了未定义的节点，则创建一个简单节点
				n := Node{id: nextID, alias: fromAlias, displayName: fromAlias}
				nextID++
				nodeMap[fromAlias] = n
				g.AddNode(n)
				fromNode = n
			}
			toNode, ok := nodeMap[toAlias]
			if !ok {
				n := Node{id: nextID, alias: toAlias, displayName: toAlias}
				nextID++
				nodeMap[toAlias] = n
				g.AddNode(n)
				toNode = n
			}
			g.SetEdge(g.NewEdge(fromNode, toNode))
		}
	}

	fmt.Println(g)
	printGraphDot(g)
	// 示例：查找从节点 "VTFD_2" 到 "PQ_8000" 的路径
	startAlias := "PPTP_104"
	endAlias := "PQ_8000"

	startNode, ok := nodeMap[startAlias]
	if !ok {
		log.Fatalf("起始节点 %s 不存在", startAlias)
	}
	endNode, ok := nodeMap[endAlias]
	if !ok {
		log.Fatalf("目标节点 %s 不存在", endAlias)
	}

	path := findPathBidirectional(g, startNode, endNode, make(map[int64]bool))
	if path == nil {
		fmt.Printf("没有找到从 %s 到 %s 的路径\n", startAlias, endAlias)
	} else {
		fmt.Printf("从 %s 到 %s 的路径：\n", startAlias, endAlias)
		for _, n := range path {
			// 将 graph.Node 转换为我们自定义的 Node 类型
			node := n.(Node)
			fmt.Printf("%s (%s) -> ", node.displayName, node.alias)
		}
		fmt.Println("end")
	}
}

// findPath 通过深度优先搜索查找从 start 到 target 的一条路径
func findPath(g graph.Graph, start, target graph.Node, visited map[int64]bool) []graph.Node {
	if start.ID() == target.ID() {
		return []graph.Node{start}
	}
	visited[start.ID()] = true
	for neighbors := g.From(start.ID()); neighbors.Next(); {
		neighbor := neighbors.Node()
		if visited[neighbor.ID()] {
			continue
		}
		if path := findPath(g, neighbor, target, visited); path != nil {
			return append([]graph.Node{start}, path...)
		}
	}
	return nil
}

func findPathBidirectional(g graph.Graph, start, target graph.Node, visited map[int64]bool) []graph.Node {
	if start.ID() == target.ID() {
		return []graph.Node{start}
	}
	visited[start.ID()] = true

	// 收集出边和入边的所有邻居
	var neighbors []graph.Node
	for it := g.From(start.ID()); it.Next(); {
		neighbors = append(neighbors, it.Node())
	}
	if dg, ok := g.(graph.Directed); ok {
		for it := dg.To(start.ID()); it.Next(); {
			neighbors = append(neighbors, it.Node())
		}
	}
	for _, neighbor := range neighbors {
		if visited[neighbor.ID()] {
			continue
		}
		if path := findPathBidirectional(g, neighbor, target, visited); path != nil {
			return append([]graph.Node{start}, path...)
		}
	}
	return nil
}

func printGraphDot(g *simple.DirectedGraph) {
	b, err := dot.Marshal(g, "MyGraph", "", "  ")
	if err != nil {
		log.Fatalf("导出 DOT 格式失败: %v", err)
	}
	fmt.Println(string(b))
}
