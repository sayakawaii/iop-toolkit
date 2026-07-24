package omciSchema

import (
	"encoding/hex"
	"encoding/json"
	"omciAnalyzer/service/omcianalyzer/omciShape"
	"omciAnalyzer/utils"
	"os"
	"slices"
	"time"
)

type OmciMetaInfo struct {
	Timestamp time.Time `json:"Timestamp"`
	Omci      string    `json:"Omci"`
}

func in[T comparable](val T, list ...T) bool {
	return slices.Contains(list, val)
}

func omciParse(omci string, schemaDef *SchemaDef) string {
	if len(omci) != 0 {
		payloadString := string(omci)
		payload, _ := hex.DecodeString(payloadString)
		if payload[3] == 0x0A && len(payload) < 40 {
			return "" //skip illegal omci
		}
		if payload[2]&0x40 == 0x40 {
			omciMsgInterface, _ := OmciParseMsgInterface(TX, payload, schemaDef)
			if omciMsgInterface == nil {
				utils.Log("TX omciMsgInterface nil : " + string(omci))
			} else if v, e := OmciParseFromInterface(TX, omciMsgInterface); e == nil {
				return v.EasyString()
			}
		} else {
			omciMsgInterface, _ := OmciParseMsgInterface(RX, payload, schemaDef)
			if omciMsgInterface == nil {
				utils.Log("RX omciMsgInterface nil : " + string(omci))
			} else if v, e := OmciParseFromInterface(RX, omciMsgInterface); e == nil {
				return v.EasyString()
			}
		}
	}
	return ""
}

func OmciRawDataParse(logPath string, outputPath string, logType uint8) string {
	utils.Log("omci raw data parse start")
	file, err := os.Open(logPath)
	if err != nil {
		utils.Log("open file failed")
		return ""
	} else {
		utils.Log("open file success")
	}

	defer file.Close()
	var tmpSchemaDef *SchemaDef
	if !in(logType, omciShape.LightSpan, omciShape.NokiaOnu, omciShape.CopperOnu, omciShape.Isam) {
		utils.Log("StandardSchemaDef using")
		tmpSchemaDef = StandardSchemaDef()
	} else {
		utils.Log("DefaultSchemaDef using")
		tmpSchemaDef = DefaultSchemaDef()
	}
	decoder := json.NewDecoder(file)
	var omciMetaList []OmciMetaInfo
	var omciMeta omciShape.OmciMetaData
	decoder.Token()
	for decoder.More() {
		err = decoder.Decode(&omciMeta)
		if err != nil {
			utils.Log("Decode failed")
			return ""
		}
		omciMetaList = append(omciMetaList, OmciMetaInfo{
			Timestamp: omciMeta.Timestamp,
			Omci:      omciParse(omciMeta.RawData, tmpSchemaDef),
		})
	}

	writer, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		utils.Log("open file failed")
		return ""
	}
	defer writer.Close()
	jsonData, err := json.MarshalIndent(omciMetaList, "", "  ")
	if err != nil {
		utils.Log("json marshal failed")
		return ""
	}
	writer.Write(jsonData)
	return outputPath
}
