/*
# ------------------------------------------------------------
# -- configAnalyzerMsg.go
# --
# -- Huang Minghe
# -- 2023-02-02
# ------------------------------------------------------------
*/
package api

import (
	"errors"
	"fmt"
	"net/url"
	"omciAnalyzer/dao"
	"omciAnalyzer/global"
	"omciAnalyzer/models"
	"omciAnalyzer/service/plantUml"
	"omciAnalyzer/utils"
	"os"
	"path"
	"strings"
)

const (
	TagConfigAnalyzer = "configAnalyzer"
)

type ConfigAnalyzerToken struct {
	Name string `json:"key"`
	File string `json:"url"`
}

type ConfigAnalyzerRequest struct {
	Path []ConfigAnalyzerToken `json:"path_list"`
}

// type ConfigAnalyzerResponse struct {
// 	Plantuml string `json:"plantuml"`
// 	Report   string `json:"report"`
// }

type YangHelperGraph struct {
	Key     string `json:"key"`
	Content string `json:"content"`
}

type ConfigAnalyzerResponse struct {
	Graphics []YangHelperGraph     `json:"graph_list"`
	PathList []ConfigAnalyzerToken `json:"path_list"`
}

type configAnalyzer struct {
	// // msgInfo
	// requestHandlerBase
}

// func (msg *configAnalyzer) RequestProcess(tag string, key string, content interface{}) *KafkaMsg {
// 	kafkaMsg := new(KafkaMsg)
// 	kafkaMsg.Tag = tag
// 	kafkaMsg.Pkg.Key = key
// 	kafkaMsg.Pkg.Content = content
// 	fmt.Println("child pkg: ", content)
// 	return kafkaMsg
// }

func umlStringCreate(dir string, graph []YangHelperGraph) string {
	uml := ""
	layer2UmlList := make(map[string]string)
	linkStr := ""
	for _, v := range graph {
		if v.Key == "0" {
			uml = v.Content
		} else {
			token := "[[[" + v.Key + " "
			path := dir + "_" + v.Key + ".wsd"
			linkStr = "[[[http://" + global.AppConf.Server.Addr + "/" + dir + "_" + v.Key + ".svg" + " "
			// writePlantUmlFile(path, v.Content)
			//step1 handle layer1 uml link
			uml = strings.ReplaceAll(uml, token, linkStr)
			//step2 handle layer2 uml link
			for k, v := range layer2UmlList {
				v = strings.ReplaceAll(v, token, linkStr)
				layer2UmlList[k] = v
			}
			//step3 cache current object
			if strings.Contains(v.Content, "@startuml\njson JSON") {
				writePlantUmlFile(path, v.Content)
			} else {
				layer2UmlList[path] = v.Content
			}
		}
	}
	for k, v := range layer2UmlList {
		writePlantUmlFile(k, v)
	}
	return uml
}

func writePlantUmlFile(fileName string, line string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		utils.Log("open file failed")
	} else {
		// utils.Log("open file success")
		defer file.Close()
		file.WriteString(line)
	}
}

func getResponseActionKey(filePath []ConfigAnalyzerToken) string {
	key := ""
	for _, token := range filePath {
		u, _ := url.Parse(token.File)
		filename := path.Base(u.Path)
		extName := path.Ext(filename)
		if extName == ".json" {
			key = strings.TrimSuffix(filename, extName)
			utils.Log("response action key: " + key)
		}
	}
	return key
}

func (msg *configAnalyzer) ResponseProcess(tag string, key string, content interface{}) error {
	var response ConfigAnalyzerResponse
	var err error
	// fmt.Println("pkg: ", content)
	err = models.GetContentFromInterface(content, &response)
	if err != nil {
		return err
	}
	// fmt.Println("response: ", response)
	record := new(dao.ConfigAnalyzerRequestRecord)
	dao.MysqlRecordDataRead(record, "request_key", key)
	// fmt.Println("record: ", record)
	if record.ID == 0 {
		return errors.New("no match record")
	}
	//update progress
	dao.MysqlRecordDataUpdate(record, "progress", 50)
	finish := true
	actionKey := getResponseActionKey(response.PathList)
	for key, onu := range record.Onus {
		// utils.Log("response action key: " + actionKey + "local action key: " + onu.Copy.GetActionKey())
		if actionKey == onu.Copy.GetActionKey() {
			utils.Log("action key match")
			//update response info
			onu.Response.PlantUmlPath = onu.Copy.OnuCopyDir + onu.Copy.ActionKey + ".wsd"
			onu.Response.DiagramPath = onu.Copy.OnuCopyDir + onu.Copy.ActionKey + ".svg"
			onu.Finished = true
			// umlStr := response.Graphics[0].Content
			umlStr := umlStringCreate(onu.Copy.OnuCopyDir+onu.Copy.ActionKey, response.Graphics)
			writePlantUmlFile(onu.Response.PlantUmlPath, umlStr)
			fmt.Println("onu: ", onu)
			record.Onus[key] = onu
		}
		finish = finish && onu.Finished
	}
	//update record
	err = dao.MysqlRecordDataUpdate(record, "onus_json", record.OnusJSON)
	if err != nil {
		fmt.Println("err: ", err)
	}
	old := new(dao.ConfigAnalyzerRequestRecord)
	dao.MysqlRecordDataReadFirst(old, "id", record.ID)
	err = dao.MysqlRecordDataUpdateByRecord(old, record)
	if err != nil {
		fmt.Println("err: ", err)
	}

	if finish {
		// plantUml.Draw(record.PlantUmlPath, record.LogPath)
		plantUml.Draw(record.LogDir+"*.wsd", record.LogPath)
		//finish
		// dao.MysqlRecordDataUpdate(record, "progress", 100)
		//save to DB
		record.Progress = 100
		old := new(dao.ConfigAnalyzerRequestRecord)
		dao.MysqlRecordDataReadFirst(old, "id", record.ID)
		err = dao.MysqlRecordDataUpdateByRecord(old, record)
		if err != nil {
			fmt.Println("err: ", err)
		}
	}
	return err
}

func init() {
	// msg := new(configAnalyzer)
	// msg.msgInfo.name = TagConfigAnalyzer
	// kafkaMsgHandlerRegist(TagConfigAnalyzer, msg)
}
