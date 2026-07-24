package controller

import (
	"encoding/json"
	"net/http"
	"omciAnalyzer/controller/collector"
	"omciAnalyzer/global"
	"omciAnalyzer/models"
	"omciAnalyzer/utils"
	"strings"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (controller *AppController) CollectorConnect(context *gin.Context) {
	type connectInfo struct {
		OamIp    string `json:"oamIP"`
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	var con connectInfo
	if err := context.Bind(&con); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
		return
	}
	utils.Log("connect info: ", con)
	requestID := models.GenerateRequestID()
	oamIP := strings.TrimSpace(con.OamIp)
	collector.CreateTask(requestID, oamIP)
	collector.UpdateTaskStatus(requestID, collector.Running)
	// send kafka message
	kafkaKey := collector.GetKafkaKey(requestID)
	type connectInfoToKafka struct {
		RequestID string `json:"request_id"`
		MsgType   string `json:"msg_type"`
		Action    string `json:"action"`
		OamIp     string `json:"oam_ip"`
		UserName  string `json:"username"`
		Password  string `json:"password"`
	}
	var conToKafka connectInfoToKafka = connectInfoToKafka{
		RequestID: requestID,
		MsgType:   "connect_request",
		Action:    "get_topology",
		OamIp:     oamIP,
		UserName:  con.UserName,
		Password:  con.Password,
	}
	payload, _ := json.Marshal(conToKafka)
	global.KafkaSyncProducer(kafkaKey, payload, global.AppConf.Kafka.TopicCollectorRequest)
	conToKafka.Action = "get_topology_with_onus"
	payload, _ = json.Marshal(conToKafka)
	global.KafkaSyncProducer(kafkaKey, payload, global.AppConf.Kafka.TopicCollectorRequest)
	// respond with request ID
	context.JSON(http.StatusOK, gin.H{
		"request_id": requestID,
	})
}

func (controller *AppController) CollectorQuery(context *gin.Context) {
	id := context.Query("request_id")
	task, ok := collector.GetTask(id)
	if !ok {
		context.JSON(404, gin.H{"error": "task not found"})
		return
	}

	// check cache first
	oltCache := collector.GetOltInfoCacheList()
	if oltCache.Exists(task.OamIP) {
		utils.Log("OltInfo found in cache for OamIP:", task.OamIP)
		oltInfo, _ := oltCache.Get(task.OamIP)
		topology := oltInfo.GetTopologyResponse(id)

		// If task is still running in background, return cache but keep task alive
		if task.Status == collector.Running {
			context.JSON(http.StatusOK, gin.H{
				"status":     "Done",
				"result":     topology,
				"from_cache": true,
				"updating":   true,
			})
			return
		}

		// Task is completely done
		context.JSON(http.StatusOK, gin.H{
			"status":     "Done",
			"result":     topology,
			"from_cache": true,
		})
		return
	}

	if task.Status == collector.Failed {
		var errPayload struct {
			Error string `json:"error"`
		}
		if len(task.Result) > 0 {
			json.Unmarshal(task.Result, &errPayload)
		}
		context.JSON(http.StatusOK, gin.H{
			"status": "Failed",
			"error":  errPayload.Error,
		})
		return
	}

	if task.Status != collector.Done {
		context.JSON(http.StatusOK, gin.H{
			"status":     "Running",
			"result":     nil,
			"from_cache": true,
			"updating":   true,
		})
		return
	}

	// parse result
	var topology collector.RespOltTopology
	if err := json.Unmarshal(task.Result, &topology); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error":  "Failed to parse topology data",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": "Done",
		"result": topology,
	})
	// store in cache
	oltInfo := &collector.OltInfo{
		OamIP:  task.OamIP,
		NTInfo: topology.NTInfo,
	}
	oltCache.Set(oltInfo)
}

