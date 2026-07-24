/**
 * @file omciOperator.go
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
	"encoding/json"
	"fmt"
	"omciAnalyzer/dao"
	service "omciAnalyzer/service/omcianalyzer"
)

func AnalyzerDataProc(log *dao.OmciAnalyzerRequestRecord) {
	var c = make(chan service.LogContent)
	var sc = make(chan service.RespStatusContent)

	go service.Analyzer(log.LogPath, log.LogDir, c, sc)
	//progress data
	go func() {
		for resp := range sc {
			dao.MysqlRecordDataUpdateMany(log, map[string]any{
				"progress":  resp.Progress,
				"status":    resp.Status,
				"error_msg": resp.ErrorMsg,
			})
		}
	}()
	//log data
	for content := range c {
		log.Content = content
		err := dao.MysqlRecordDataUpdate(log, "content_json", log.ContentJSON)
		if err != nil {
			fmt.Println("err: ", err)
		}
		// fmt.Println("log.Content: ", log.Content)
		log.LogType = content.LogType
		log.Progress = 100
		//save to DB
		old := new(dao.OmciAnalyzerRequestRecord)
		dao.MysqlRecordDataReadFirst(old, "id", log.ID)
		err = dao.MysqlRecordDataUpdateByRecord(old, log)
		if err != nil {
			fmt.Println("err: ", err)
		}
		//send diagram to AI
		// go sendDigramToAI(content, log.RequestKey)
	}
}

func sendDigramToAI(content service.LogContent, key string) {
	for k, v := range content.OnuContent {
		for _, v := range v.Charts {
			if v.Category == "latest" {
				//to json
				v.Name = k
				_, err := json.Marshal(v)
				if err != nil {
					continue
				}
				//build kafka msg
				// handler := GetKafkaMsgHandler(TagOmciAI)
				// if handler != nil {
				// 	tag := handler.GetMsgName()
				// 	msg := handler.RequestProcess(tag, key, string(jsonBytes))
				// 	fmt.Println("msg: ", msg)
				// 	//send to kafka
				// 	// models.MqSyncPush(msg.ToString(), global.AppConf.Kafka.TopicOmciAI)
				// }
			}
		}
	}
}
