package api

import "omciAnalyzer/service/omcianalyzer/omciSchema"

type EntityHandler interface {
	HandleEntity(msg *omciSchema.OmciContext) (string, error)
}

var (
	entityRegistry = make(map[uint64]EntityHandler)
)

func RegisterEntityHandler(entityID uint64, h EntityHandler) {
	entityRegistry[entityID] = h
}

func GetEntityHandler(entityID uint64) EntityHandler {
	return entityRegistry[entityID]
}