func handleCollectorResponse(msg *sarama.ConsumerMessage) {
	utils.Log("Received collector response message, key:", string(msg.Key), "value:", string(msg.Value))

	var errPayload struct {
		RequestID string `json:"request_id"`
		Error     string `json:"error"`
	}
	if err := json.Unmarshal(msg.Value, &errPayload); err == nil && errPayload.Error != "" {
		task, ok := collector.GetTask(errPayload.RequestID)
		if !ok {
			utils.Log("task not found for key:", errPayload.RequestID)
			return
		}
		collector.FailTask(task.RequestID, msg.Value)
		return
	}
	var logs struct {
		RequestID string          `json:"request_id"`
		Logs      []collector.Log `json:"logs"`
	}
	if err := json.Unmarshal(msg.Value, &logs); err == nil && len(logs.Logs) > 0 {
		task, ok := collector.GetTask(logs.RequestID)
		if !ok {
			utils.Log("task not found for key:", logs.RequestID)
			return
		}
		collector.SetTaskLogs(task.RequestID, logs.Logs)
		utils.Log("Set logs for task:", task.RequestID, "logs:", logs.Logs)
		return
	}
	var topology collector.RespOltTopology
	if err := json.Unmarshal(msg.Value, &topology); err == nil && topology.RequestID != "" {
		task, ok := collector.GetTask(topology.RequestID)
		if !ok {
			utils.Log("task not found for key:", topology.RequestID)
			return
		}
		collector.CompleteTask(task.RequestID, msg.Value)

		// Update cache with latest data
		oltCache := collector.GetOltInfoCacheList()
		oltInfo := &collector.OltInfo{
			OamIP:  task.OamIP,
			NTInfo: topology.NTInfo,
		}
		oltCache.Set(oltInfo)
		utils.Log("Updated cache for OamIP:", task.OamIP)

		return
	}
	utils.Log("wrong message, nothing to do", string(msg.Key), string(msg.Value))
}

func StartCollectorResponseConsumer() {
	go global.ConsumeTopic(
		global.AppConf.Kafka.TopicCollectorResponse,
		global.AppConf.Kafka.GroupCollectorResponse,
		handleCollectorResponse,
	)
}

func (controller *AppController) CollectorLoggerGet(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"logger_modules": collector.GetLoggerModules(),
	})
}

func (controller *AppController) CollectorLoggerSet(context *gin.Context) {
	type loggerModulesSettings struct {
		RequestID     string                        `json:"request_id"`
		Action        string                        `json:"action"`
		RequestParent string                        `json:"request_parent"`
		LoggerModules map[string][]collector.Module `json:"logger_modules"`
	}
	var settings loggerModulesSettings
	if err := context.Bind(&settings); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request payload",
		})
		return
	}
	utils.Log("Received logger settings for request ID:", settings.RequestID)
	utils.Log("Received logger settings for Action:", settings.Action)
	utils.Log("Received logger settings for request parent:", settings.RequestParent)
	utils.Log("Logger Modules:", settings.LoggerModules)
	// send kafka message
	kafkaKey := collector.GetKafkaKey(settings.RequestID)
	if kafkaKey == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request ID",
		})
		return
	}
	type loggerModulesSettingsToKafka struct {
		RequestID     string                        `json:"request_id"`
		MsgType       string                        `json:"msg_type"`
		OamIp         string                        `json:"oam_ip"`
		Action        string                        `json:"action"`
		RequestParent string                        `json:"request_parent"`
		LoggerModules map[string][]collector.Module `json:"logger_modules"`
	}
	var msgToKafka loggerModulesSettingsToKafka = loggerModulesSettingsToKafka{
		RequestID:     settings.RequestID,
		MsgType:       "logger_settings",
		OamIp:         collector.GetOamIP(settings.RequestID),
		Action:        settings.Action,
		RequestParent: settings.RequestParent,
		LoggerModules: settings.LoggerModules,
	}
	payload, _ := json.Marshal(msgToKafka)
	global.KafkaSyncProducer(kafkaKey, payload, global.AppConf.Kafka.TopicCollectorRequest)
	context.JSON(http.StatusOK, gin.H{
		"resp": "Logger settings received",
	})
}

func (controller *AppController) CollectorLoggerGetLogs(context *gin.Context) {
	id := context.Query("request_id")
	task, ok := collector.GetTask(id)
	if !ok {
		context.JSON(404, gin.H{"error": "task not found"})
		return
	}
	if task.Status != collector.Done {
		context.JSON(http.StatusBadRequest, gin.H{"error": "task not completed"})
		return
	}
	logs, ok := collector.GetTaskLogsPresignedURL(id)
	if !ok {
		context.JSON(404, gin.H{"error": "logs not found"})
		return
	}
	utils.Log("Presigned URL for logs:", logs)
	context.JSON(http.StatusOK, gin.H{
		"logs": logs,
	})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins, adjust for production
	},
}

func (controller *AppController) CollectorWebSocket(context *gin.Context) {
	requestID := context.Query("request_id")
	if requestID == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "request_id is required",
		})
		return
	}

	// Verify task exists
	_, ok := collector.GetTask(requestID)
	if !ok {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	// Upgrade connection to WebSocket
	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		utils.Log("Failed to upgrade WebSocket connection:", err)
		return
	}

	utils.Log("WebSocket connection established for request ID:", requestID)

	// Handle the connection
	wsManager := collector.GetWebSocketManager()
	wsManager.HandleConnection(conn, requestID)
}
