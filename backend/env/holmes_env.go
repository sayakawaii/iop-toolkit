package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

type HolmesSoloConfig interface{}

type HolmesCommonConfig struct {
	Interval   string
	DumpPath   string
	DumpType   string
	DumpCpuMax int
	Enable     bool
}

type HolmesSoloConfig1 struct {
	Enable   bool
	Min      int
	Abs      int
	Diff     int
	CoolDown int
}

type HolmesSoloConfig2 struct {
	HolmesSoloConfig1
	Max int
}

type ConfigType int

const (
	Mem ConfigType = iota
	Cpu
	Thread
	GcHeap
	Goroutine
	SoloConfigEnd
)

const (
	SoloConfig1End ConfigType = GcHeap + 1
	SoloConfig2End            = Goroutine + 1
)

var type2name = map[ConfigType]string{
	Mem:       "MEM",
	Cpu:       "CPU",
	Thread:    "THREAD",
	Goroutine: "GOROUTINE",
	GcHeap:    "GC_HEAP",
}

type HolmesConfig struct {
	Common HolmesCommonConfig
	// Mem       HolmesSoloConfig1[0]
	// Cpu       HolmesSoloConfig1[1]
	// Thread    HolmesSoloConfig1[2]
	// GcHeap    HolmesSoloConfig1[3]

	// Goroutine HolmesSoloConfig2[4]
	Solo [SoloConfigEnd]HolmesSoloConfig
}

// New returns a new Config struct
func NewHolmesConfig() *HolmesConfig {
	res := &HolmesConfig{
		Common: HolmesCommonConfig{
			Interval:   getEnv("HOLMES_COLLECTTION_INTERVAL", "5s"),
			DumpPath:   getEnv("HOLMES_DUMP_PATH", "/tmp"),
			DumpType:   getEnv("HOLMES_DUMP_TYPE", "text"),
			DumpCpuMax: getEnvAsInt("HOLMES_DUMP_CPU_MAX", 90),
		},
	}
	var i ConfigType

	for i = 0; i < SoloConfig1End; i++ {
		res.Solo[i] = HolmesSoloConfig1{
			Enable:   getEnvAsBool("HOLMES_"+type2name[i]+"_DUMP_ENABLE", false),
			Min:      getEnvAsInt("HOLMES_"+type2name[i]+"_DUMP_TRIGGER_MIN", 5),
			Abs:      getEnvAsInt("HOLMES_"+type2name[i]+"_DUMP_TRIGGER_ABS", 80),
			Diff:     getEnvAsInt("HOLMES_"+type2name[i]+"_DUMP_TRIGGER_DIFF", 25),
			CoolDown: getEnvAsInt("HOLMES_"+type2name[i]+"_DUMP_TRIGGER_COOLDOWN", 60),
		}
		if !res.Common.Enable && res.Solo[i].(HolmesSoloConfig1).Enable {
			res.Common.Enable = true
		}
	}
	for i = SoloConfig1End; i < SoloConfig2End; i++ {
		res.Solo[i] = HolmesSoloConfig2{
			HolmesSoloConfig1: HolmesSoloConfig1{
				Enable:   getEnvAsBool("HOLMES_"+type2name[i]+"_DUMP_ENABLE", false),
				Min:      getEnvAsInt("HOLMES_"+type2name[i]+"_DUMP_TRIGGER_MIN", 5),
				Abs:      getEnvAsInt("HOLMES_"+type2name[i]+"_DUMP_TRIGGER_ABS", 80),
				Diff:     getEnvAsInt("HOLMES_"+type2name[i]+"_DUMP_TRIGGER_DIFF", 25),
				CoolDown: getEnvAsInt("HOLMES_"+type2name[i]+"_DUMP_TRIGGER_COOLDOWN", 60),
			},
			Max: getEnvAsInt("HOLMES_"+type2name[i]+"_DUMP_TRIGGER_MAX", 90),
		}
		if !res.Common.Enable && res.Solo[i].(HolmesSoloConfig2).Enable {
			res.Common.Enable = true
		}
	}
	return res
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
// func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
// 	valStr := getEnv(name, "")

// 	if valStr == "" {
// 		return defaultVal
// 	}

// 	val := strings.Split(valStr, sep)

// 	return val
// }
