package collector

import (
	"encoding/json"
	"omciAnalyzer/utils"
	"os"
	"path/filepath"
)

type Module struct {
	ID         int    `json:"id"`
	ModuleName string `json:"module_name"`
	Level      int    `json:"level,omitempty"`
}

var logger_modules map[string][]Module = make(map[string][]Module)

func GetLoggerModules() map[string][]Module {
	return logger_modules
}

func newFileLoggerModule(path string) ([]Module, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return []Module{}, err
	}
	var modules []Module
	err = json.Unmarshal(content, &modules)
	if err != nil {
		return []Module{}, err
	}
	return modules, nil
}

func loadLoggerModules(path string) {
	// read directory all json files
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		return
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}

	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			continue
		}

		fullPath := filepath.Join(path, e.Name())

		if info.IsDir() {
			continue
		}
		filename := e.Name()
		if len(filename) < 6 || filename[len(filename)-5:] != ".json" {
			continue
		}
		moduleName := filename[:len(filename)-5]
		modules, err := newFileLoggerModule(fullPath)
		if err != nil {
			utils.Log("Failed to load logger module:", moduleName, "error:", err)
			continue
		}
		logger_modules[moduleName] = modules
	}
}

var loggerModulePath string = "./resource/logger/"

func init() {
	loadLoggerModules(loggerModulePath)
}
