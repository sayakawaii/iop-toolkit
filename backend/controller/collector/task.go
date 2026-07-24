package collector

import (
	"omciAnalyzer/api"
	"omciAnalyzer/utils"
	"sync"
	"time"
)

type TaskStatus int

var taskStore sync.Map // map[request_id]*Task

const (
	Pending TaskStatus = iota
	Running
	Done
	Failed
)

type Log struct {
	App   string  `json:"app"`
	Minio string  `json:"minio"`
	Size  float64 `json:"size"`
}

type Task struct {
	RequestID string
	KafkaKey  string
	OamIP     string
	Status    TaskStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	Result    []byte
	Logs      []Log
}

func taskGC() {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		now := time.Now()
		taskStore.Range(func(k, v any) bool {
			task := v.(*Task)
			if now.Sub(task.UpdatedAt) > 24*time.Hour {
				taskStore.Delete(k)
			}
			return true
		})
	}
}

func CreateTask(requestID, oamIP string) *Task {
	now := time.Now()
	task := &Task{
		RequestID: requestID,
		KafkaKey:  "olt:" + oamIP,
		OamIP:     oamIP,
		Status:    Pending,
		CreatedAt: now,
		UpdatedAt: now,
	}
	taskStore.Store(requestID, task)
	return task
}

func GetTask(requestID string) (*Task, bool) {
	v, ok := taskStore.Load(requestID)
	if !ok {
		return nil, false
	}
	return v.(*Task), true
}

func GetKafkaKey(requestID string) string {
	v, ok := taskStore.Load(requestID)
	if !ok {
		return ""
	}
	task := v.(*Task)
	return task.KafkaKey
}

func GetOamIP(requestID string) string {
	v, ok := taskStore.Load(requestID)
	if !ok {
		return ""
	}
	task := v.(*Task)
	return task.OamIP
}

func UpdateTaskStatus(requestID string, status TaskStatus) bool {
	v, ok := taskStore.Load(requestID)
	if !ok {
		return false
	}

	task := v.(*Task)
	task.Status = status
	task.UpdatedAt = time.Now()
	return true
}

func CompleteTask(requestID string, result []byte) bool {
	v, ok := taskStore.Load(requestID)
	if !ok {
		return false
	}

	task := v.(*Task)
	task.Status = Done
	task.Result = result
	task.UpdatedAt = time.Now()

	// Notify WebSocket clients about the update
	wsManager := GetWebSocketManager()
	wsManager.NotifyUpdate(requestID, string(result))

	return true
}

func FailTask(requestID string, result []byte) bool {
	v, ok := taskStore.Load(requestID)
	if !ok {
		return false
	}

	task := v.(*Task)
	task.Status = Failed
	task.Result = result
	task.UpdatedAt = time.Now()
	return true
}

func GetTaskLogs(requestID string) ([]Log, bool) {
	v, ok := taskStore.Load(requestID)
	if !ok {
		return nil, false
	}

	task := v.(*Task)
	return task.Logs, true
}

func SetTaskLogs(requestID string, logs []Log) bool {
	v, ok := taskStore.Load(requestID)
	if !ok {
		return false
	}

	task := v.(*Task)
	task.Logs = logs
	task.UpdatedAt = time.Now()
	return true
}

type PresignedURL struct {
	Log
	URL string `json:"url"`
}

func GetTaskLogsPresignedURL(requestID string) ([]PresignedURL, bool) {
	v, ok := taskStore.Load(requestID)
	if !ok {
		return nil, false
	}

	task := v.(*Task)
	if len(task.Logs) == 0 {
		return []PresignedURL{}, true
	}

	var presignedURLs []PresignedURL
	for _, log := range task.Logs {
		// Generate presigned URL for each log file with 1 hour expiration
		url, err := api.GeneratePresignedURL(log.Minio, 1*time.Hour)
		if err != nil {
			// If error occurs, append empty string to maintain index consistency
			utils.Log("Failed to generate presigned URL for log file:", log.Minio, "error:", err)
			presignedURLs = append(presignedURLs, PresignedURL{
				Log: log,
				URL: "",
			})
		} else {
			presignedURLs = append(presignedURLs, PresignedURL{
				Log: log,
				URL: url,
			})
		}
	}

	return presignedURLs, true
}

func init() {
	go taskGC()
}
