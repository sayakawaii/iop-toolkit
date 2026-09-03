package omci

import (
	"encoding/json"
	"fmt"
	"omciAnalyzer/controller/omci/api"
	_ "omciAnalyzer/controller/omci/msgType"
	"omciAnalyzer/service/omcianalyzer/omciSchema"
	"omciAnalyzer/utils"
	"os"
	"time"

	"github.com/iancoleman/orderedmap"
)

type Contents struct {
	Header  omciSchema.OmciHeader  `json:"header"`
	Payload *orderedmap.OrderedMap `json:"payload"`
}

type RespOmciData struct {
	Id        uint64   `json:"id"`
	Latency   float64  `json:"latency"`
	TCID      string   `json:"tcid"`
	Name      string   `json:"name"`
	Class     uint64   `json:"class"`
	Type      string   `json:"type"`
	MsgFormat string   `json:"msgFormat"`
	Direction string   `json:"direction"`
	Status    string   `json:"status"`
	Content   Contents `json:"content"`
}

func OmciContentTable(msg *omciSchema.OmciContext) (r *orderedmap.OrderedMap) {
	h := api.GetTypeHandler(msg.Header.MsgType)
	if h != nil {
		r, _ = h.HandleMessage(msg)
	} else {
		r = nil
	}
	return prependMsgFormat(r, msg.Header.DevId)
}

func prependMsgFormat(payload *orderedmap.OrderedMap, devId byte) *orderedmap.OrderedMap {
	enriched := orderedmap.New()
	enriched.Set("MsgFormat", omciSchema.FormatDevId(devId))
	if payload == nil {
		return enriched
	}
	for _, key := range payload.Keys() {
		val, _ := payload.Get(key)
		enriched.Set(key, val)
	}
	return enriched
}

func enrichOmciHeader(header omciSchema.OmciHeader) omciSchema.OmciHeader {
	header.MsgFormat = omciSchema.FormatDevId(header.DevId)
	return header
}

func OmciContentResult(msg *omciSchema.OmciContext) uint8 {
	meAction := msg.Header.MsgTypeName
	switch meAction {
	case "MIB Upload":
		// handle mibupload
		return 0
	case "MIB Upload Next":
		// handle mibupload next
		return 0
	case "Alarm":
		// handle alarm
		return 0
	default:
		if msg.Header.AK == 1 {
			// handle service
			contents := omciSchema.OmciAkContents{}
			err := omciSchema.GetContentFromInterface(msg.Contents, &contents)
			if err != nil {
				return 2
			}
			if (uint8)(contents.Result) != 0 {
				return 1
			} else {
				return 0
			}
		} else {
			return 0
		}
	}
}

func updateLatencyMap(latencyMap map[uint64]float64, tcid uint64, timestamp time.Time) {
	currentMilliseconds := float64(timestamp.UnixNano()) / float64(time.Millisecond)
	if latencyMap[tcid] == 0 {
		latencyMap[tcid] = currentMilliseconds
	} else {
		latencyMap[tcid] = currentMilliseconds - latencyMap[tcid]
	}
}

func AssembleOmciData(filePath string) []RespOmciData {
	status := []string{"success", "warning", "danger", "info"}
	var latencyMap = map[uint64]float64{}

	var omciData []RespOmciData
	file, err := os.Open(filePath)
	if err != nil {
		utils.Log("open file failed: " + filePath)
		return omciData
	}
	defer file.Close()
	var msg omciSchema.OmciContext
	var omciMeta omciSchema.OmciMetaInfo
	var index uint64
	decoder := json.NewDecoder(file)
	decoder.Token()
	for decoder.More() {
		err = decoder.Decode(&omciMeta)
		if err != nil {
			utils.Log("Decode failed")
			continue
		}
		err = json.Unmarshal([]byte(omciMeta.Omci), &msg)
		if err != nil {
			utils.Log("Decode omciMeta.Omci failed")
			continue
		}
		if msg.Header.Tcid != 0 {
			updateLatencyMap(latencyMap, msg.Header.Tcid, omciMeta.Timestamp)
		}
		index++
		var dir string
		var latency float64
		if msg.Header.AR == 1 {
			dir = "OLT --> ONU"
			latency = 0
		} else if msg.Header.AK == 1 {
			dir = "OLT <-- ONU"
			latency = latencyMap[msg.Header.Tcid]
		} else {
			dir = "OLT <-- ONU"
			latency = 0
		}
		result := OmciContentResult(&msg)
		msgFormat := omciSchema.FormatDevId(msg.Header.DevId)
		omciData = append(omciData, RespOmciData{
			Id:        index,
			Latency:   latency,
			TCID:      fmt.Sprintf("%d(0x%x)", msg.Header.Tcid, msg.Header.Tcid),
			Name:      msg.Header.MeClassName,
			Class:     msg.Header.MeClass,
			Type:      msg.Header.MsgTypeName,
			MsgFormat: msgFormat,
			Direction: dir,
			Status:    status[result],
			Content: Contents{
				Header:  enrichOmciHeader(msg.Header),
				Payload: OmciContentTable(&msg),
			},
		})
	}
	return omciData
}
