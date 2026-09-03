/**
 * @file yangHelper.go
 * @author Minghe Huang (minghe.huang@nokia-sbell.com)
 * @brief
 * @version 0.1
 * @date 2023-11-22
 *
 * @copyright Copyright (c) 2023
 *
 */

package api

import (
	"fmt"
	"io"
	"omciAnalyzer/dao"
	"omciAnalyzer/models"
	"omciAnalyzer/utils"
	"os"
	"strings"
)

func OnuLogShape(logPath string, outputPath string) map[string]dao.ConfigAnalyzerContent {
	utils.Log("yang helper onu log shape")
	onus := make(map[string]dao.ConfigAnalyzerContent)
	var ontid string = ""
	var copyBuff string = ""

	file, err := os.Open(logPath)
	if err != nil {
		utils.Log("open file failed: " + fmt.Sprintf("%s", err))
		return nil
	}

	defer file.Close()
	reader := models.GetReader(file)
	line := ""
	for {
		buf, isPrefix, err := reader.ReadLine()
		line += string(buf)
		if isPrefix {
			continue
		}
		if err == io.EOF {
			utils.Log("end of read")
			break
		}
		if len(line) != 0 {
			str := string(line)
			if strings.Contains(str, "handleProxyMsgs: ") && (strings.Contains(str, "\"operation\":\"copy\"") || strings.Contains(str, "\"operation\":\"update\"")) {
				leftIndex := strings.LastIndex(str, "\"onu_name\":\"")
				rightIndex := strings.LastIndex(str, "\",\"event\"")
				ontidstr := str[leftIndex:rightIndex]
				ontid = models.StringStrip(string(ontidstr[len("\"onu_name\":\""):]))
				if _, res := onus[ontid]; !res {
					onu := dao.ConfigAnalyzerContent{}
					onu.Copy.OnuNativeName = string(ontidstr[len("\"onu_name\":\""):])
					onu.Copy.Identifier, _ = models.GetSubstringBetween(str, "\"identifier\":\"", "\",\"operation\"")
					onu.Copy.Operation, _ = models.GetSubstringBetween(str[strings.LastIndex(str, "\"operation\":\""):strings.LastIndex(str, "\"operation\":\"")+21], "\"operation\":\"", "\",")
					onu.Copy.OnuCopyDir = outputPath
					onu.Copy.OnuCopyName = ontid
					onu.Copy.NewActionKey()
					onu.Copy.OnuCopyPath = outputPath + onu.Copy.GetActionKey() + ".json"
					utils.Log("path: " + onu.Copy.OnuCopyPath)
					copyBuff = string(str[strings.LastIndex(str, "{\"onu_name\":"):])
					models.WriteFile(onu.Copy.OnuCopyPath, copyBuff)
					onus[ontid] = onu
				}
			}
			line = ""
		}
	}

	return onus
}

func OnuJsonBound(key string, ontid string, outputPath string) map[string]dao.ConfigAnalyzerContent {
	utils.Log("yang helper onu json bound")
	onus := make(map[string]dao.ConfigAnalyzerContent)

	if _, res := onus[ontid]; !res {
		onu := dao.ConfigAnalyzerContent{}
		onu.Copy.OnuNativeName = ontid
		onu.Copy.Identifier = "0"
		onu.Copy.Operation = "dummy"
		onu.Copy.OnuCopyDir = outputPath
		onu.Copy.OnuCopyName = ontid
		onu.Copy.SetActionKey(key)
		onu.Copy.OnuCopyPath = outputPath + key + ".json"
		onus[ontid] = onu
	}

	return onus
}

func SendDataToYangHelper(log *dao.ConfigAnalyzerRequestRecord, ponLogPath string) {
	//build kafka msg
	// handler := taskDistribute.GetKafkaMsgHandler(taskDistribute.TagConfigAnalyzer)
	// if handler != nil {
	// 	for _, onu := range log.Onus {
	// 		onu.Finished = false
	// 		payload := taskDistribute.ConfigAnalyzerRequest{}
	// 		token := taskDistribute.ConfigAnalyzerToken{}
	// 		token.Name = onu.Copy.OnuNativeName
	// 		token.File = "http://" + global.AppConf.Server.Addr + "/" + onu.Copy.OnuCopyPath
	// 		// token.File = "http://135.251.206.244:3334" + "/cp1_onu01.json"
	// 		// payload.Path = append(payload.Path, "http://"+global.AppConf.Server.Addr+"/"+thisLog.LogPath)
	// 		payload.Path = append(payload.Path, token)
	// 		// payload.Path = append(payload.Path, "http://135.251.206.244:3334"+"/cp1_onu01.json")
	// 		if ponLogPath != "" {
	// 			token.Name = ""
	// 			token.File = "http://" + global.AppConf.Server.Addr + "/" + ponLogPath
	// 			payload.Path = append(payload.Path, token)
	// 		}
	// 		tag := handler.GetMsgName()
	// 		msg := handler.RequestProcess(tag, log.RequestKey, payload)
	// 		fmt.Println("msg: ", msg)
	// 		//send to kafka
	// 		// models.MqSyncPush(msg.ToString(), global.AppConf.Kafka.TopicConfigAnalyzer)
	// 	}
	// }
}
